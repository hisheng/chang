package xueqiu

import (
	"net/url"
)

var Request  Request_

func init()()  {
	//Request.SearchUrl = "https://xueqiu.com/service/screener/screen?category=CN&exchange=sh_sz&areacode=&indcode=&order_by=pettm&order=asc&page=1&size=2&only_count=0&current=&pct=&pettm=4_10&oiy.20190331=-99.42_44065.32&npay.20190331=-28250.4_80737.8&_=1564377012355"
	Request.SearchUrl = "https://xueqiu.com/service/screener/screen"
	initSearchParms()
}

type Request_ struct {
	SearchUrl string
	SearchParms url.Values
 }

func initSearchParms()  {
	Request.SearchParms = url.Values{}
	Request.SearchParms.Add("category","CN")
	Request.SearchParms.Add("exchange","sh_sz")
	Request.SearchParms.Add("order_by","pettm")
	Request.SearchParms.Add("order","asc")
	Request.SearchParms.Add("page","1")
	Request.SearchParms.Add("size","2")
	Request.SearchParms.Add("current","")
	Request.SearchParms.Add("pettm","4_10")
	Request.SearchParms.Add("oiy.20190331","-99.42_44065.32")
	Request.SearchParms.Add("npay.20190331","-28250.4_80737.8")
	Request.SearchParms.Add("_","1564377012355")
	Request.SearchParms.Add("pct","")
	Request.SearchParms.Add("only_count","0")


}


