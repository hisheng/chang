package xueqiu

import (
	"github.com/hisheng/chang/db"
	"github.com/jinzhu/gorm"
)

type GuxiRecord struct {
	gorm.Model
	Id            int
	Symbol        string `sql:"comment:'股票代码'"`
	SymbolName    string
	Moneytype     string
	Dividend      float64
	Shareholdings int `sql:"comment:'持有股份'"`
	tax           float64
	DividendHand  float64 `sql:"comment:'到手股息'"`
	GatherDay     string
}

func (guxi GuxiRecord) getTableName() string {
	return "guxi_record"
}

func (guxi GuxiRecord) CreateTable() {
	tableKey := guxi.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&guxi)
	}
}

func (guxi GuxiRecord) getList() []GuxiRecord {
	var guxis []GuxiRecord
	db.DB.Table(guxi.getTableName()).Find(&guxis)
	return guxis
}
