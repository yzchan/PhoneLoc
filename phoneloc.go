package phoneloc

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"os"
)

const (
	Others       byte = 0b00
	ChinaMobile  byte = 0b01 // CMCC
	ChinaTelecom byte = 0b10 // CTCC
	ChinaUnicom  byte = 0b11 // CUCC
)

var ispMap = map[byte]string{
	0b00: "其他",
	0b01: "中国移动",
	0b10: "中国电信",
	0b11: "中国联通",
}

type PhoneLoc struct {
	Prov     string
	ProvCode int
	City     string
	CityCode int
	Isp      string // 运营商
	Virtual  bool   // 是否虚拟号段
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
	loc = &PhoneLoc{}
	if sec < 1000000 || sec > 2000000 {
		return nil, errors.New("invalid phone section")
	}
	mac := sec / 10000
	if mac > 170 && mac < 180 { // 虚拟号段
		loc.Virtual = true
	}
	hlr := sec % 10000
	blockId := int(p.buffer[mac])
	offset := 200 + (blockId-1)*3*10000 + hlr*3
	buff := make([]byte, 4)
	copy(buff, p.buffer[offset:offset+3])
	ispBits := buff[2] >> 6
	loc.Isp, _ = ispMap[ispBits]
	buff[2] = buff[2] & 0b00111111
	loc.CityCode = int(binary.LittleEndian.Uint32(buff))
	loc.ProvCode = (loc.CityCode / 10000) * 10000
	loc.Prov, _ = ProvAndCities[loc.ProvCode]
	loc.City, _ = ProvAndCities[loc.CityCode]
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
