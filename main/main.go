package main

import (
	"fmt"
	"github.com/hisheng/chang/conf"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/xueqiu"
	"github.com/hisheng/chang/xueqiu/caiwu"
)

func init()  {
	conf.Conf.Init()
	conf.XueqiuCookie.Init()
	db.GetDb()
}

func main()  {
	//xueqiuRun()
	//xueqiu.StockSummaryModel.Run()
	//xueqiu.StockChartModel.Run()
	//caiwu.LirunbiaoRequest.Run()

	//fmt.Println(xueqiu.Areas)
	xueqiu.SymbolRequest.Run()
	//caiwu.XianjinliuRequest.Run()
	//caiwu.ZichanfuzhaiRequest.Run()


	//fmt.Println(b.val(a,"name"))

}

//func (b BB)val(a AA,key string) BB {
//	b.key = a.key
//	return b
//}

func xueqiuRun()  {
	ss := xueqiu.Symbol.Gets()

	for _,s := range ss{
		fmt.Println(s)
		//caiwu.XianjinliuRequest.Run(s.Symbol)
		caiwu.LirunbiaoRequest.Run(s.Symbol)
		//caiwu.ZichanfuzhaiRequest.Run(s.Symbol)
		//time.Sleep(time.Second * 1)
	}
}