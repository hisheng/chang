package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hisheng/chang/model"
	"github.com/hisheng/chang/repository"
	"github.com/hisheng/jsoncached"
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
	//1 从文件中读取
	guxiRecords, err := jsoncached.Get("guxiRecords")
	if err != nil {
		return c.Json([]struct{}{})
	}

	//2 返回
	var rs []model.GuxiRecord
	_ = json.Unmarshal(guxiRecords, &rs)
	return c.Json(rs)
}

func (c GuxiRecordController) Init() (int, gin.Negotiate) {
	//1 读取mysql
	rs := repository.NewGuxiRecordRepository().GetList()
	//2 写入文件
	rsjson, _ := json.Marshal(rs)
	ok := jsoncached.Set("guxiRecords", rsjson)

	return c.Json(ok)
}
