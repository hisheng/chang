package xueqiu

import (
	"fmt"
)

func InitData()  {
	createTable()
	SymbolRequest.Run()
	ss := Symbol.Gets()

	for _,s := range ss{
		fmt.Println(s)
		go XianjinliuRequest.InitRun(s.Symbol)
		go LirunbiaoRequest.InitRun(s.Symbol)
		go ZichanfuzhaiRequest.InitRun(s.Symbol)
		StockChartRequest.Run(s.Symbol,"")
	}
}

func UpdateData()  {
	createTable()
	
	updateSymbol()

	ss := Symbol.Gets()
	for _,s := range ss{
		fmt.Println(s)
		updateStock(s.Symbol)
		UpdateCaiwuBaobiao(s.Symbol)
	}
}

func updateSymbol()  {
	SymbolRequest.Run()
}

func updateStock(symbol string)  {
	StockChartRequest.Run(symbol,"10")
	go StockQuoteRequest.Run(symbol)
}

func UpdateCaiwuBaobiao(symbol string)  {
	XianjinliuRequest.Update(symbol)
	LirunbiaoRequest.Update(symbol)
	ZichanfuzhaiRequest.Upadte(symbol)
}