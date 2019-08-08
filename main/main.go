package main

import (
	"github.com/hisheng/chang/conf"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/xueqiu"
)

func init()  {
	conf.Conf.Init()
	conf.XueqiuCookie.Init()
	db.GetDb()
}


func main()  {
	//xueqiu.StockSummaryModel.Run()
	//xueqiu.StockChartModel.Run()
	//caiwu.LirunbiaoRequest.Run()

	//fmt.Println(xueqiu.Areas)
	xueqiu.SymbolRequest.Run()

}