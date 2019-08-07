package caiwu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/xueqiu"
	"net/url"
)

/*
https://stock.xueqiu.com/v5/stock/finance/cn/income.json?symbol=SH601155&type=Q4&is_detail=true&count=20&timestamp=
*/

var Lirunbiao Lirunbiao_
var LirunbiaoRequest LirunbiaoRequest_

type Lirunbiao_ struct {
	report_name string //1546185600000
	report_date string  //"2018年报"
	total_revenue float64  //营业总收入
	//operatingCosts  //营业总成本构成
	op float64 //营业利润
	non_operating_income float64 //加：营业外收入
	non_operating_payout float64 //减：营业外支出
	profit_total_amt float64  //利润总额
	income_tax_expenses float64 //减：所得税费用
	net_profit float64 //净利润
	continous_operating_np float64 //（一）持续经营净利润
	net_profit_atsopc float64 //归属于母公司所有者的净利润
	minority_gal float64 //少数股东损益
	net_profit_after_nrgal_atsolc float64 //扣除非经常性损益后的净利润
	basic_eps float64  //基本每股收益
	dlt_earnings_per_share float64 //稀释每股收益
	othr_compre_income float64 //其他综合收益
	othr_compre_income_atoopc float64 //归属母公司所有者的其他综合收益
	total_compre_income float64 //综合收益总额
	total_compre_income_atsopc float64 //归属于母公司股东的综合收益总额
	total_compre_income_atms float64 //归属于少数股东的综合收益总额
}

//营业总成本
type operatingCosts struct {
	operating_costs float64 //营业总成本
	operating_cost float64  //营业成本
	operating_taxes_and_surcharge float64 //营业税金及附加
	sales_fee float64 //销售费用
	manage_fee float64  //管理费用
	financing_expenses float64 //财务费用
	finance_cost_interest_fee float64 //利息费用
	finance_cost_interest_income float64 //利息收入
	asset_impairment_loss float64  //资产减值损失
	credit_impairment_loss float64 //信用减值损失
	income_from_chg_in_fv  float64 //加：公允价值变动收益
	invest_income float64    //投资收益
	invest_incomes_from_rr float64 //其中：对联营企业和合营企业的投资收益
	asset_disposal_income float64 //资产处置收益
	other_income float64  //其他收益
}

func (s Lirunbiao_) getTableName() string {
	return "lirunbiao"
}

func (s Lirunbiao_) createTable()  {
	tableKey := s.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&s)
	}
}


type LirunbiaoRequest_ xueqiu.Request_
func (request LirunbiaoRequest_) initRequest() LirunbiaoRequest_{
	request.SearchUrl  = "https://stock.xueqiu.com/v5/stock/finance/cn/income.json"

	request.SearchParms = url.Values{}
	request.SearchParms.Add("symbol","SH601155")
	request.SearchParms.Add("type","Q4")
	request.SearchParms.Add("is_detail","true")
	request.SearchParms.Add("count","20")
	request.SearchParms.Add("timestamp","")
	return request
}



func (request LirunbiaoRequest_) Run ()  {
	//Lirunbiao.createTable()

	request = request.initRequest()
	data := xueqiu.Get(request.SearchUrl,request.SearchParms)

	str:=[]byte(data)


	rs := LirunbiaoJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	for _,stock := range rs.Data.List{
		sm := make(map[string]interface{})
		sm["symbol"] = xueqiu.GetSymbol(request.SearchParms.Get("symbol"))
		fmt.Println(stock)
	}
}


//chart json
type  LirunbiaoJsonResponse struct{
	ErrorCode int `json:"error_code"`
	Data LirunbiaoJsonData  `json:"data"`
}

type  LirunbiaoJsonData struct{
	Symbol string
	QuoteName string
	Currency string
	CurrencyName string
	OrgType int
	List []LirunbiaoJsonDataItem
}
type  LirunbiaoJsonDataItem struct{
	ReportName string //1546185600000
	ReportDate string  //"2018年报"
	Total_revenue []float64  //营业总收入
	Op []float64 //营业利润
	Non_operating_income []float64 //加：营业外收入
	Non_operating_payout []float64 //减：营业外支出
	Profit_total_amt []float64  //利润总额
	Income_tax_expenses []float64 //减：所得税费用
	Net_profit []float64 //净利润
	Continous_operating_np []float64 //（一）持续经营净利润
	Net_profit_atsopc []float64 //归属于母公司所有者的净利润
	Minority_gal []float64 //少数股东损益
	Net_profit_after_nrgal_atsolc []float64 //扣除非经常性损益后的净利润
	Basic_eps []float64  //基本每股收益
	Dlt_earnings_per_share []float64 //稀释每股收益
	Othr_compre_income []float64 //其他综合收益
	Othr_compre_income_atoopc []float64 //归属母公司所有者的其他综合收益
	Total_compre_income []float64 //综合收益总额
	Total_compre_income_atsopc []float64 //归属于母公司股东的综合收益总额
	Total_compre_income_atms []float64 //归属于少数股东的综合收益总额
	Operating_costs []float64 //营业总成本
	Operating_cost []float64  //营业成本
	Operating_taxes_and_surcharge []float64 //营业税金及附加
	Sales_fee []float64 //销售费用
	Manage_fee []float64  //管理费用
	Financing_expenses []float64 //财务费用
	Finance_cost_interest_fee []float64 //利息费用
	Finance_cost_interest_income []float64 //利息收入
	Asset_impairment_loss []float64  //资产减值损失
	Credit_impairment_loss []float64 //信用减值损失
	Income_from_chg_in_fv  []float64 //加：公允价值变动收益
	Invest_income []float64    //投资收益
	Invest_incomes_from_rr []float64 //其中：对联营企业和合营企业的投资收益
	Asset_disposal_income []float64 //资产处置收益
	Other_income []float64  //其他收益
}












