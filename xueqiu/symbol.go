package xueqiu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/db"
	"net/url"
)

var Symbol Symbol_

type Symbol_ struct {
	Symbol string
	Pct float64 //当日涨幅
	Volume float64 //本日交易量
	Current float64 //当前价格
	Mc float64 //市值
	Name string //names
	Exchange string //市场范围
	Type int
	Areacode string //地区
	Tick_size float64
	Indcode string
}

func (s Symbol_) getTableName() string {
	return "symbol"
}

func (s Symbol_) createTable()  {
	tableKey := s.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&s)
	}
}

func (s Symbol_)Add(){
	db.DB.Table(s.getTableName()).Create(&s)
}

type SymbolRequest_ Request_

var SymbolRequest  SymbolRequest_

func (request SymbolRequest_)initRequest() SymbolRequest_{

	request.SearchUrl  = "https://xueqiu.com/service/screener/screen"


	request.SearchParms = url.Values{}
	request.SearchParms.Add("category","CN")
	request.SearchParms.Add("exchange","sh_sz")
	request.SearchParms.Add("areacode","")
	request.SearchParms.Add("indcode","")
	request.SearchParms.Add("order_by","mc")
	request.SearchParms.Add("order","desc")
	request.SearchParms.Add("page","1")
	request.SearchParms.Add("size","3800")
	request.SearchParms.Add("only_count","0")
	return request

}


func (s SymbolRequest_) Run ()  {
	Symbol.createTable()

	request := s.initRequest()
	data := Get(request.SearchUrl,request.SearchParms)


	str:=[]byte(data)


	rs := SymbolJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	for _,symbol := range rs.Data.List{
		symbol.Add()
	}
}



//chart json
type  SymbolJsonResponse struct{
	ErrorCode int `json:"error_code"`
	Data SymbolJsonData  `json:"data"`
}

type  SymbolJsonData struct{
	Count int
	List []Symbol_
}


