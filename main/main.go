package main

import (
	"fmt"
	"github.com/hisheng/chang/conf"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/xueqiu"
	"github.com/robfig/cron"
)

var (
	CRON  *cron.Cron
	everyMinuteFifteenSecond = "15 * * * * *"
	everyFourSecond = "*/4 * * * * *"
	everyDay = "4 21 17 * * *"
)


func init()  {
	conf.InitConf()
	db.GetDb()
}



func main()  {

	fmt.Println("chang start")
	//xueqiu.InitData()
	//xueqiu.StockSummaryModel.Run()
	//xueqiu.StockChartModel.Run()
	//caiwu.LirunbiaoRequest.Run()

	//fmt.Println(xueqiu.Areas)
	//xueqiu.SymbolRequest.Run()
	//caiwu.XianjinliuRequest.Run()
	//caiwu.ZichanfuzhaiRequest.Run()


	//fmt.Println(b.val(a,"name"))

	CRON = cron.New()

	// debug 开发环境 配置测试
	if conf.Conf.Debug {
		// 每5秒 打印一下时间
		//CRON.AddFunc(everyMinuteFifteenSecond, func() {fmt.Println(time.Now().Date())})
	}


	CRON.AddFunc(everyDay, func() {
		xueqiu.UpdateData()
	})



	CRON.Start()
	defer CRON.Stop()


	select {

	}

}

