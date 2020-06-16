package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hisheng/chang/repository"
)

type GuxiRecordController struct {
	Controller
}

func NewGuxiRecordController(c *gin.Context) GuxiRecordController {
	return GuxiRecordController{Controller{
		Ctx: c,
	}}
}

func (c GuxiRecordController) List() (int, gin.Negotiate) {
	return c.Json(repository.NewGuxiRecordRepository().GetList())
}
