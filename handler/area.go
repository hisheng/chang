package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hisheng/chang/controller"
)

func Area(c *gin.Context) {
	c.Negotiate(controller.NewAreaController(c).List())
}

func AreaInit(c *gin.Context) {
	c.Negotiate(controller.NewAreaController(c).Init())
}
