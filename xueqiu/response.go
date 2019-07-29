package xueqiu

type  XueqiuJsonResponse struct{
	ErrorCode int `json:"error_code"`
	Data XueqiuJsonData
}
type XueqiuJsonData struct {
	Count int `json:"count"`
	List []StockSummary `json:"list"`
}

type StockSummary struct {
	Symbol string `json:"symbol"`
	Pettm float32 `json:"pettm"`// pe ttm
	Npay float64 `json:"npay"`// 净利润同比增
	Current float32 `json:"current"`//目前价格
	Name string `json:"name"`//名
	Oiy float32 `json:"oiy"`// 营业收入同比增长
}


