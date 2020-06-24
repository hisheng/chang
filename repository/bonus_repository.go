package repository

import (
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/model"
)

type BonusRepository struct {
}

func NewBonusRepository() *BonusRepository {
	return &BonusRepository{}
}

func (r BonusRepository) getTableName() string {
	return "bonus"
}

func (r BonusRepository) CreateTable() {
	tableKey := r.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Bonus{})
	}
}

func (r BonusRepository) GetList() []model.Bonus {
	var bonus []model.Bonus
	db.DB.Table(r.getTableName()).Find(&bonus)
	return bonus
}

func (r BonusRepository) Add(bonus *model.Bonus) bool {
	b := r.FindOne(bonus.Symbol, bonus.DividendDate)
	if b.Id > 0 {
		return true
	}
	if err := db.DB.Create(bonus).Error; err != nil {
		//fmt.Println("插入失败", err)
		return false
	}
	return true
}

func (r BonusRepository) Update(bonus *model.Bonus) bool {
	b := r.FindOne(bonus.Symbol, bonus.DividendDate)
	if b.Id > 0 {
		//update
		if b.DividendDate == 0 {
			b.DividendDate = bonus.DividendDate
			db.DB.Save(&b)
		}
		return true
	}
	//insert
	if err := db.DB.Create(bonus).Error; err != nil {
		return false
	}
	return true
}

func (r BonusRepository) FindOne(symbol string, dividendDate int) model.Bonus {
	bonus := model.Bonus{}
	db.DB.Table(r.getTableName()).Where("symbol = ? and dividend_date = ?", symbol, dividendDate).First(&bonus)
	return bonus
}
