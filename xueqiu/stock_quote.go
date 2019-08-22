package xueqiu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/jinzhu/gorm"
	"net/url"
	"time"
)

// https://stock.xueqiu.com/v5/stock/quote.json?symbol=SH601155&extend=detail
//只能查看当天的，所以脚本不跑的话，就没有数据，数据不一定全，不过有股息，以及今天比较全面


var StockQuote  StockQuote_
var StockQuoteRequest  StockQuoteRequest_
type StockQuote_ struct {
	gorm.Model
	Symbol string
	Code string
	Currency string
	Name string
	Exchange string
	GatherDay string  `gorm:"column:gather_day" sql:"type:date"`
	High52w float64
	Avg_price float64
	Delayed int
	Percent float64
	Tick_size float64
	Float_shares float64
	Limit_down float64
	Amplitude float64
	Current float64
	High float64
	Current_year_percent float64 //年增幅
	Float_market_capital float64 //流通市值
	Issue_date float64
	Low float64
	Sub_type string
	Market_capital float64 //市值
	Dividend float64 //股息
	Dividend_yield float64 //股息率
	Lot_size int
	Navps float64 //每股净资产
	Profit float64
	Timestamp float64
	Pe_lyr float64 //静态市盈率
	Pe_ttm float64 //市盈率
	Pe_forecast float64 //市盈率(动)
	Amount float64 //成交额
	Chg float64
	Eps float64
	Last_close float64
	Profit_four float64
	Volume float64
	Volume_ratio float64
	Pb float64
	Limit_up float64
	Turnover_rate float64
	Low52w float64
	Time float64
	Total_shares float64
	Open float64
}

func (s StockQuote_) getTableName() string {
	return "stock_quote"
}

func (s StockQuote_) CreateTable()  {
	tableKey := s.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&s)
		db.DB.Table(tableKey).AddUniqueIndex(s.getTableName() +"symbol_gatherday", "symbol","gather_day")
	}
}

func (s StockQuote_)Add() StockQuote_{
	is := s.FindOne()
	if is.ID == 0 {
		db.DB.Table(s.getTableName()).Create(&s)
	}else {
		return is
	}
	return s
	//Exce 会释放连接池
}

func (s StockQuote_) FindOne() StockQuote_{
	dbn := db.DB.Table(s.getTableName()).Where("symbol = ?",s.Symbol).Where("gather_day = ?",s.GatherDay)
	dbn.First(&s)
	return s
}




type StockQuoteRequest_ Request_

func (request StockQuoteRequest_) initRequest(symbol string) StockQuoteRequest_{
	request.SearchUrl  = "https://stock.xueqiu.com/v5/stock/quote.json"
	request.SearchParms = url.Values{}
	request.SearchParms.Add("symbol",symbol)
	request.SearchParms.Add("extend","detail")
	return request
}


func (request StockQuoteRequest_) Run (symbol string)  {
	request = request.initRequest(symbol)
	data := Get(request.SearchUrl,request.SearchParms)


	str:=[]byte(data)


	rs := XueqiuStockQuoteJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}

	stock := rs.Data.Quote
	va := stock.Timestamp /1000
	int64_ := int64(va)
	tm := time.Unix(int64_ , 0)
	stock.GatherDay = tm.Format("2006-01-02")
	stock.Add()

}



//chart json
type  XueqiuStockQuoteJsonResponse struct{
	ErrorCode int `json:"error_code"`
	Data XueqiuStockQuoteJsonData  `json:"data"`
}

type  XueqiuStockQuoteJsonData struct{
	Quote StockQuote_
}






