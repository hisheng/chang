package xueqiu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/phpgo"
	"github.com/jinzhu/gorm"
	"net/url"
	"strconv"
	"time"
)

/*
xueqiu stock chart json
*/

var StockChartModel StockChart


type StockChart struct {
	gorm.Model
	Symbol string `sql:"type:varchar(20)"`
	Timestamp string `sql:"type:int(11)"`
	Volume string  `sql:"type:bigint DEFAULT '0' "`
	Open string   `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	High string  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Low string   `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Close string  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Chg string  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Percent string  `sql:"type:decimal(10,2) DEFAULT '0.00' "`
	Turnoverrate string  `sql:"type:decimal(10,2) DEFAULT '0.00' "`
	Amount string  `sql:"type:bigint DEFAULT '0' "`
	VolumePost string `json:"volume_post"`
	AmountPost string `json:"amount_post"`
	Pe string  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Pb string  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Ps string  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	Pcf string  `sql:"type:decimal(10,4) DEFAULT '0.0000' "`
	MarketCapital string  `sql:"type:decimal(20,2) DEFAULT '0.00' "`
	Balance string   `sql:"type:decimal(20,2) DEFAULT '0.00' "`
	GatherDay string  `gorm:"column:gather_day" sql:"type:date"`
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
			s.Timestamp = val[:10]

			int64t, err := strconv.ParseInt(s.Timestamp, 10, 64)
			if err != nil {}
			tm := time.Unix(int64t, 0)
			s.GatherDay = tm.Format("2006-01-02")
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
			fmt.Println(val)
			if len(val) > 0 {
				s.Balance = val
			}else {
				s.Balance = "0.00"
			}
		case "symbol":
			s.Symbol = val
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


	str:=[]byte(data)


	rs := XueqiuChartJsonResponse{}

	err:= json.Unmarshal(str,&rs)
	if err != nil {
		fmt.Println(err)
	}
	for _,stock := range rs.Data.Item{
		sm := make(map[string]interface{})
		sm["symbol"] = getSymbol(rs.Data.Symbol)
		for index,d := range stock{
			sm[rs.Data.Column[index]] = d
		}
		StockChartModel.Add(sm)
	}
}

func getSymbol(symbol string) string {
	switch symbol[0:2] {
	case "SH":
		return symbol[2:]
	default:
		return ""
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
