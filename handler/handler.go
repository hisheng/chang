package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hisheng/chang/controller"
)

func TestDetail(c *gin.Context) {
	c.Negotiate(controller.NewTestController(c).Detail())
}

func GuxiRecord(c *gin.Context) {
	c.Negotiate(controller.NewGuxiRecordController(c).List())
}
