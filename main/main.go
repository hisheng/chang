package main

import (
	"github.com/hisheng/chang/conf"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/xueqiu"
	"github.com/robfig/cron"
)

var (
//CRON Expression Format
//A cron expression represents a set of times, using 6 space-separated fields.
//
//Field name   | Mandatory? | Allowed values  | Allowed special characters
//----------   | ---------- | --------------  | --------------------------
//Seconds      | Yes        | 0-59            | * / , -
//Minutes      | Yes        | 0-59            | * / , -
//Hours        | Yes        | 0-23            | * / , -
//Day of month | Yes        | 1-31            | * / , - ?
//Month        | Yes        | 1-12 or JAN-DEC | * / , -
//Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
	CRON  *cron.Cron
	everyMinuteFifteenSecond = "15 * * * * *"
	everyFourSecond = "*/4 * * * * *"
	everyDay = "4 21 16 * * *"
)


func init()  {
	conf.InitConf()
	db.GetDb()
}



func main()  {


	//go moni.AddmoniGroup1("2019-08-20")
	//xueqiu.InitData()
	//xueqiu.StockSummaryModel.Run()
	//xueqiu.StockChartModel.Run()
	//caiwu.LirunbiaoRequest.Run()

	//fmt.Println(xueqiu.Areas)
	//xueqiu.SymbolRequest.Run()
	//caiwu.XianjinliuRequest.Run()
	//caiwu.ZichanfuzhaiRequest.Run()

	//go xueqiu.UpdateData()
	//select {
	//
	//}
	go xueqiu.UpdateData()
	//fmt.Println(b.val(a,"name"))

	CRON = cron.New()

	// debug 开发环境 配置测试
	if conf.Conf.Debug {
		// 每5秒 打印一下时间
		//CRON.AddFunc(everyMinuteFifteenSecond, func() {fmt.Println(time.Now().Date())})
	}


	CRON.AddFunc(everyDay, func() {
		go xueqiu.UpdateData()
	})



	CRON.Start()
	defer CRON.Stop()


	select {

	}

}

