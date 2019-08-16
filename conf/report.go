package conf

import "github.com/BurntSushi/toml"

var Report Report_

type Report_ struct {
	MysqlReport `toml:"mysql"`

}


func (r *Report_)Init() {
	_, err := toml.DecodeFile("report.toml", r)
	if err != nil {
		panic(err)
	}
}


type MysqlReport struct {
	Ip string
	Port string
	User string
	Password string
	Database string
}