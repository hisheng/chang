package main

import (
	"fmt"
	"github.com/hisheng/chang/curl"
	"net/url"
)

func main()  {
	xueqiuUrl := "https://xueqiu.com/service/screener/screen?category=CN&exchange=sh_sz&areacode=&indcode=&order_by=pettm&order=asc&page=1&size=300&only_count=0&current=&pct=&pettm=4_10&oiy.20190331=-99.42_44065.32&npay.20190331=-28250.4_80737.8&_=1564377012355"

	r := curl.Get(xueqiuUrl,url.Values{})
	fmt.Println(r)
}