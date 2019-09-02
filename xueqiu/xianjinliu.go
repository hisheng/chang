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

var Xianjinliu  Xianjinliu_
var XianjinliuRequest XianjinliuRequest_

type Xianjinliu_ struct {
	gorm.Model
	Report_type string
	Symbol string `sql:"type:varchar(20)"`
	Report_name string //1546185600000
	Report_date int  //"2018年报"
	Cash_received_of_sales_service float64 //销售商品、提供劳务收到的现金
	Cash_received_of_othr_oa float64 //收到其他与经营活动有关的现金
	Sub_total_of_ci_from_oa float64 //经营活动现金流入小计
	Goods_buy_and_service_cash_pay float64 //购买商品、接受劳务支付的现金
	Cash_paid_to_employee_etc float64 //支付给职工以及为职工支付的现金
	Payments_of_all_taxes float64 //支付的各项税费
	Othrcash_paid_relating_to_oa float64 //支付其他与经营活动有关的现金
	Sub_total_of_cos_from_oa float64 //经营活动现金流出小计
	Ncf_from_oa float64 //经营活动产生的现金流量净额
	Cash_received_of_dspsl_invest float64 //收回投资收到的现金
	Invest_income_cash_received float64 //取得投资收益收到的现金
	Net_cash_of_disposal_assets float64 //处置固定资产、无形资产和其他长期资产收回的现金净额
	Net_cash_of_disposal_branch float64 //处置子公司及其他营业单位收到的现金净额
	Cash_received_of_othr_ia float64 //收到其他与投资活动有关的现金
	Sub_total_of_ci_from_ia float64 //投资活动现金流入小计
	Cash_paid_for_assets float64 //购建固定资产、无形资产和其他长期资产支付的现金
	Invest_paid_cash float64 //投资支付的现金
	Net_cash_amt_from_branch float64 //取得子公司及其他营业单位支付的现金净额
	Othrcash_paid_relating_to_ia float64 //支付其他与投资活动有关的现金
	Sub_total_of_cos_from_ia float64 //投资活动现金流出小计
	Ncf_from_ia float64 //投资活动产生的现金流量净额
	Cash_received_of_absorb_invest float64 //吸收投资收到的现金
	Cash_received_from_investor float64 //其中：子公司吸收少数股东投资收到的现金
	Cash_received_of_borrowing float64 //取得借款收到的现金
	Cash_received_from_bond_issue float64 //发行债券收到的现金
	Cash_received_of_othr_fa float64 //收到其他与筹资活动有关的现金
	Sub_total_of_ci_from_fa float64 //筹资活动现金流入小计
	Cash_pay_for_debt float64 //偿还债务支付的现金
	Cash_paid_of_distribution float64 //分配股利、利润或偿付利息支付的现金
	Branch_paid_to_minority_holder float64 //其中：子公司支付给少数股东的股利
	Sub_total_of_cos_from_fa float64 //筹资活动现金流出小计
	Ncf_from_fa float64 //筹资活动产生的现金流量净额
	Effect_of_exchange_chg_on_cce float64 //汇率变动对现金及现金等价物的影响
	Net_increase_in_cce float64 //现金及现金等价物净增加额
	Initial_balance_of_cce float64 //加：期初现金及现金等价物余额
	Final_balance_of_cce float64 //期末现金及现金等价物余额
	GatherDay string
	XianjinliuPercent
}

func (x Xianjinliu_) getTableName() string {
	return "xianjinliu"
}

func (x Xianjinliu_) CreateTable()  {
	tableKey := x.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&x)
		db.DB.Table(tableKey).AddUniqueIndex(x.getTableName() +"symbol_gatherday", "symbol","gather_day")

	}
}




func (x Xianjinliu_)Add() Xianjinliu_{
	is := x.FindOne()
	if is.ID == 0 {
		db.DB.Table(x.getTableName()).Create(&x)
	}else {
		return is
	}
	return x
	//Exce 会释放连接池
}

func (x Xianjinliu_) FindOne() Xianjinliu_{
	dbn := db.DB.Table(x.getTableName()).Where("symbol = ?",x.Symbol).Where("gather_day = ?",x.GatherDay)
	dbn.First(&x)
	return x
}



type XianjinliuPercent struct {
	Cash_received_of_sales_service_percent float64 //销售商品、提供劳务收到的现金
	Cash_received_of_othr_oa_percent float64 //收到其他与经营活动有关的现金
	Sub_total_of_ci_from_oa_percent float64 //经营活动现金流入小计
	Goods_buy_and_service_cash_pay_percent float64 //购买商品、接受劳务支付的现金
	Cash_paid_to_employee_etc_percent float64 //支付给职工以及为职工支付的现金
	Payments_of_all_taxes_percent float64 //支付的各项税费
	Othrcash_paid_relating_to_oa_percent float64 //支付其他与经营活动有关的现金
	Sub_total_of_cos_from_oa_percent float64 //经营活动现金流出小计
	Ncf_from_oa_percent float64 //经营活动产生的现金流量净额
	Cash_received_of_dspsl_invest_percent float64 //收回投资收到的现金
	Invest_income_cash_received_percent float64 //取得投资收益收到的现金
	Net_cash_of_disposal_assets_percent float64 //处置固定资产、无形资产和其他长期资产收回的现金净额
	Net_cash_of_disposal_branch_percent float64 //处置子公司及其他营业单位收到的现金净额
	Cash_received_of_othr_ia_percent float64 //收到其他与投资活动有关的现金
	Sub_total_of_ci_from_ia_percent float64 //投资活动现金流入小计
	Cash_paid_for_assets_percent float64 //购建固定资产、无形资产和其他长期资产支付的现金
	Invest_paid_cash_percent float64 //投资支付的现金
	Net_cash_amt_from_branch_percent float64 //取得子公司及其他营业单位支付的现金净额
	Othrcash_paid_relating_to_ia_percent float64 //支付其他与投资活动有关的现金
	Sub_total_of_cos_from_ia_percent float64 //投资活动现金流出小计
	Ncf_from_ia_percent float64 //投资活动产生的现金流量净额
	Cash_received_of_absorb_invest_percent float64 //吸收投资收到的现金
	Cash_received_from_investor_percent float64 //其中：子公司吸收少数股东投资收到的现金
	Cash_received_of_borrowing_percent float64 //取得借款收到的现金
	Cash_received_from_bond_issue_percent float64 //发行债券收到的现金
	Cash_received_of_othr_fa_percent float64 //收到其他与筹资活动有关的现金
	Sub_total_of_ci_from_fa_percent float64 //筹资活动现金流入小计
	Cash_pay_for_debt_percent float64 //偿还债务支付的现金
	Cash_paid_of_distribution_percent float64 //分配股利、利润或偿付利息支付的现金
	Branch_paid_to_minority_holder_percent float64 //其中：子公司支付给少数股东的股利
	Sub_total_of_cos_from_fa_percent float64 //筹资活动现金流出小计
	Ncf_from_fa_percent float64 //筹资活动产生的现金流量净额
	Effect_of_exchange_chg_on_cce_percent float64 //汇率变动对现金及现金等价物的影响
	Net_increase_in_cce_percent float64 //现金及现金等价物净增加额
	Initial_balance_of_cce_percent float64 //加：期初现金及现金等价物余额
	Final_balance_of_cce_percent float64 //期末现金及现金等价物余额
}











type XianjinliuRequest_ Request_
func (request XianjinliuRequest_) initRequest(symbol string,count string) XianjinliuRequest_{
	request.SearchUrl  = "https://stock.xueqiu.com/v5/stock/finance/cn/cash_flow.json"

	request.SearchParms = url.Values{}
	request.SearchParms.Add("symbol",symbol)
	request.SearchParms.Add("type","all")
	request.SearchParms.Add("is_detail","true")
	request.SearchParms.Add("count",count)
	request.SearchParms.Add("timestamp","")
	return request
}





func (request XianjinliuRequest_) InitRun (symbol string)  {
	request.Run(symbol,"20")
}


func (request XianjinliuRequest_) Run (symbol string,count string)  {
	request = request.initRequest(symbol,count)
	fmt.Println(request.SearchParms.Get("type"))

	for i:= 1;i<=4;i++ {
		fmt.Println("XianjinliuRequest_ "+ symbol + " Q" + strconv.Itoa(i))

		request.SearchParms.Set("type","Q"+strconv.Itoa(i))  //Q1代表一季度
		request.RunGet()
	}

}

func (request XianjinliuRequest_) Update(symbol string)  {
	request.Run(symbol,"2")
}

func (request XianjinliuRequest_) RunGet()  {
	data := Get(request.SearchUrl,request.SearchParms)

	str:=[]byte(data)


	rs := XianjinliuJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	for _,stock := range rs.Data.List{
		xianjin := Xianjinliu_{}
		xianjin.Symbol = request.SearchParms.Get("symbol")
		xianjin.Report_name = stock.Report_name
		xianjin.Report_type = request.SearchParms.Get("type")

		val := stock.Report_date /1000
		int64_ := int64(val)
		tm := time.Unix(int64_ , 0)
		xianjin.GatherDay = tm.Format("2006-01-02")
		xianjin.Report_date = val


		xianjin.Ncf_from_fa = stock.Ncf_from_fa[0]
		xianjin.Ncf_from_fa_percent = stock.Ncf_from_fa[1]
		xianjin.Ncf_from_ia = stock.Ncf_from_ia[0]
		xianjin.Ncf_from_ia_percent = stock.Ncf_from_ia[1]
		xianjin.Ncf_from_fa = stock.Ncf_from_fa[0]
		xianjin.Ncf_from_fa_percent = stock.Ncf_from_fa[1]
		xianjin.Cash_received_of_othr_oa = stock.Cash_received_of_othr_oa[0]
		xianjin.Cash_received_of_othr_oa_percent = stock.Cash_received_of_othr_oa[1]
		xianjin.Sub_total_of_ci_from_oa = stock.Sub_total_of_ci_from_oa[0]
		xianjin.Sub_total_of_ci_from_oa_percent = stock.Sub_total_of_ci_from_oa[1]
		xianjin.Cash_paid_to_employee_etc = stock.Cash_paid_to_employee_etc[0]
		xianjin.Cash_paid_to_employee_etc_percent = stock.Cash_paid_to_employee_etc[1]
		xianjin.Payments_of_all_taxes = stock.Payments_of_all_taxes[0]
		xianjin.Payments_of_all_taxes_percent = stock.Payments_of_all_taxes[1]
		xianjin.Othrcash_paid_relating_to_oa = stock.Othrcash_paid_relating_to_oa[0]
		xianjin.Othrcash_paid_relating_to_oa_percent = stock.Othrcash_paid_relating_to_oa[1]
		xianjin.Sub_total_of_cos_from_oa = stock.Sub_total_of_cos_from_oa[0]
		xianjin.Sub_total_of_cos_from_oa_percent = stock.Sub_total_of_cos_from_oa[1]
		xianjin.Cash_received_of_dspsl_invest = stock.Cash_received_of_dspsl_invest[0]
		xianjin.Cash_received_of_dspsl_invest_percent = stock.Cash_received_of_dspsl_invest[1]
		xianjin.Invest_income_cash_received = stock.Invest_income_cash_received[0]
		xianjin.Invest_income_cash_received_percent = stock.Invest_income_cash_received[1]
		xianjin.Net_cash_of_disposal_assets = stock.Net_cash_of_disposal_assets[0]
		xianjin.Net_cash_of_disposal_assets_percent = stock.Net_cash_of_disposal_assets[1]
		xianjin.Net_cash_of_disposal_branch = stock.Net_cash_of_disposal_branch[0]
		xianjin.Net_cash_of_disposal_branch_percent = stock.Net_cash_of_disposal_branch[1]
		xianjin.Cash_received_of_othr_ia = stock.Cash_received_of_othr_ia[0]
		xianjin.Cash_received_of_othr_ia_percent = stock.Cash_received_of_othr_ia[1]
		xianjin.Sub_total_of_ci_from_ia = stock.Sub_total_of_ci_from_ia[0]
		xianjin.Sub_total_of_ci_from_ia_percent = stock.Sub_total_of_ci_from_ia[1]
		xianjin.Invest_paid_cash = stock.Invest_paid_cash[0]
		xianjin.Invest_paid_cash_percent = stock.Invest_paid_cash[1]
		xianjin.Cash_paid_for_assets = stock.Cash_paid_for_assets[0]
		xianjin.Cash_paid_for_assets_percent = stock.Cash_paid_for_assets[1]
		xianjin.Othrcash_paid_relating_to_ia = stock.Othrcash_paid_relating_to_ia[0]
		xianjin.Othrcash_paid_relating_to_ia_percent = stock.Othrcash_paid_relating_to_ia[1]
		xianjin.Sub_total_of_cos_from_ia = stock.Sub_total_of_cos_from_ia[0]
		xianjin.Sub_total_of_cos_from_ia_percent = stock.Sub_total_of_cos_from_ia[1]
		xianjin.Cash_received_of_absorb_invest = stock.Cash_received_of_absorb_invest[0]
		xianjin.Cash_received_of_absorb_invest_percent = stock.Cash_received_of_absorb_invest[1]
		xianjin.Cash_received_from_investor = stock.Cash_received_from_investor[0]
		xianjin.Cash_received_from_investor_percent = stock.Cash_received_from_investor[1]
		xianjin.Cash_received_from_bond_issue = stock.Cash_received_from_bond_issue[0]
		xianjin.Cash_received_from_bond_issue_percent = stock.Cash_received_from_bond_issue[1]
		xianjin.Cash_received_of_borrowing = stock.Cash_received_of_borrowing[0]
		xianjin.Cash_received_of_borrowing_percent = stock.Cash_received_of_borrowing[1]
		xianjin.Cash_received_of_othr_fa = stock.Cash_received_of_othr_fa[0]
		xianjin.Cash_received_of_othr_fa_percent = stock.Cash_received_of_othr_fa[1]
		xianjin.Sub_total_of_ci_from_fa = stock.Sub_total_of_ci_from_fa[0]
		xianjin.Sub_total_of_ci_from_fa_percent = stock.Sub_total_of_ci_from_fa[1]
		xianjin.Cash_pay_for_debt = stock.Cash_pay_for_debt[0]
		xianjin.Cash_pay_for_debt_percent = stock.Cash_pay_for_debt[1]
		xianjin.Cash_paid_of_distribution = stock.Cash_paid_of_distribution[0]
		xianjin.Cash_paid_of_distribution_percent = stock.Cash_paid_of_distribution[1]
		xianjin.Branch_paid_to_minority_holder = stock.Branch_paid_to_minority_holder[0]
		xianjin.Branch_paid_to_minority_holder_percent = stock.Branch_paid_to_minority_holder[1]
		xianjin.Sub_total_of_cos_from_fa = stock.Sub_total_of_cos_from_fa[0]
		xianjin.Sub_total_of_cos_from_fa_percent = stock.Sub_total_of_cos_from_fa[1]
		xianjin.Effect_of_exchange_chg_on_cce = stock.Effect_of_exchange_chg_on_cce[0]
		xianjin.Effect_of_exchange_chg_on_cce_percent = stock.Effect_of_exchange_chg_on_cce[1]
		xianjin.Net_increase_in_cce = stock.Net_increase_in_cce[0]
		xianjin.Net_increase_in_cce_percent = stock.Net_increase_in_cce[1]
		xianjin.Initial_balance_of_cce = stock.Initial_balance_of_cce[0]
		xianjin.Initial_balance_of_cce_percent = stock.Initial_balance_of_cce[1]
		xianjin.Final_balance_of_cce = stock.Final_balance_of_cce[0]
		xianjin.Final_balance_of_cce_percent = stock.Final_balance_of_cce[1]
		xianjin.Cash_received_of_sales_service = stock.Cash_received_of_sales_service[0]
		xianjin.Cash_received_of_sales_service_percent = stock.Cash_received_of_sales_service[1]
		xianjin.Goods_buy_and_service_cash_pay = stock.Goods_buy_and_service_cash_pay[0]
		xianjin.Goods_buy_and_service_cash_pay_percent = stock.Goods_buy_and_service_cash_pay[1]
		xianjin.Net_cash_amt_from_branch = stock.Net_cash_amt_from_branch[0]
		xianjin.Net_cash_amt_from_branch_percent = stock.Net_cash_amt_from_branch[1]
		xianjin.Add()
	}
}




type  XianjinliuJsonResponse struct{
	ErrorCode int `json:"error_code"`
	Data XianjinliuJsonData  `json:"data"`
}

type  XianjinliuJsonData struct {
	Symbol       string
	QuoteName    string
	Currency     string
	CurrencyName string
	OrgType      int
	List         []XianjinliuJsonDataItem
}

type  XianjinliuJsonDataItem struct{
	Report_type string
	Symbol string `sql:"type:varchar(20)"`
	Report_name string //1546185600000
	Report_date int  //"2018年报"
	Cash_received_of_sales_service [2]float64 //销售商品、提供劳务收到的现金
	Cash_received_of_othr_oa [2]float64 //收到其他与经营活动有关的现金
	Sub_total_of_ci_from_oa [2]float64 //经营活动现金流入小计
	Goods_buy_and_service_cash_pay [2]float64 //购买商品、接受劳务支付的现金
	Cash_paid_to_employee_etc [2]float64 //支付给职工以及为职工支付的现金
	Payments_of_all_taxes [2]float64 //支付的各项税费
	Othrcash_paid_relating_to_oa [2]float64 //支付其他与经营活动有关的现金
	Sub_total_of_cos_from_oa [2]float64 //经营活动现金流出小计
	Ncf_from_oa [2]float64 //经营活动产生的现金流量净额
	Cash_received_of_dspsl_invest [2]float64 //收回投资收到的现金
	Invest_income_cash_received [2]float64 //取得投资收益收到的现金
	Net_cash_of_disposal_assets [2]float64 //处置固定资产、无形资产和其他长期资产收回的现金净额
	Net_cash_of_disposal_branch [2]float64 //处置子公司及其他营业单位收到的现金净额
	Cash_received_of_othr_ia [2]float64 //收到其他与投资活动有关的现金
	Sub_total_of_ci_from_ia [2]float64 //投资活动现金流入小计
	Cash_paid_for_assets [2]float64 //购建固定资产、无形资产和其他长期资产支付的现金
	Invest_paid_cash [2]float64 //投资支付的现金
	Net_cash_amt_from_branch [2]float64 //取得子公司及其他营业单位支付的现金净额
	Othrcash_paid_relating_to_ia [2]float64 //支付其他与投资活动有关的现金
	Sub_total_of_cos_from_ia [2]float64 //投资活动现金流出小计
	Ncf_from_ia [2]float64 //投资活动产生的现金流量净额
	Cash_received_of_absorb_invest [2]float64 //吸收投资收到的现金
	Cash_received_from_investor [2]float64 //其中：子公司吸收少数股东投资收到的现金
	Cash_received_of_borrowing [2]float64 //取得借款收到的现金
	Cash_received_from_bond_issue [2]float64 //发行债券收到的现金
	Cash_received_of_othr_fa [2]float64 //收到其他与筹资活动有关的现金
	Sub_total_of_ci_from_fa [2]float64 //筹资活动现金流入小计
	Cash_pay_for_debt [2]float64 //偿还债务支付的现金
	Cash_paid_of_distribution [2]float64 //分配股利、利润或偿付利息支付的现金
	Branch_paid_to_minority_holder [2]float64 //其中：子公司支付给少数股东的股利
	Sub_total_of_cos_from_fa [2]float64 //筹资活动现金流出小计
	Ncf_from_fa [2]float64 //筹资活动产生的现金流量净额
	Effect_of_exchange_chg_on_cce [2]float64 //汇率变动对现金及现金等价物的影响
	Net_increase_in_cce [2]float64 //现金及现金等价物净增加额
	Initial_balance_of_cce [2]float64 //加：期初现金及现金等价物余额
	Final_balance_of_cce [2]float64 //期末现金及现金等价物余额
}
