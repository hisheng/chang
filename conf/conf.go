package conf

import "github.com/BurntSushi/toml"

var Conf Config

type Config struct {
	Debug bool
	Crontab bool
	Redis *Redis	`toml:"redis"`
	Admin *Ad `toml:"ad"`
	HttpServer `toml:"httpserver"`
	MysqlConf `toml:"mysql"`
}
func (c *Config)Init() {
	_, err := toml.DecodeFile("conf.toml", c)
	if err != nil {
		panic(err)
	}
}

type HttpServer struct {
	Addr string `toml:"addr"`
}

type MysqlConf struct {
	Ip string
	Port string
	User string
	Password string
	Database string
}

type Redis struct {
	IP string	`toml:"ip"`
	Password string	`toml:"password"`
}

type Ad struct {
	Clientid string	`toml:"clientid"`
	Secretkey string	`toml:"secretkey"`
	Trackurl string `toml:"trackurl"`
}



