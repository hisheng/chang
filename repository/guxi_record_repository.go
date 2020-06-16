package repository

import (
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/model"
)

type GuxiRecordRepository struct {
}

func NewGuxiRecordRepository() GuxiRecordRepository {
	return GuxiRecordRepository{}
}

func (guxi GuxiRecordRepository) getTableName() string {
	return "guxi_record"
}

func (guxi GuxiRecordRepository) CreateTable() {
	tableKey := guxi.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.GuxiRecord{})
	}
}

func (guxi GuxiRecordRepository) GetList() []model.GuxiRecord {
	var guxis []model.GuxiRecord
	db.DB.Table(guxi.getTableName()).Find(&guxis)
	return guxis
}
