package xueqiu

import "github.com/hisheng/chang/xueqiu/models"

type  XueqiuJsonResponse struct{
	ErrorCode int `json:"error_code"`
	Data XueqiuJsonData
}
type XueqiuJsonData struct {
	Count int `json:"count"`
	List []models.StockSummary `json:"list"`
}



