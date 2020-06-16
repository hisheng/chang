package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hisheng/chang/model"
	"github.com/hisheng/jsoncached"
)

type AreaController struct {
	Controller
}

func NewAreaController(c *gin.Context) AreaController {
	return AreaController{Controller{
		Ctx: c,
	}}
}

func (c AreaController) List() (int, gin.Negotiate) {
	var areaMap map[int]string
	_ = json.Unmarshal([]byte(jsoncached.Get("area")), &areaMap)
	return c.Json(areaMap)
}

func (c AreaController) Init() (int, gin.Negotiate) {
	//1 写入文件
	areaJson, _ := json.Marshal(model.Areas)
	ok := jsoncached.SetByte("area", areaJson)
	return c.Json(ok)
}
