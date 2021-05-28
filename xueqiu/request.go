package xueqiu

import (
	"fmt"
	"github.com/hisheng/chang/conf"
	"io/ioutil"
	"net/http"
	"net/url"
)

var XueqiuRequest Request_

type Request_ struct {
	SearchUrl   string
	SearchParms url.Values
}

func Get(url string, parms url.Values) string {
	cookies := conf.XueqiuCookie.Cookies

	client := &http.Client{}

	req, err := http.NewRequest("GET", url+"?"+parms.Encode(), nil)
	if err != nil {
		fmt.Println("获取地址错误")
	}
	req.Header.Set("Cookie", cookies)
	//req.Header.Add("Agent",GetRandomUserAgent() )
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("登录错误")
	}
	resp_byte, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	return string(resp_byte)

}

func GetSymbol(symbol string) string {
	switch symbol[0:2] {
	case "SH":
		return symbol[2:]
	default:
		return ""
	}

}

func createTable() {
	Lirunbiao.CreateTable()
	Zichanfuzhai.CreateTable()
	Xianjinliu.CreateTable()
	StockChartModel.CreateTable()
	StockQuote.CreateTable()
}
