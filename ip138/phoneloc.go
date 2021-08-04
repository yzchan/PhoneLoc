package ip138

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

const macChromeUA string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"

func PhoneLoc(phone string) (address string, sp string, err error) {
	client := http.Client{}
	url := fmt.Sprintf("https://ip138.com/mobile.asp?mobile=%s&action=mobile", phone)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("User-Agent", macChromeUA)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	s := dom.Find("div.table table tbody")
	address = s.Find("tr ").Eq(1).Find("td").Eq(1).Text()
	sp = s.Find("tr ").Eq(2).Find("td").Eq(1).Text()
	return
}
