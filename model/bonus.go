package model

type Bonus struct {
	Id                   int    `sql:"primary_key;AUTO_INCREMENT",json:"-"`
	Symbol               string `sql:"comment:'股票代码'"`
	SymbolName           string
	DividendYear         string `json:"dividend_year"`
	AshareExDividendDate int    `json:"ashare_ex_dividend_date"`
	ExDividendDate       int    `json:"ex_dividend_date"`
	EquityDate           int    `json:"equity_date"`
	PlanExplain          string `json:"plan_explain"`
	PlanSchedule         string `json:"plan_schedule"`
	CancleDividendDate   int    `json:"cancle_dividend_date"`
	DividendDate         int    `json:"dividend_date"`
}
