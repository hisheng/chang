package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hisheng/jsoncached"
)

type TestController struct {
	Controller
}

func NewTestController(c *gin.Context) TestController {
	return TestController{Controller{
		Ctx: c,
	}}
}

func (c TestController) Detail() (int, gin.Negotiate) {
	guxiRecords, err := jsoncached.Get("guxiRecords")
	if err != nil {
		fmt.Println("没有值")
	}
	fmt.Println(guxiRecords)
	//ok := "i am ok"
	return c.Json(guxiRecords)
}
