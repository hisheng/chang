package moni

import (
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/xueqiu"
	"time"
)

//增加每天的 group 1 每日创新低股票
func AddmoniGroup1(gatherDay string)  {
	Moni.CreateTable()

	sql := "select a.symbol,a.close,a.pe,a.pb,a.ps,a.gather_day from chang.stock_chart a left join "
	sql += "(select symbol,  min(pe) as pe from chang.stock_chart where pe > 1 and gather_day <= ? group by symbol) b on a.symbol = b.symbol "
	sql += "where gather_day = ? and a.pe=b.pe;"

	rows1,_ := db.DB.Raw(sql, gatherDay,gatherDay).Rows() // (*sql.Rows, error)
	defer rows1.Close()
	for rows1.Next() {
		var st xueqiu.StockChart
		db.DB.ScanRows(rows1, &st)
		m := Moni_{Symbol:st.Symbol,Start_price:st.Close,GatherDay:st.GatherDay[:10],Group_code:1,Group_name:"每日创新低股票",
			Pe:st.Pe,Pb:st.Pb,Ps:st.Ps,}
		m.Add()
	}
}

func Init()  {
	for i:=626;i<=5000;i++ {
		yt := beforeToday(i)
		if isWorkDay(yt) {
			AddmoniGroup1(yt.String()[:10])
			fmt.Println(i)
			fmt.Println(yt.String()[:10] + "ok")
		}else {
			fmt.Println(i)
			fmt.Println(yt.String()[:10] + " is not work day")
		}

	}
}


func  beforeToday(day int) time.Time {
	nTime := time.Now()
	nTime = nTime.AddDate(0,0,-day)
	return nTime
}

func isWorkDay(t time.Time)  bool {
	switch t.Weekday() {
	case time.Sunday:
		return false
	case time.Saturday:
		return false
	default:
		return true
	}
}
