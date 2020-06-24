package xueqiu

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/library/curl"
	"github.com/jinzhu/gorm"
	"net/url"
)

var StockSummaryModel StockSummary

type StockSummary struct {
	gorm.Model
	Symbol  string  `json:"symbol"`
	Pettm   float32 `json:"pettm"`   // pe ttm
	Npay    float32 `json:"npay"`    // 净利润同比增
	Current float32 `json:"current"` //目前价格
	Name    string  `json:"name"`    //名
	Oiy     float32 `json:"oiy"`     // 营业收入同比增长
	OiyPe   float32 //  营业收入同比增长 / Pettm
}

func (stockSummary *StockSummary) InitOiyPe() {
	stockSummary.OiyPe = stockSummary.Oiy / stockSummary.Pettm
}

func (s StockSummary) createTable() {
	tableKey := s.getTableName()
	if !db.DB.HasTable(tableKey) {
		db.DB.Table(tableKey).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&s)
	}
}

func (s StockSummary) getTableName() string {
	tableKey := "stock_summary"
	return tableKey
}

func initSearchRequest() Request_ {
	request := new(Request_)

	request.SearchUrl = "https://xueqiu.com/service/screener/screen"

	request.SearchParms = url.Values{}
	request.SearchParms.Add("category", "CN")
	request.SearchParms.Add("exchange", "sh_sz")
	request.SearchParms.Add("order_by", "pettm")
	request.SearchParms.Add("order", "asc")
	request.SearchParms.Add("page", "1")
	request.SearchParms.Add("size", "60")
	request.SearchParms.Add("current", "")
	request.SearchParms.Add("pettm", "4_10")
	request.SearchParms.Add("oiy.20190331", "10_1000")
	request.SearchParms.Add("npay.20190331", "10_1000")
	request.SearchParms.Add("_", "1564377012355")
	request.SearchParms.Add("pct", "")
	request.SearchParms.Add("only_count", "0")

	return *request
}

func (s StockSummary) Run() {
	request := initSearchRequest()
	data := curl.Get(request.SearchUrl, request.SearchParms)

	fmt.Println(data)

	str := []byte(data)

	rs := XueqiuJsonResponse{}

	err := json.Unmarshal(str, &rs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs)

	fmt.Println(rs.Data.List)

	for _, stock := range rs.Data.List {
		stock.InitOiyPe()
		s := fmt.Sprintf("%s %4s 营业收入增长 %.2f 利润增长 %.2f pe为 %.2f 性价比指数为 %.2f", stock.Symbol, stock.Name, stock.Oiy, stock.Npay, stock.Pettm, stock.Npay/stock.Pettm)
		fmt.Println(s)
		//return fmt.Sprintf("sm,SM,um,UM=%d,%d,%d,%d", l.min, l.max, l.umin, l.umax)

	}
}

type XueqiuJsonResponse struct {
	ErrorCode int `json:"error_code"`
	Data      XueqiuJsonData
}
type XueqiuJsonData struct {
	Count int            `json:"count"`
	List  []StockSummary `json:"list"`
}
