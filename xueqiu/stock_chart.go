package xueqiu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/phpgo"
	"github.com/jinzhu/gorm"
	"net/url"
)

/*
xueqiu stock chart json
*/

var StockChartModel StockChart


type StockChart struct {
	gorm.Model
	Timestamp string
	Volume string
	Open string
	High string
	Low string
	Close string
	Chg string
	Percent string
	Turnoverrate string
	Amount string
	VolumePost string `json:"volume_post"`
	AmountPost string `json:"amount_post"`
	Pe string
	Pb string
	Ps string
	Pcf string
	MarketCapital string
	Balance string
}

func (s StockChart) getTableName() string {
	return "stock_chart"
}

func (s StockChart) createTable()  {
	tableKey := s.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&s)
	}
}


func (s StockChart)Add(sm map[string]interface{}){
	for k,v := range sm{
		val := phpgo.Strval(v)

		switch k {
		case "pe":
			s.Pe = val
		case "pb":
			s.Pb = val
		case "timestamp":
			s.Timestamp = val
		case "volume":
			s.Volume = val
		case "open":
			s.Open = val
		case "high":
			s.High = val
		case "low":
			s.Low = val
		case "close":
			s.Close = val
		case "chg":
			s.Chg = val
		case "percent":
			s.Percent = val
		case "turnoverrate":
			s.Turnoverrate = val
		case "amount":
			s.Amount = val
		case "volume_post":
			s.VolumePost = val
		case "ps":
			s.Ps = val
		case "pcf":
			s.Pcf = val
		case "market_capital":
			s.MarketCapital = val
		case "balance":
			s.Balance = val
		}
	}
	db.DB.Table(s.getTableName()).Create(&s) //会引起 Lock wait timeout exceeded; try restarting transaction

}

func initChartRequest() Request_{
	request := new(Request_)

	request.SearchUrl  = "https://stock.xueqiu.com/v5/stock/chart/kline.json"


	request.SearchParms = url.Values{}
	request.SearchParms.Add("symbol","SH601155")
	request.SearchParms.Add("begin","1564744635504")
	request.SearchParms.Add("period","day")
	request.SearchParms.Add("type","before")
	request.SearchParms.Add("count","-142")
	request.SearchParms.Add("indicator","kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance")


	return *request

}


func (s StockChart) Run ()  {
	s.createTable()

	request := initChartRequest()
	data := Get(request.SearchUrl,request.SearchParms)

	fmt.Println(data)

	str:=[]byte(data)


	rs := XueqiuChartJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs)


	fmt.Println(rs.Data.Symbol)
	fmt.Println(rs.Data.Column)
	fmt.Println(rs.Data.Item)

	for _,stock := range rs.Data.Item{
		sm := make(map[string]interface{})
		for index,d := range stock{
			sm[rs.Data.Column[index]] = d
		}
		fmt.Println(sm)
		StockChartModel.Add(sm)
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
	Item [][]interface{}
}
