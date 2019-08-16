package main

import (
	"fmt"
	"github.com/hisheng/chang/conf"
	"github.com/hisheng/chang/db"
)

func init()  {
	conf.InitConf()
	db.GetDb()
}

func main()  {

	fmt.Println("s")
	//xueqiu.InitData()
	//xueqiu.StockSummaryModel.Run()
	//xueqiu.StockChartModel.Run()
	//caiwu.LirunbiaoRequest.Run()

	//fmt.Println(xueqiu.Areas)
	//xueqiu.SymbolRequest.Run()
	//caiwu.XianjinliuRequest.Run()
	//caiwu.ZichanfuzhaiRequest.Run()


	//fmt.Println(b.val(a,"name"))

}

