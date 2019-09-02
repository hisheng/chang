package main

import (
	"github.com/hisheng/chang/conf"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/moni"
)

func init()  {
	conf.InitConf()
	db.GetDb()
}

func main()  {
	moni.Init()
	//xueqiu.InitData()
}
