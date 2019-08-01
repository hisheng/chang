package xueqiu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/curl"
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
	Request.SearchParms.Add("size","60")
	Request.SearchParms.Add("current","")
	Request.SearchParms.Add("pettm","4_10")
	Request.SearchParms.Add("oiy.20190331","10_1000")
	Request.SearchParms.Add("npay.20190331","10_1000")
	Request.SearchParms.Add("_","1564377012355")
	Request.SearchParms.Add("pct","")
	Request.SearchParms.Add("only_count","0")
}

func (request_ Request_) Run ()  {
	data := curl.Get(request_.SearchUrl,request_.SearchParms)

	fmt.Println(data)

	str:=[]byte(data)


	rs := XueqiuJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs)

	fmt.Println(rs.Data.List)

	for _,stock := range rs.Data.List{
		stock.InitOiyPe()
		s := fmt.Sprintf("%s %4s 营业收入增长 %.2f 利润增长 %.2f pe为 %.2f 性价比指数为 %.2f",stock.Symbol,stock.Name,stock.Oiy,stock.Npay,stock.Pettm,stock.Npay/stock.Pettm)
		fmt.Println(s)
		//return fmt.Sprintf("sm,SM,um,UM=%d,%d,%d,%d", l.min, l.max, l.umin, l.umax)

	}
}



