package xueqiu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/jinzhu/gorm"
	"net/url"
	"strconv"
	"time"
)

/*
https://stock.xueqiu.com/v5/stock/finance/cn/income.json?symbol=SH601155&type=Q4&is_detail=true&count=20&timestamp=
*/

var Lirunbiao Lirunbiao_
var LirunbiaoRequest LirunbiaoRequest_

type Lirunbiao_ struct {
	gorm.Model
	Report_type string
	Symbol string `sql:"type:varchar(20)"`
	Report_name string //1546185600000
	Report_date int  //"2018年报"
	Total_revenue float64  //营业总收入
	Revenue float64 //营业收入
	OperatingCosts  //营业总成本构成
	Op float64 //营业利润
	Non_operating_income float64 //加：营业外收入
	Non_operating_payout float64 //减：营业外支出
	Profit_total_amt float64  //利润总额
	Income_tax_expenses float64 //减：所得税费用
	Net_profit float64 //净利润
	Continous_operating_np float64 //（一）持续经营净利润
	Net_profit_atsopc float64 //归属于母公司所有者的净利润
	Minority_gal float64 //少数股东损益
	Net_profit_after_nrgal_atsolc float64 //扣除非经常性损益后的净利润
	Basic_eps float64  //基本每股收益
	Dlt_earnings_per_share float64 //稀释每股收益
	CompreIncome //公司与综合收益
	LirunbiaoPercent
}

type CompreIncome struct {
	Othr_compre_income float64 //其他综合收益
	Othr_compre_income_atoopc float64 //归属母公司所有者的其他综合收益
	Total_compre_income float64 //综合收益总额
	Total_compre_income_atsopc float64 //归属于母公司股东的综合收益总额
	Total_compre_income_atms float64 //归属于少数股东的综合收益总额
}

//营业总成本
type OperatingCosts struct {
	Operating_costs float64 //营业总成本
	Operating_cost float64  //营业成本
	Operating_taxes_and_surcharge float64 //营业税金及附加
	Sales_fee float64 //销售费用
	Manage_fee float64  //管理费用
	Financing_expenses float64 //财务费用
	Finance_cost_interest_fee float64 //利息费用
	Finance_cost_interest_income float64 //利息收入
	Asset_impairment_loss float64  //资产减值损失
	Credit_impairment_loss float64 //信用减值损失
	Income_from_chg_in_fv  float64 //加：公允价值变动收益
	Invest_income float64    //投资收益
	Invest_incomes_from_rr float64 //其中：对联营企业和合营企业的投资收益
	Asset_disposal_income float64 //资产处置收益
	Other_income float64  //其他收益
	GatherDay string
}

type LirunbiaoPercent struct {
	Total_revenue_percent float64  //营业总收入
	Revenue_percent float64  //营业收入
	Operating_costs_percent float64 //营业总成本
	Operating_cost_percent float64  //营业成本
	Operating_taxes_and_surcharge_percent float64 //营业税金及附加
	Sales_fee_percent float64 //销售费用
	Manage_fee_percent float64  //管理费用
	Financing_expenses_percent float64 //财务费用
	Finance_cost_interest_fee_percent float64 //利息费用
	Finance_cost_interest_income_percent float64 //利息收入
	Asset_impairment_loss_percent float64  //资产减值损失
	Credit_impairment_loss_percent float64 //信用减值损失
	Income_from_chg_in_fv_percent  float64 //加：公允价值变动收益
	Invest_income_percent float64    //投资收益
	Invest_incomes_from_rr_percent float64 //其中：对联营企业和合营企业的投资收益
	Asset_disposal_income_percent float64 //资产处置收益
	Other_income_percent float64  //其他收益
	Op_percent float64 //营业利润
	Non_operating_income_percent float64 //加：营业外收入
	Non_operating_payout_percent float64 //减：营业外支出
	Profit_total_amt_percent float64  //利润总额
	Income_tax_expenses_percent float64 //减：所得税费用
	Net_profit_percent float64 //净利润
	Continous_operating_np_percent float64 //（一）持续经营净利润
	Net_profit_atsopc_percent float64 //归属于母公司所有者的净利润
	Minority_gal_percent float64 //少数股东损益
	Net_profit_after_nrgal_atsolc_percent float64 //扣除非经常性损益后的净利润
	Basic_eps_percent float64  //基本每股收益
	Dlt_earnings_per_share_percent float64 //稀释每股收益
	Othr_compre_income_percent float64 //其他综合收益
	Othr_compre_income_atoopc_percent float64 //归属母公司所有者的其他综合收益
	Total_compre_income_percent float64 //综合收益总额
	Total_compre_income_atsopc_percent float64 //归属于母公司股东的综合收益总额
	Total_compre_income_atms_percent float64 //归属于少数股东的综合收益总额
}

func (l Lirunbiao_) getTableName() string {
	return "lirunbiao"
}

func (l Lirunbiao_) CreateTable()  {
	tableKey := l.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&l)
		db.DB.Table(tableKey).AddUniqueIndex(l.getTableName() +"symbol_gatherday", "symbol","gather_day")
	}
}


func (l Lirunbiao_)Add() Lirunbiao_{
	is := l.FindOne()
	if is.ID == 0 {
		db.DB.Table(l.getTableName()).Create(&l)
	}else {
		return is
	}
	return l
	//Exce 会释放连接池
}

func (l Lirunbiao_) FindOne() Lirunbiao_{
	dbn := db.DB.Table(l.getTableName()).Where("symbol = ?",l.Symbol).Where("gather_day = ?",l.GatherDay)
	dbn.First(&l)
	return l
}





type LirunbiaoRequest_ Request_
func (request LirunbiaoRequest_) initRequest(symbol string) LirunbiaoRequest_{
	request.SearchUrl  = "https://stock.xueqiu.com/v5/stock/finance/cn/income.json"

	request.SearchParms = url.Values{}
	request.SearchParms.Add("symbol",symbol)
	request.SearchParms.Add("type","all")
	request.SearchParms.Add("is_detail","true")
	request.SearchParms.Add("count","20")
	request.SearchParms.Add("timestamp","")
	return request
}



func (request LirunbiaoRequest_) Run (symbol string)  {

	request = request.initRequest(symbol)
	fmt.Println(request.SearchParms.Get("type"))

	for i:= 1;i<=4;i++ {
		fmt.Println("LirunbiaoRequest_ "+ symbol + " Q" + strconv.Itoa(i))
		request.SearchParms.Set("type","Q"+strconv.Itoa(i))  //Q1代表一季度
		request.RunGet()
	}

}

func (request LirunbiaoRequest_) RunGet()  {
	data := Get(request.SearchUrl,request.SearchParms)

	str:=[]byte(data)


	rs := LirunbiaoJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	for _,stock := range rs.Data.List{
		l := Lirunbiao_{}
		l.Symbol = request.SearchParms.Get("symbol")
		l.Report_name = stock.Report_name

		val := stock.Report_date /1000
		int64_ := int64(val)
		tm := time.Unix(int64_ , 0)
		l.GatherDay = tm.Format("2006-01-02")
		l.Report_date = val

		l.Report_type = request.SearchParms.Get("type")
		l.Net_profit = stock.Net_profit[0]
		l.Net_profit_percent = stock.Net_profit[1]
		l.Net_profit_atsopc = stock.Net_profit_atsopc[0]
		l.Net_profit_atsopc_percent = stock.Net_profit_atsopc[1]
		l.Total_revenue = stock.Total_revenue[0]
		l.Total_revenue_percent = stock.Total_revenue[1]
		l.Op = stock.Op[0]
		l.Op_percent = stock.Op[1]
		l.Income_from_chg_in_fv = stock.Income_from_chg_in_fv[0]
		l.Income_from_chg_in_fv_percent = stock.Income_from_chg_in_fv[1]
		l.Invest_incomes_from_rr = stock.Invest_incomes_from_rr[0]
		l.Invest_incomes_from_rr_percent = stock.Invest_incomes_from_rr[1]
		l.Invest_income = stock.Invest_income[0]
		l.Invest_income_percent = stock.Invest_income[1]
		l.Operating_taxes_and_surcharge = stock.Operating_taxes_and_surcharge[0]
		l.Operating_taxes_and_surcharge_percent = stock.Operating_taxes_and_surcharge[1]
		l.Asset_impairment_loss = stock.Asset_impairment_loss[0]
		l.Asset_impairment_loss_percent = stock.Asset_impairment_loss[1]
		l.Non_operating_income = stock.Non_operating_income[0]
		l.Non_operating_income_percent = stock.Non_operating_income[1]
		l.Non_operating_payout = stock.Non_operating_payout[0]
		l.Non_operating_payout_percent = stock.Non_operating_payout[1]
		l.Profit_total_amt = stock.Profit_total_amt[0]
		l.Profit_total_amt_percent = stock.Profit_total_amt[1]
		l.Minority_gal = stock.Minority_gal[0]
		l.Minority_gal_percent = stock.Minority_gal[1]
		l.Basic_eps = stock.Basic_eps[0]
		l.Basic_eps_percent = stock.Basic_eps[1]
		l.Dlt_earnings_per_share = stock.Dlt_earnings_per_share[0]
		l.Dlt_earnings_per_share_percent = stock.Dlt_earnings_per_share[1]
		l.Othr_compre_income_atoopc = stock.Othr_compre_income_atoopc[0]
		l.Othr_compre_income_atoopc_percent = stock.Othr_compre_income_atoopc[1]
		l.Total_compre_income = stock.Total_compre_income[0]
		l.Total_compre_income_percent = stock.Total_compre_income[1]
		l.Total_compre_income_atsopc = stock.Total_compre_income_atsopc[0]
		l.Total_compre_income_atsopc_percent = stock.Total_compre_income_atsopc[1]
		l.Total_compre_income_atms = stock.Total_compre_income_atms[0]
		l.Total_compre_income_atms_percent = stock.Total_compre_income_atms[1]
		l.Othr_compre_income = stock.Othr_compre_income[0]
		l.Othr_compre_income_percent = stock.Othr_compre_income[1]
		l.Net_profit_after_nrgal_atsolc = stock.Net_profit_after_nrgal_atsolc[0]
		l.Net_profit_after_nrgal_atsolc_percent = stock.Net_profit_after_nrgal_atsolc[1]
		l.Income_tax_expenses = stock.Income_tax_expenses[0]
		l.Income_tax_expenses_percent = stock.Income_tax_expenses[1]
		l.Credit_impairment_loss = stock.Credit_impairment_loss[0]
		l.Credit_impairment_loss_percent = stock.Credit_impairment_loss[1]
		l.Revenue = stock.Revenue[0]
		l.Revenue_percent = stock.Revenue[1]
		l.Operating_costs = stock.Operating_costs[0]
		l.Operating_costs_percent = stock.Operating_costs[1]
		l.Operating_cost = stock.Operating_cost[0]
		l.Operating_cost_percent = stock.Operating_cost[1]
		l.Sales_fee = stock.Sales_fee[0]
		l.Sales_fee_percent = stock.Sales_fee[1]
		l.Manage_fee = stock.Manage_fee[0]
		l.Manage_fee_percent = stock.Manage_fee[1]
		l.Financing_expenses = stock.Financing_expenses[0]
		l.Financing_expenses_percent = stock.Financing_expenses[1]
		l.Finance_cost_interest_fee = stock.Finance_cost_interest_fee[0]
		l.Finance_cost_interest_fee_percent = stock.Finance_cost_interest_fee[1]
		l.Finance_cost_interest_income = stock.Finance_cost_interest_income[0]
		l.Finance_cost_interest_income_percent = stock.Finance_cost_interest_income[1]
		l.Asset_disposal_income = stock.Asset_disposal_income[0]
		l.Asset_disposal_income_percent = stock.Asset_disposal_income[1]
		l.Other_income = stock.Other_income[0]
		l.Other_income_percent = stock.Other_income[1]
		l.Continous_operating_np = stock.Continous_operating_np[0]
		l.Continous_operating_np_percent = stock.Continous_operating_np[1]
		l.Add()
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
	Report_name string //1546185600000
	Report_date int  //"2018年报"
	Total_revenue [2]float64  //营业总收入
	Revenue [2]float64 //营业收入
	Op [2]float64 //营业利润
	Non_operating_income [2]float64 //加：营业外收入
	Non_operating_payout [2]float64 //减：营业外支出
	Profit_total_amt [2]float64  //利润总额
	Income_tax_expenses [2]float64 //减：所得税费用
	Net_profit [2]float64 //净利润
	Continous_operating_np [2]float64 //（一）持续经营净利润
	Net_profit_atsopc [2]float64 //归属于母公司所有者的净利润
	Minority_gal [2]float64 //少数股东损益
	Net_profit_after_nrgal_atsolc [2]float64 //扣除非经常性损益后的净利润
	Basic_eps [2]float64  //基本每股收益
	Dlt_earnings_per_share [2]float64 //稀释每股收益
	Othr_compre_income [2]float64 //其他综合收益
	Othr_compre_income_atoopc [2]float64 //归属母公司所有者的其他综合收益
	Total_compre_income [2]float64 //综合收益总额
	Total_compre_income_atsopc [2]float64 //归属于母公司股东的综合收益总额
	Total_compre_income_atms [2]float64 //归属于少数股东的综合收益总额
	Operating_costs [2]float64 //营业总成本
	Operating_cost [2]float64  //营业成本
	Operating_taxes_and_surcharge [2]float64 //营业税金及附加
	Sales_fee [2]float64 //销售费用
	Manage_fee [2]float64  //管理费用
	Financing_expenses [2]float64 //财务费用
	Finance_cost_interest_fee [2]float64 //利息费用
	Finance_cost_interest_income [2]float64 //利息收入
	Asset_impairment_loss [2]float64  //资产减值损失
	Credit_impairment_loss [2]float64 //信用减值损失
	Income_from_chg_in_fv  [2]float64 //加：公允价值变动收益
	Invest_income [2]float64    //投资收益
	Invest_incomes_from_rr [2]float64 //其中：对联营企业和合营企业的投资收益
	Asset_disposal_income [2]float64 //资产处置收益
	Other_income [2]float64  //其他收益
}












