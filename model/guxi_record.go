package model

type GuxiRecord struct {
	Id            int
	Symbol        string `sql:"comment:'股票代码'"`
	SymbolName    string
	Moneytype     string
	Dividend      float64
	Shareholdings int `sql:"comment:'持有股份'"`
	tax           float64
	DividendHand  float64 `sql:"comment:'到手股息'"`
	GatherDay     string
}
