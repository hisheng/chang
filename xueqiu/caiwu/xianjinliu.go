package caiwu

import (
	"fmt"
	"github.com/hisheng/chang/xueqiu"
	"net/url"
	"strconv"
)

var XianjinliuRequest XianjinliuRequest_













type XianjinliuRequest_ xueqiu.Request_
func (request XianjinliuRequest_) initRequest() XianjinliuRequest_{
	request.SearchUrl  = "https://stock.xueqiu.com/v5/stock/finance/cn/cash_flow.json"

	request.SearchParms = url.Values{}
	request.SearchParms.Add("symbol","SH601155")
	request.SearchParms.Add("type","all")
	request.SearchParms.Add("is_detail","true")
	request.SearchParms.Add("count","20")
	request.SearchParms.Add("timestamp","")
	return request
}






func (request XianjinliuRequest_) Run ()  {
	//Lirunbiao.createTable()

	request = request.initRequest()
	fmt.Println(request.SearchParms.Get("type"))

	for i:= 1;i<=4;i++ {
		request.SearchParms.Set("type","Q"+strconv.Itoa(i))  //Q1代表一季度
		//request.RunGet()
		fmt.Println(request)
	}

}



type  XianjinliuJsonResponse struct{
	ErrorCode int `json:"error_code"`
	Data LirunbiaoJsonData  `json:"data"`
}

type  XianjinliuJsonData struct {
	Symbol       string
	QuoteName    string
	Currency     string
	CurrencyName string
	OrgType      int
	List         []LirunbiaoJsonDataItem
}