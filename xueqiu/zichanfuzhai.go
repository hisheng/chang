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

var ZichanfuzhaiRequest  ZichanfuzhaiRequest_
var Zichanfuzhai Zichanfuzhai_


type Zichanfuzhai_ struct {
	gorm.Model
	Report_type string
	Symbol string `sql:"type:varchar(20)"`
	Report_name string //1546185600000
	Report_date int  //"2018年报"
	GatherDay string
	Currency_funds float64 //货币资金
	Ar_and_br float64 //应收票据及应收账款
	Bills_receivable float64 //其中：应收票据
	Account_receivable float64 //应收账款
	Pre_payment float64 //预付款项
	Othr_receivables float64 //其他应收款
	Inventory float64 //存货
	Othr_current_assets float64 //其他流动资产
	Total_current_assets float64 //流动资产合计
	Lt_receivable float64 //长期应收款
	Lt_equity_invest float64 //长期股权投资
	Other_illiquid_fnncl_assets float64 //其他非流动金融资产
	Invest_property float64 //投资性房地产
	Fixed_asset_sum float64 //固定资产合计
	Fixed_asset float64 //其中：固定资产
	Intangible_assets float64 //无形资产
	Lt_deferred_expense float64 //长期待摊费用
	Dt_assets float64 //递延所得税资产
	Othr_noncurrent_assets float64 //其他非流动资产
	Total_noncurrent_assets float64 //非流动资产合计
	Total_assets float64 //资产合计
	St_loan float64 //短期借款
	Derivative_fnncl_liab float64 //衍生金融负债
	Bp_and_ap float64 //应付票据及应付账款
	Bill_payable float64 //应付票据
	Accounts_payable float64 //应付账款
	Pre_receivable float64 //预收款项
	Contract_liabilities float64 //合同负债
	Payroll_payable float64 //应付职工薪酬
	Tax_payable float64 //应交税费
	Interest_payable float64 //应付利息
	Dividend_payable float64 //应付股利
	Othr_payables float64 //其他应付款
	Noncurrent_liab_due_in1y float64 //一年内到期的非流动负债
	Othr_current_liab float64 //其他流动负债
	Total_current_liab float64 //流动负债合计
	Lt_loan float64 //长期借款
	Bond_payable float64 //应付债券
	Lt_payable float64 //长期应付款
	Dt_liab float64 //递延所得税负债
	Total_noncurrent_liab float64 //非流动负债合计
	Total_liab float64 //负债合计
	Shares float64 //实收资本(或股本)
	Othr_equity_instruments float64 //其他权益工具
	Perpetual_bond float64 //永续债
	Capital_reserve float64 //资本公积
	Treasury_stock float64 //减：库存股
	Othr_compre_income float64 //其他综合收益
	Earned_surplus float64 //盈余公积
	Undstrbtd_profit float64 //未分配利润
	Total_quity_atsopc float64 //归属于母公司股东权益合计
	Minority_equity float64 //少数股东权益
	Total_holders_equity float64 //股东权益合计
	Total_liab_and_holders_equity float64 //负债和股东权益总计
	ZichanfuzhaiPercent
}


type ZichanfuzhaiPercent struct {
	Currency_funds_percent float64 //货币资金
	Ar_and_br_percent float64 //应收票据及应收账款
	Bills_receivable_percent float64 //其中：应收票据
	Account_receivable_percent float64 //应收账款
	Pre_payment_percent float64 //预付款项
	Othr_receivables_percent float64 //其他应收款
	Inventory_percent float64 //存货
	Othr_current_assets_percent float64 //其他流动资产
	Total_current_assets_percent float64 //流动资产合计
	Lt_receivable_percent float64 //长期应收款
	Lt_equity_invest_percent float64 //长期股权投资
	Other_illiquid_fnncl_assets_percent float64 //其他非流动金融资产
	Invest_property_percent float64 //投资性房地产
	Fixed_asset_sum_percent float64 //固定资产合计
	Fixed_asset_percent float64 //其中：固定资产
	Intangible_assets_percent float64 //无形资产
	Lt_deferred_expense_percent float64 //长期待摊费用
	Dt_assets_percent float64 //递延所得税资产
	Othr_noncurrent_assets_percent float64 //其他非流动资产
	Total_noncurrent_assets_percent float64 //非流动资产合计
	Total_assets_percent float64 //资产合计
	St_loan_percent float64 //短期借款
	Derivative_fnncl_liab_percent float64 //衍生金融负债
	Bp_and_ap_percent float64 //应付票据及应付账款
	Bill_payable_percent float64 //应付票据
	Accounts_payable_percent float64 //应付账款
	Pre_receivable_percent float64 //预收款项
	Contract_liabilities_percent float64 //合同负债
	Payroll_payable_percent float64 //应付职工薪酬
	Tax_payable_percent float64 //应交税费
	Interest_payable_percent float64 //应付利息
	Dividend_payable_percent float64 //应付股利
	Othr_payables_percent float64 //其他应付款
	Noncurrent_liab_due_in1y_percent float64 //一年内到期的非流动负债
	Othr_current_liab_percent float64 //其他流动负债
	Total_current_liab_percent float64 //流动负债合计
	Lt_loan_percent float64 //长期借款
	Bond_payable_percent float64 //应付债券
	Lt_payable_percent float64 //长期应付款
	Dt_liab_percent float64 //递延所得税负债
	Total_noncurrent_liab_percent float64 //非流动负债合计
	Total_liab_percent float64 //负债合计
	Shares_percent float64 //实收资本(或股本)
	Othr_equity_instruments_percent float64 //其他权益工具
	Perpetual_bond_percent float64 //永续债
	Capital_reserve_percent float64 //资本公积
	Treasury_stock_percent float64 //减：库存股
	Othr_compre_income_percent float64 //其他综合收益
	Earned_surplus_percent float64 //盈余公积
	Undstrbtd_profit_percent float64 //未分配利润
	Total_quity_atsopc_percent float64 //归属于母公司股东权益合计
	Minority_equity_percent float64 //少数股东权益
	Total_holders_equity_percent float64 //股东权益合计
	Total_liab_and_holders_equity_percent float64 //负债和股东权益总计
}


func (z Zichanfuzhai_) getTableName() string {
	return "zichanfuzhai"
}

func (z Zichanfuzhai_) CreateTable()  {
	tableKey := z.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&z)
		db.DB.Table(tableKey).AddUniqueIndex(z.getTableName() +"symbol_gatherday", "symbol","gather_day")
	}
}




func (z Zichanfuzhai_)Add() Zichanfuzhai_{
	is := z.FindOne()
	if is.ID == 0 {
		db.DB.Table(z.getTableName()).Create(&z)
	}else {
		return is
	}
	return z
	//Exce 会释放连接池
}

func (z Zichanfuzhai_) FindOne() Zichanfuzhai_{
	dbn := db.DB.Table(z.getTableName()).Where("symbol = ?",z.Symbol).Where("gather_day = ?",z.GatherDay)
	dbn.First(&z)
	return z
}





type ZichanfuzhaiRequest_ Request_
func (request ZichanfuzhaiRequest_) initRequest(symbol string) ZichanfuzhaiRequest_{
	request.SearchUrl  = "https://stock.xueqiu.com/v5/stock/finance/cn/balance.json"

	request.SearchParms = url.Values{}
	request.SearchParms.Add("symbol",symbol)
	request.SearchParms.Add("type","all")
	request.SearchParms.Add("is_detail","true")
	request.SearchParms.Add("count","20")
	request.SearchParms.Add("timestamp","")
	return request
}






func (request ZichanfuzhaiRequest_) Run (symbol string)  {

	request = request.initRequest(symbol)
	fmt.Println(request.SearchParms.Get("type"))

	for i:= 1;i<=4;i++ {
		fmt.Println("ZichanfuzhaiRequest_ "+ symbol + " Q" + strconv.Itoa(i))

		request.SearchParms.Set("type","Q"+strconv.Itoa(i))  //Q1代表一季度
		//fmt.Println(request)
		request.RunGet()
	}

}


func (request ZichanfuzhaiRequest_) RunGet()  {
	data := Get(request.SearchUrl,request.SearchParms)

	str:=[]byte(data)


	rs := ZichanfuzhaiJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	for _,stock := range rs.Data.List{

		zichanfuzhai := Zichanfuzhai_{}
		zichanfuzhai.Symbol = request.SearchParms.Get("symbol")
		zichanfuzhai.Report_name = stock.Report_name
		zichanfuzhai.Report_type = request.SearchParms.Get("type")

		val := stock.Report_date /1000
		int64_ := int64(val)
		tm := time.Unix(int64_ , 0)
		zichanfuzhai.GatherDay = tm.Format("2006-01-02")
		zichanfuzhai.Report_date = val






		zichanfuzhai.Currency_funds = stock.Currency_funds[0]
		zichanfuzhai.Currency_funds_percent = stock.Currency_funds[1]
		zichanfuzhai.Ar_and_br = stock.Ar_and_br[0]
		zichanfuzhai.Ar_and_br_percent = stock.Ar_and_br[1]
		zichanfuzhai.Bills_receivable = stock.Bills_receivable[0]
		zichanfuzhai.Bills_receivable_percent = stock.Bills_receivable[1]
		zichanfuzhai.Account_receivable = stock.Account_receivable[0]
		zichanfuzhai.Account_receivable_percent = stock.Account_receivable[1]
		zichanfuzhai.Pre_payment = stock.Pre_payment[0]
		zichanfuzhai.Pre_payment_percent = stock.Pre_payment[1]
		zichanfuzhai.Othr_receivables = stock.Othr_receivables[0]
		zichanfuzhai.Othr_receivables_percent = stock.Othr_receivables[1]
		zichanfuzhai.Inventory = stock.Inventory[0]
		zichanfuzhai.Inventory_percent = stock.Inventory[1]
		zichanfuzhai.Othr_current_assets = stock.Othr_current_assets[0]
		zichanfuzhai.Othr_current_assets_percent = stock.Othr_current_assets[1]
		zichanfuzhai.Total_current_assets = stock.Total_current_assets[0]
		zichanfuzhai.Total_current_assets_percent = stock.Total_current_assets[1]
		zichanfuzhai.Lt_receivable = stock.Lt_receivable[0]
		zichanfuzhai.Lt_receivable_percent = stock.Lt_receivable[1]
		zichanfuzhai.Lt_equity_invest = stock.Lt_equity_invest[0]
		zichanfuzhai.Lt_equity_invest_percent = stock.Lt_equity_invest[1]
		zichanfuzhai.Other_illiquid_fnncl_assets = stock.Other_illiquid_fnncl_assets[0]
		zichanfuzhai.Other_illiquid_fnncl_assets_percent = stock.Other_illiquid_fnncl_assets[1]
		zichanfuzhai.Invest_property = stock.Invest_property[0]
		zichanfuzhai.Invest_property_percent = stock.Invest_property[1]
		zichanfuzhai.Fixed_asset_sum = stock.Fixed_asset_sum[0]
		zichanfuzhai.Fixed_asset_sum_percent = stock.Fixed_asset_sum[1]
		zichanfuzhai.Fixed_asset = stock.Fixed_asset[0]
		zichanfuzhai.Fixed_asset_percent = stock.Fixed_asset[1]
		zichanfuzhai.Intangible_assets = stock.Intangible_assets[0]
		zichanfuzhai.Intangible_assets_percent = stock.Intangible_assets[1]
		zichanfuzhai.Lt_deferred_expense = stock.Lt_deferred_expense[0]
		zichanfuzhai.Lt_deferred_expense_percent = stock.Lt_deferred_expense[1]
		zichanfuzhai.Dt_assets = stock.Dt_assets[0]
		zichanfuzhai.Dt_assets_percent = stock.Dt_assets[1]
		zichanfuzhai.Othr_noncurrent_assets = stock.Othr_noncurrent_assets[0]
		zichanfuzhai.Othr_noncurrent_assets_percent = stock.Othr_noncurrent_assets[1]
		zichanfuzhai.Total_noncurrent_assets = stock.Total_noncurrent_assets[0]
		zichanfuzhai.Total_noncurrent_assets_percent = stock.Total_noncurrent_assets[1]
		zichanfuzhai.Total_assets = stock.Total_assets[0]
		zichanfuzhai.Total_assets_percent = stock.Total_assets[1]
		zichanfuzhai.St_loan = stock.St_loan[0]
		zichanfuzhai.St_loan_percent = stock.St_loan[1]
		zichanfuzhai.Derivative_fnncl_liab = stock.Derivative_fnncl_liab[0]
		zichanfuzhai.Derivative_fnncl_liab_percent = stock.Derivative_fnncl_liab[1]
		zichanfuzhai.Bp_and_ap = stock.Bp_and_ap[0]
		zichanfuzhai.Bp_and_ap_percent = stock.Bp_and_ap[1]
		zichanfuzhai.Bill_payable = stock.Bill_payable[0]
		zichanfuzhai.Bill_payable_percent = stock.Bill_payable[1]
		zichanfuzhai.Accounts_payable = stock.Accounts_payable[0]
		zichanfuzhai.Accounts_payable_percent = stock.Accounts_payable[1]
		zichanfuzhai.Pre_receivable = stock.Pre_receivable[0]
		zichanfuzhai.Pre_receivable_percent = stock.Pre_receivable[1]
		zichanfuzhai.Contract_liabilities = stock.Contract_liabilities[0]
		zichanfuzhai.Contract_liabilities_percent = stock.Contract_liabilities[1]
		zichanfuzhai.Payroll_payable = stock.Payroll_payable[0]
		zichanfuzhai.Payroll_payable_percent = stock.Payroll_payable[1]
		zichanfuzhai.Tax_payable = stock.Tax_payable[0]
		zichanfuzhai.Tax_payable_percent= stock.Tax_payable[1]
		zichanfuzhai.Interest_payable = stock.Interest_payable[0]
		zichanfuzhai.Interest_payable_percent = stock.Interest_payable[1]
		zichanfuzhai.Dividend_payable = stock.Dividend_payable[0]
		zichanfuzhai.Dividend_payable_percent = stock.Dividend_payable[1]
		zichanfuzhai.Othr_payables = stock.Othr_payables[0]
		zichanfuzhai.Othr_payables_percent = stock.Othr_payables[1]
		zichanfuzhai.Noncurrent_liab_due_in1y = stock.Noncurrent_liab_due_in1y[0]
		zichanfuzhai.Noncurrent_liab_due_in1y_percent = stock.Noncurrent_liab_due_in1y[1]
		zichanfuzhai.Othr_current_liab = stock.Othr_current_liab[0]
		zichanfuzhai.Othr_current_liab_percent = stock.Othr_current_liab[1]
		zichanfuzhai.Total_current_liab = stock.Total_current_liab[0]
		zichanfuzhai.Total_current_liab_percent = stock.Total_current_liab[1]
		zichanfuzhai.Lt_loan = stock.Lt_loan[0]
		zichanfuzhai.Lt_loan_percent = stock.Lt_loan[1]
		zichanfuzhai.Bond_payable = stock.Bond_payable[0]
		zichanfuzhai.Bond_payable_percent = stock.Bond_payable[1]
		zichanfuzhai.Lt_payable = stock.Lt_payable[0]
		zichanfuzhai.Lt_payable_percent = stock.Lt_payable[1]
		zichanfuzhai.Dt_liab = stock.Dt_liab[0]
		zichanfuzhai.Dt_liab_percent = stock.Dt_liab[1]
		zichanfuzhai.Total_noncurrent_liab = stock.Total_noncurrent_liab[0]
		zichanfuzhai.Total_noncurrent_liab_percent = stock.Total_noncurrent_liab[1]
		zichanfuzhai.Total_liab = stock.Total_liab[0]
		zichanfuzhai.Total_liab_percent = stock.Total_liab[1]
		zichanfuzhai.Shares = stock.Shares[0]
		zichanfuzhai.Shares_percent = stock.Shares[1]
		zichanfuzhai.Othr_equity_instruments = stock.Othr_equity_instruments[0]
		zichanfuzhai.Othr_equity_instruments_percent = stock.Othr_equity_instruments[1]
		zichanfuzhai.Perpetual_bond = stock.Perpetual_bond[0]
		zichanfuzhai.Perpetual_bond_percent = stock.Perpetual_bond[1]
		zichanfuzhai.Capital_reserve = stock.Capital_reserve[0]
		zichanfuzhai.Capital_reserve_percent = stock.Capital_reserve[1]
		zichanfuzhai.Treasury_stock = stock.Treasury_stock[0]
		zichanfuzhai.Treasury_stock_percent = stock.Treasury_stock[1]
		zichanfuzhai.Othr_compre_income = stock.Othr_compre_income[0]
		zichanfuzhai.Othr_compre_income_percent = stock.Othr_compre_income[1]
		zichanfuzhai.Earned_surplus = stock.Earned_surplus[0]
		zichanfuzhai.Earned_surplus_percent = stock.Earned_surplus[1]
		zichanfuzhai.Undstrbtd_profit = stock.Undstrbtd_profit[0]
		zichanfuzhai.Undstrbtd_profit_percent = stock.Undstrbtd_profit[1]
		zichanfuzhai.Total_quity_atsopc = stock.Total_quity_atsopc[0]
		zichanfuzhai.Total_quity_atsopc_percent = stock.Total_quity_atsopc[1]
		zichanfuzhai.Minority_equity = stock.Minority_equity[0]
		zichanfuzhai.Minority_equity_percent = stock.Minority_equity[1]
		zichanfuzhai.Total_holders_equity = stock.Total_holders_equity[0]
		zichanfuzhai.Total_holders_equity_percent = stock.Total_holders_equity[1]
		zichanfuzhai.Total_liab_and_holders_equity = stock.Total_liab_and_holders_equity[0]
		zichanfuzhai.Total_liab_and_holders_equity_percent = stock.Total_liab_and_holders_equity[1]

		zichanfuzhai.Add()
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
	Currency_funds [2]float64 //货币资金
	Ar_and_br [2]float64 //应收票据及应收账款
	Bills_receivable [2]float64 //其中：应收票据
	Account_receivable [2]float64 //应收账款
	Pre_payment [2]float64 //预付款项
	Othr_receivables [2]float64 //其他应收款
	Inventory [2]float64 //存货
	Othr_current_assets [2]float64 //其他流动资产
	Total_current_assets [2]float64 //流动资产合计
	Lt_receivable [2]float64 //长期应收款
	Lt_equity_invest [2]float64 //长期股权投资
	Other_illiquid_fnncl_assets [2]float64 //其他非流动金融资产
	Invest_property [2]float64 //投资性房地产
	Fixed_asset_sum [2]float64 //固定资产合计
	Fixed_asset [2]float64 //其中：固定资产
	Intangible_assets [2]float64 //无形资产
	Lt_deferred_expense [2]float64 //长期待摊费用
	Dt_assets [2]float64 //递延所得税资产
	Othr_noncurrent_assets [2]float64 //其他非流动资产
	Total_noncurrent_assets [2]float64 //非流动资产合计
	Total_assets [2]float64 //资产合计
	St_loan [2]float64 //短期借款
	Derivative_fnncl_liab [2]float64 //衍生金融负债
	Bp_and_ap [2]float64 //应付票据及应付账款
	Bill_payable [2]float64 //应付票据
	Accounts_payable [2]float64 //应付账款
	Pre_receivable [2]float64 //预收款项
	Contract_liabilities [2]float64 //合同负债
	Payroll_payable [2]float64 //应付职工薪酬
	Tax_payable [2]float64 //应交税费
	Interest_payable [2]float64 //应付利息
	Dividend_payable [2]float64 //应付股利
	Othr_payables [2]float64 //其他应付款
	Noncurrent_liab_due_in1y [2]float64 //一年内到期的非流动负债
	Othr_current_liab [2]float64 //其他流动负债
	Total_current_liab [2]float64 //流动负债合计
	Lt_loan [2]float64 //长期借款
	Bond_payable [2]float64 //应付债券
	Lt_payable [2]float64 //长期应付款
	Dt_liab [2]float64 //递延所得税负债
	Total_noncurrent_liab [2]float64 //非流动负债合计
	Total_liab [2]float64 //负债合计
	Shares [2]float64 //实收资本(或股本)
	Othr_equity_instruments [2]float64 //其他权益工具
	Perpetual_bond [2]float64 //永续债
	Capital_reserve [2]float64 //资本公积
	Treasury_stock [2]float64 //减：库存股
	Othr_compre_income [2]float64 //其他综合收益
	Earned_surplus [2]float64 //盈余公积
	Undstrbtd_profit [2]float64 //未分配利润
	Total_quity_atsopc [2]float64 //归属于母公司股东权益合计
	Minority_equity [2]float64 //少数股东权益
	Total_holders_equity [2]float64 //股东权益合计
	Total_liab_and_holders_equity [2]float64 //负债和股东权益总计
}
