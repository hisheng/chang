package caiwu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/xueqiu"
	"net/url"
	"strconv"
)

var ZichanfuzhaiRequest  ZichanfuzhaiRequest_



type Zichanfuzhai_ struct {
	Report_type string
	Symbol string `sql:"type:varchar(20)"`
	Report_name string //1546185600000
	Report_date int  //"2018年报"
	GatherDay string
	Currency_funds float64 //货币资金
	ar_and_br float64 //应收票据及应收账款
	bills_receivable float64 //其中：应收票据
	account_receivable float64 //应收账款
	pre_payment float64 //预付款项
	othr_receivables float64 //其他应收款
	inventory float64 //存货
	othr_current_assets float64 //其他流动资产
	total_current_assets float64 //流动资产合计
	lt_receivable float64 //长期应收款
	lt_equity_invest float64 //长期股权投资
	other_illiquid_fnncl_assets float64 //其他非流动金融资产
	invest_property float64 //投资性房地产
	fixed_asset_sum float64 //固定资产合计
	fixed_asset float64 //其中：固定资产
	intangible_assets float64 //无形资产
	lt_deferred_expense float64 //长期待摊费用


}

func (z Zichanfuzhai_) getTableName() string {
	return "zichanfuzhai"
}

func (z Zichanfuzhai_) createTable()  {
	tableKey := z.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&z)
	}
}


func (z Zichanfuzhai_)Add(){
	db.DB.Table(z.getTableName()).Create(&z)
}





type ZichanfuzhaiRequest_ xueqiu.Request_
func (request ZichanfuzhaiRequest_) initRequest() ZichanfuzhaiRequest_{
	request.SearchUrl  = "https://stock.xueqiu.com/v5/stock/finance/cn/balance.json"

	request.SearchParms = url.Values{}
	request.SearchParms.Add("symbol","SH601155")
	request.SearchParms.Add("type","all")
	request.SearchParms.Add("is_detail","true")
	request.SearchParms.Add("count","20")
	request.SearchParms.Add("timestamp","")
	return request
}






func (request ZichanfuzhaiRequest_) Run ()  {
	Xianjinliu.createTable()

	request = request.initRequest()
	fmt.Println(request.SearchParms.Get("type"))

	for i:= 1;i<=4;i++ {
		request.SearchParms.Set("type","Q"+strconv.Itoa(i))  //Q1代表一季度
		//fmt.Println(request)
		request.RunGet()
	}

}


func (request ZichanfuzhaiRequest_) RunGet()  {
	data := xueqiu.Get(request.SearchUrl,request.SearchParms)

	str:=[]byte(data)


	rs := ZichanfuzhaiJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	for _,stock := range rs.Data.List{

		fmt.Println(stock)

		//xianjin.Add()
	}
}


type  ZichanfuzhaiJsonResponse struct{
	ErrorCode int `json:"error_code"`
	Data ZichanfuzhaiJsonData  `json:"data"`
}

type  ZichanfuzhaiJsonData struct {
	Symbol       string
	QuoteName    string
	Currency     string
	CurrencyName string
	OrgType      int
	List         []ZichanfuzhaiJsonDataItem
}

type  ZichanfuzhaiJsonDataItem struct{
	Report_type string
	Symbol string `sql:"type:varchar(20)"`
	Report_name string //1546185600000
	Report_date int  //"2018年报"
}
