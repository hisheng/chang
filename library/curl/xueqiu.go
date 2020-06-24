package curl

import (
	"fmt"
	"github.com/hisheng/chang/conf"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
包含了 雪球的 cookie 封装的 get 请求
*/
func XueqiuGet(url string, parms url.Values) string {
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
