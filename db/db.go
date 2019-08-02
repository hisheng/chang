package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hisheng/chang/conf"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)


// gorm.Model definition
type Model struct {
	Id int
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;"`
	DeletedAt time.Time  `gorm:"column:deleted_at;"`

}

var (
	DB *gorm.DB
	err error
)

func GetDb() {
	DB, err = gorm.Open(
		"mysql",
		conf.Conf.MysqlConf.User +
			":"+conf.Conf.MysqlConf.Password +
			"@tcp("+conf.Conf.MysqlConf.Ip +
			conf.Conf.MysqlConf.Port+")/" +
			conf.Conf.MysqlConf.Database +
			"?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
		defer DB.Close()
	}

	//return db
}



