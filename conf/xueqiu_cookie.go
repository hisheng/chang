package conf

import "github.com/BurntSushi/toml"

var XueqiuCookie XueqiuCookie_

type XueqiuCookie_ struct {
	Cookies string `toml:"cookies"`
}


func (x *XueqiuCookie_)Init() {
	_, err := toml.DecodeFile("cookie.toml", x)
	if err != nil {
		panic(err)
	}
}
