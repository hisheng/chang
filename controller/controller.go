package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const codeOk = 200

type Controller struct {
	Ctx *gin.Context
}

func (c Controller) Json(obj interface{}) (int, gin.Negotiate) {
	return codeOk, gin.Negotiate{
		Offered: []string{binding.MIMEJSON},
		Data:    obj,
	}
}
