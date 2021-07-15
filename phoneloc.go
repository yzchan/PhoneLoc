package phoneloc

const (
	ChinaMobile  string = "中国移动" // CMCC
	ChinaTelecom string = "中国电信" // CTCC
	ChinaUnicom  string = "中国联通" // CUCC
)

type PhoneLoc struct {
	Prov     string
	ProvCode int
	City     string
	CityCode int
	Isp      string // 运营商
	Virtual  bool   // 是否虚拟号段
}
