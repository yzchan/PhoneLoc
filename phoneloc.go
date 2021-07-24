package phoneloc

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"os"
)

var ispMapping = map[byte]string{
	0b00: "其他", // OtherIsp
	0b01: "移动", // ChinaMobile
	0b10: "电信", // ChinaTelecom
	0b11: "联通", // ChinaUnicom
}

type PhoneLoc struct {
	Section  int    `json:"section"`  // 号段
	Prov     string `json:"prov"`     // 归属省
	City     string `json:"city"`     // 归属城市
	Adcode   int    `json:"adcode"`   // 归属城市的行政区划代码
	Postcode int    `json:"postcode"` // 邮编
	Telcode  int    `json:"telcode"`  // 电话区号
	Sp       string `json:"sp"`       // 运营商
	Virtual  bool   `json:"virtual"`  // 是否虚拟号段
}

type Parser struct {
	buffer []byte
	len    int
}

func NewParser(file string) (p *Parser, err error) {
	p = &Parser{}

	f, err := os.OpenFile(file, os.O_RDONLY, 0400)
	if err != nil {
		return
	}
	defer f.Close()

	buffer, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	p.buffer = buffer
	return
}

func (p *Parser) Find(sec int) (loc *PhoneLoc, err error) {
	if sec < 1000000 || sec > 2000000 {
		return nil, errors.New("invalid phone section")
	}
	loc = &PhoneLoc{Section: sec}
	mac := sec / 10000
	if mac > 170 && mac < 180 { // 虚拟号段
		loc.Virtual = true
	}
	hlr := sec % 10000
	blockId := int(p.buffer[mac])
	if blockId == 0 {
		return nil, errors.New("invalid phone section")
	}
	offset := 200 + (blockId-1)*3*10000 + hlr*3
	buff := make([]byte, 4)
	copy(buff, p.buffer[offset:offset+3])
	ispBits := buff[2] >> 6
	loc.Sp, _ = ispMapping[ispBits]
	buff[2] = buff[2] & 0b00111111
	loc.Adcode = int(binary.LittleEndian.Uint32(buff))
	provCode := (loc.Adcode / 10000) * 10000
	loc.Prov, _ = DistrictMapping[provCode]
	loc.City, _ = DistrictMapping[loc.Adcode]
	return
}

func (p *Parser) Macs() (ret []int) {
	ret = make([]int, 0)
	for i := 100; i < 200; i++ {
		if p.buffer[i] != 0x00 {
			ret = append(ret, i)
		}
	}
	return
}

func (p *Parser) Version() string {
	end := 0
	for {
		end++
		if p.buffer[end] == 0x00 {
			break
		}
	}
	return string(p.buffer[0:end])
}
