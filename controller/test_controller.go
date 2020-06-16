package controller

import (
	"github.com/gin-gonic/gin"
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
	ok := "i am ok"
	return c.Json(ok)
}
