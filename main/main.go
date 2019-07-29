package main

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/curl"
	"github.com/hisheng/chang/xueqiu"
)



func main()  {
	xueqiuUrl := xueqiu.Request.SearchUrl
	parms := xueqiu.Request.SearchParms
	fmt.Println(xueqiuUrl + "?" + parms.Encode())

	data := curl.Get(xueqiuUrl,parms)
	fmt.Println(data)

	str:=[]byte(data)


	rs := xueqiu.XueqiuJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs)

}