package xueqiu

import (
	"fmt"
	"time"
)

func InitData()  {
	createTable()

	ss := Symbol.Gets()

	for _,s := range ss{
		fmt.Println(s)
		//go XianjinliuRequest.Run(s.Symbol)
		//go LirunbiaoRequest.Run(s.Symbol)
		//go ZichanfuzhaiRequest.Run(s.Symbol)
		go StockChartRequest.Run(s.Symbol)
		time.Sleep(time.Millisecond * 1000)
	}
}
