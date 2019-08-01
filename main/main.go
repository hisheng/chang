package main

import (
	"github.com/hisheng/chang/conf"
	"github.com/hisheng/chang/models"
	"github.com/hisheng/chang/xueqiu"
)

func init()  {
	conf.Conf.Init()
	models.GetDb()
}

func main()  {
	xueqiu.Request.Run()
}