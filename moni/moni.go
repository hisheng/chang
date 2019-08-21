package moni

import (
	"github.com/hisheng/chang/db"
	"github.com/jinzhu/gorm"
)

var Moni Moni_

type Moni_ struct {
	gorm.Model
	Symbol,Name,Desc string
	Start_price float64
	Pe,Pb,Ps float64
	GatherDay string
	Group_code int
	Group_name string
}

func (m Moni_) getTableName() string {
	return "moni"
}

func (m Moni_) CreateTable()  {
	tableKey := m.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&m)
	}
}

func (m Moni_)Add() Moni_{
	is := m.FindOne()
	if is.ID == 0 {
		db.DB.Table(m.getTableName()).Create(&m)
	}else {
		return is
	}
	return m
	//Exce 会释放连接池
}

func (m Moni_) FindOne() Moni_{
	dbn := db.DB.Table(m.getTableName()).Where("symbol = ?",m.Symbol)
	dbn = dbn.Where("gather_day = ?",m.GatherDay)
	dbn = dbn.Where("group_code = ?",m.Group_code)
	dbn.First(&m)
	return m
}




