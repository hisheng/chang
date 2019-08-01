package models

import (
	"github.com/hisheng/chang/models"
	"github.com/jinzhu/gorm"
)

type StockSummary struct {
	gorm.Model
	Symbol string `json:"symbol"`
	Pettm float32 `json:"pettm"`// pe ttm
	Npay float32 `json:"npay"`// 净利润同比增
	Current float32 `json:"current"`//目前价格
	Name string `json:"name"`//名
	Oiy float32 `json:"oiy"`// 营业收入同比增长
	OiyPe float32    //  营业收入同比增长 / Pettm
}

func (stockSummary *StockSummary) InitOiyPe ()  {
	stockSummary.OiyPe = stockSummary.Oiy / stockSummary.Pettm
}


func (s StockSummary) createTable()  {
	tableKey := s.getTableName()
	if !models.DB.HasTable(tableKey) {
		models.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&s)
	}
}

func (s StockSummary) getTableName() string {
	tableKey := "stock_summary"
	return tableKey
}
