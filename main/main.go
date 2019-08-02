package main

import (
	"github.com/hisheng/chang/conf"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/xueqiu"
)

func init()  {
	conf.Conf.Init()
	db.GetDb()
}

func main()  {
	xueqiu.StockSummaryModel.Run()
}