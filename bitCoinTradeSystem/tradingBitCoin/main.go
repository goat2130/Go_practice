package main

import (
	"tradingBitCoin/app/controllers"
	"tradingBitCoin/config"
	"tradingBitCoin/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	//apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	//ticker, _ := apiClient.GetTicker("BTC_USD")
	//fmt.Println(ticker)
	//fmt.Println(ticker.GetMidPrice())
	//fmt.Println(ticker.DateTime())
	//fmt.Println(ticker.TruncateDateTime(time.Hour))

	//order := &bitflyer.Order{
	//	ProductCode:     config.Config.ProductCode,
	//	ChildOrderType:  "LIMIT",
	//	Side:            "BUY",
	//	Price:           7000,
	//	Size:            0.01,
	//	MinuteToExpires: 1,
	//	TimeInForce:     "GTC",
	//}
	//res, _ := apiClient.SendOrder(order)
	//fmt.Println(res.ChildOrderAcceptanceID)

	//i := "JRF20181012-144016-140584"
	//params := map[string]string{
	//	"product_code": config.Config.ProductCode,
	//	"child_order_acceptance_id": i,
	//}
	//r, _ := apiClient.ListOrder(params)
	//fmt.Println(r)

	// fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
