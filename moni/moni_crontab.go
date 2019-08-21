package moni

import (
	"fmt"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/xueqiu"
)

//增加每天的 group 1 每日创新低股票

func AddmoniGroup1(gatherDay string)  {
	sql := "select a.symbol,a.close,a.pe,a.pb,a.ps,a.gather_day from chang.stock_chart a left join "
	sql += "(select symbol,  min(pe) as pe from chang.stock_chart where pe > 1 group by symbol) b on a.symbol = b.symbol "
	sql += "where gather_day = ? and a.pe=b.pe;"

	rows1,_ := db.DB.Raw(sql, gatherDay).Rows() // (*sql.Rows, error)
	defer rows1.Close()
	for rows1.Next() {
		var st xueqiu.StockChart
		db.DB.ScanRows(rows1, &st)
		fmt.Println(st)
	}


}
