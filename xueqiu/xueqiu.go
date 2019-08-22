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
		go XianjinliuRequest.Run(s.Symbol)
		go LirunbiaoRequest.Run(s.Symbol)
		go ZichanfuzhaiRequest.Run(s.Symbol)
		StockChartRequest.Run(s.Symbol,"")
	}
}

func UpdateData()  {
	createTable()
	SymbolRequest.Run()

	ss := Symbol.Gets()
	for _,s := range ss{
		fmt.Println(s)
		StockChartRequest.Run(s.Symbol,"10")
		go StockQuoteRequest.Run(s.Symbol)
	}
}