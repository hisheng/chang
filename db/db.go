package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hisheng/chang/conf"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// gorm.Model definition
type Model struct {
	Id        int
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;"`
	DeletedAt time.Time `gorm:"column:deleted_at;"`
}

var (
	DB       *gorm.DB
	err      error
	ReportDb *gorm.DB
)

func GetDb() {
	initCreateDb()
	initChangDb()
	initReportDb()
	//initTestDb()
	//return db
}

func initCreateDb() {
	db, err := sql.Open("mysql", conf.Conf.MysqlConf.User+
		":"+conf.Conf.MysqlConf.Password+
		"@tcp("+conf.Conf.MysqlConf.Ip+
		conf.Conf.MysqlConf.Port+")/"+
		"?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("create database IF NOT EXISTS chang character set utf8mb4 collate utf8mb4_general_ci")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create database IF NOT EXISTS chang_report character set utf8mb4 collate utf8mb4_general_ci")
	if err != nil {
		panic(err)
	}
}

func initChangDb() {
	fmt.Println(conf.Conf)
	DB, err = gorm.Open(
		"mysql",
		conf.Conf.MysqlConf.User+
			":"+conf.Conf.MysqlConf.Password+
			"@tcp("+conf.Conf.MysqlConf.Ip+
			conf.Conf.MysqlConf.Port+")/"+
			conf.Conf.MysqlConf.Database+
			"?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
		defer DB.Close()
	}
}

func initReportDb() {
	ReportDb, err = gorm.Open(
		"mysql",
		conf.Report.MysqlReport.User+
			":"+conf.Report.MysqlReport.Password+
			"@tcp("+conf.Report.MysqlReport.Ip+
			conf.Report.MysqlReport.Port+")/"+
			conf.Report.MysqlReport.Database+
			"?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
		defer ReportDb.Close()
	}
}
