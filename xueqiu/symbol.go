package xueqiu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/jinzhu/gorm"
	"net/url"
	"strconv"
)

var Symbol Symbol_

type Symbol_ struct {
	gorm.Model
	Symbol string
	Name string //names
	Exchange string //市场范围
	Type int
	Areacode string //地区
	Area_name string //地区
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
		db.DB.Table(tableKey).AddUniqueIndex(s.getTableName() +"symbol", "symbol")
		//db.DB.Table(tableKey).AddUniqueIndex(s.getTableName() +"symbol", "symbol")
	}
}


func (s Symbol_) Gets() (ss []Symbol_){
	db.DB.Table(s.getTableName()).Find(&ss)
	return ss
}


func (s Symbol_)Add() Symbol_{
	is := s.FindOne(s.Symbol)
	if is.ID == 0 {
		db.DB.Table(s.getTableName()).Create(&s)
	}else {
		return is
	}
	return s
}

func (s Symbol_) FindOne(symbol string) Symbol_{
	dbn := db.DB.Table(s.getTableName()).Where("symbol = ?",symbol)
	dbn.First(&s)
	return s
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
	request.SearchParms.Add("size","6000")
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
		i,_:=strconv.Atoi(symbol.Areacode)
		symbol.Area_name = Areas[i]
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


