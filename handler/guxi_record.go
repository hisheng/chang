package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hisheng/chang/controller"
)

func GuxiRecordList(c *gin.Context) {
	c.Negotiate(controller.NewGuxiRecordController(c).List())
}

func GuxiRecordInit(c *gin.Context) {
	c.Negotiate(controller.NewGuxiRecordController(c).Init())
}
