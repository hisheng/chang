package service

import (
	"encoding/json"
	"fmt"
	"github.com/hisheng/chang/model"
	"github.com/hisheng/chang/repository"
	"github.com/hisheng/chang/xueqiu"
	"net/url"
)

type xueqiuResponse struct {
	ErrorCode        int
	ErrorDescription string
	Data             xueqiuResponseData `json:"data"`
}

type xueqiuResponseData struct {
	addtions []int
	allots   []int
	Items    []model.Bonus `json:"items"`
}

type BonusService struct {
	url             string
	bonusRepository *repository.BonusRepository
}

//https://stock.xueqiu.com/v5/stock/f10/cn/bonus.json?symbol=SH601155&size=10&page=1&extend=true

func NewBonusService() BonusService {
	return BonusService{
		url: "https://stock.xueqiu.com/v5/stock/f10/cn/bonus.json",
	}
}

func (s BonusService) CurlGet(symbol string) {
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("extend", "true")
	data := xueqiu.Get(s.url, params)
	fmt.Println(data)
	jsonData := []byte(data)

	rs := xueqiuResponse{}
	_ = json.Unmarshal(jsonData, &rs)
	fmt.Println(rs)
	fmt.Println(rs.Data.Items)
	s.bonusRepository = repository.NewBonusRepository()
	for _, bonus := range rs.Data.Items {
		fmt.Println(bonus)
		bonus.Symbol = symbol
		bonus.AshareExDividendDate = bonus.AshareExDividendDate / 1000
		bonus.ExDividendDate = bonus.ExDividendDate / 1000
		bonus.EquityDate = bonus.EquityDate / 1000
		bonus.DividendDate = bonus.DividendDate / 1000
		bonus.CancleDividendDate = bonus.CancleDividendDate / 1000
		//db.DB.Create(&bonus)
		s.bonusRepository.Update(&bonus)
	}
}
