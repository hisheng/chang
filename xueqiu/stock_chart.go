package xueqiu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/jinzhu/gorm"
	"net/url"
	"strconv"
	"time"
)

/*
xueqiu stock chart json
*/

var StockChartModel StockChart
var StockChartRequest StockChartRequest_


type StockChart struct {
	gorm.Model
	Symbol string `sql:"type:varchar(20)"`
	Timestamp float64
	Volume float64
	Open float64   `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	High float64  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Low float64   `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Close float64  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Chg float64  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Percent float64  `sql:"type:decimal(10,2) DEFAULT '0.00' "`
	Turnoverrate float64  `sql:"type:decimal(10,2) DEFAULT '0.00' "`
	Amount float64
	VolumePost float64 `json:"volume_post"`
	AmountPost float64 `json:"amount_post"`
	Pe float64  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Pb float64  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Ps float64  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Pcf float64  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	MarketCapital float64  `sql:"type:decimal(20,2) DEFAULT '0.00' "`
	Balance float64   `sql:"type:decimal(20,2) DEFAULT '0.00' "`
	GatherDay string  `gorm:"column:gather_day" sql:"type:date"`
}

func (s StockChart) getTableName() string {
	return "stock_chart"
}

func (s StockChart) CreateTable()  {
	tableKey := s.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&s)
		db.DB.Table(tableKey).AddUniqueIndex(s.getTableName() +"symbol_gatherday", "symbol","gather_day")
	}
}



func (s StockChart)Add() StockChart{
	is := s.FindOne()
	if is.ID == 0 {
		db.DB.Table(s.getTableName()).Create(&s)
	}else {
		return is
	}
	return s
	//Exce 会释放连接池
}

func (s StockChart) FindOne() StockChart{
	dbn := db.DB.Table(s.getTableName()).Where("symbol = ?",s.Symbol).Where("gather_day = ?",s.GatherDay)
	dbn.First(&s)
	return s
}







type StockChartRequest_ Request_

func (request StockChartRequest_) initRequest(symbol,day string) StockChartRequest_{

	request.SearchUrl  = "https://stock.xueqiu.com/v5/stock/chart/kline.json"

	request.SearchParms = url.Values{}
	request.SearchParms.Add("symbol",symbol)

	now := time.Now().Unix()
	request.SearchParms.Add("begin",strconv.FormatInt(now,10)+"000")
	request.SearchParms.Add("period","day")
	request.SearchParms.Add("type","before")
	count := "10000"
	if len(day) > 0{
		count = day
	}
	request.SearchParms.Add("count","-"+count) //往前多少天的数据

	request.SearchParms.Add("indicator","kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance")
	return request

}


func (request StockChartRequest_) Run (symbol string,day string)  {
	request = request.initRequest(symbol,day)
	data := Get(request.SearchUrl,request.SearchParms)


	str:=[]byte(data)


	rs := XueqiuChartJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	for _,stock := range rs.Data.Item{
		stockChartModel := StockChart{}
		stockChartModel.Symbol = rs.Data.Symbol


		columns := rs.Data.Column

		for index,column := range columns {
			val := stock[index]
			switch column {
			case "timestamp":
				stockChartModel.Timestamp = val

				va := stockChartModel.Timestamp /1000
				int64_ := int64(va)
				tm := time.Unix(int64_ , 0)
				stockChartModel.GatherDay = tm.Format("2006-01-02")

			case "volume":
				stockChartModel.Volume = val
			case "open":
				stockChartModel.Open = val
			case "high":
				stockChartModel.High = val
			case "low":
				stockChartModel.Low = val
			case "close":
				stockChartModel.Close = val
			case "chg":
				stockChartModel.Chg = val
			case "percent":
				stockChartModel.Percent = val
			case "turnoverrate":
				stockChartModel.Turnoverrate = val
			case "amount":
				stockChartModel.Amount = val
			case "volume_post":
				stockChartModel.VolumePost = val
			case "amount_post":
				stockChartModel.AmountPost = val
			case "pe":
				stockChartModel.Pe = val
			case "pb":
				stockChartModel.Pb = val
			case "ps":
				stockChartModel.Ps = val
			case "pcf":
				stockChartModel.Pcf = val
			case "market_capital":
				stockChartModel.MarketCapital = val
			case "balance":
				stockChartModel.Balance = val
			}

		}
		stockChartModel.Add()
	}
}



//chart json
type  XueqiuChartJsonResponse struct{
	ErrorCode int `json:"error_code"`
	Data XueqiuChartJsonData  `json:"data"`
}

type  XueqiuChartJsonData struct{
	Symbol string
	Column []string
	Item []XueqiuChartJsonDataItme
}

type XueqiuChartJsonDataItme []float64
