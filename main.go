package main

import (
	"fmt"
	"time"

	"github.com/giver-yell/goTrading/bitflyer"
	"github.com/giver-yell/goTrading/config"
	"github.com/giver-yell/goTrading/utils"
)

func main() {
	// 95.アプリのロギングの設定
	utils.LoggingSettings(config.Config.LogFile)
	// log.Println("test")
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	// fmt.Println(apiClient.GetBalance())

	// 98.Ticker APIでBitcoinデータを取得
	ticker, _ := apiClient.GetTicker("BTC_USD")
	fmt.Println(ticker)
	fmt.Println(ticker.GetMidPrice())
	fmt.Println(ticker.DateTime())
	fmt.Println(ticker.TruncateDateTime(time.Hour))

	// tickerChannel := make(chan bitflyer.Ticker)
	// go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	// for ticker := range tickerChannel {
	// 	fmt.Println(ticker)
	// 	fmt.Println(ticker.GetMidPrice())
	// 	fmt.Println(ticker.DateTime())
	// 	fmt.Println(ticker.TruncateDateTime(time.Second))
	// 	fmt.Println(ticker.TruncateDateTime(time.Minute))
	// 	fmt.Println(ticker.TruncateDateTime(time.Hour))
	// }

	// 101.Bitcoinを売買するAPI
	// order := &bitflyer.Order{
	// 	ProductCode: config.Config.ProductCode,
	// 	// 成り行き
	// 	ChildOrderType:  "MARKET",
	// 	Side:            "BUY",
	// 	Size:            0.01,
	// 	MinuteToExpires: 1,
	// 	TimeInForce:     "GTC",
	// }
	// res, _ := apiClient.SendOrder(order)
	// fmt.Println(res.ChildOrderAcceptanceID)

	// fmt.Println(models.DbConnection)

	// 94.アプリのconfigの設定
	// fmt.Println(config.Config.ApiKey)
	// fmt.Println(config.Config.ApiSecret)
}
