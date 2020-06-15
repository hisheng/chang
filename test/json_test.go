package test

import (
	"encoding/json"
	"github.com/hisheng/chang/conf"
	"github.com/hisheng/chang/db"
	"github.com/jinzhu/gorm"
	"testing"
)

type TestJsonModel struct {
	gorm.Model
	Json string
}
type JsonName struct {
	Name string
	Age int
}

func (tj TestJsonModel) getTableName() string {
	return "testjson"
}

func (tj TestJsonModel) CreateTable()  {
	tableKey := tj.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&tj)
	}
}

func TestJson(t *testing.T) {
	conf.InitConf()
	db.GetDb()
	tj := TestJsonModel{}
	tj.CreateTable()
	t.Log("test")
	jn := JsonName{Name:"hisheng",Age:12}
	b, _ := json.Marshal(jn)
	tj.Json = string(b)
	db.DB.Table(tj.getTableName()).Create(&tj)
}
