package xueqiu

import "github.com/jinzhu/gorm"

/*
xueqiu stock chart json
*/

type StockChart struct {
	gorm.Model
	Timestamp string
	Volume float32
	Open float32
	High float32
	Low float32
	Close float32
	Chg float32
	Percent float32
	Turnoverrate float32
	Amount float32
	VolumePost float32 `json:"volume_post"`
	AmountPost float32 `json:"amount_post"`
	Pe float32
	Pb float32
	Ps float32
	Pcf float32
	Market_capital float32
	Balance float32
}

//chart json
type  XueqiuChartJsonResponse struct{
	ErrorCode int `json:"error_code"`
	Data XueqiuChartJsonData
}

type  XueqiuChartJsonData struct{
	symbol string
	column string
	item []StockChart
}
