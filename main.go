package main

import (
	"fmt"
	"time"

	"github.com/giver-yell/goTrading/bitflyer"
	"github.com/giver-yell/goTrading/config"
	"github.com/giver-yell/goTrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	// fmt.Println(apiClient.GetBalance())

	// ticker, _ := apiClient.GetTicker("BTC_USD")
	// fmt.Println(ticker)
	// fmt.Println(ticker.GetMidPrice())
	// fmt.Println(ticker.DateTime())
	// fmt.Println(ticker.TruncateDateTime(time.Hour))

	tickerChannel := make(chan bitflyer.Ticker)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		fmt.Println(ticker)
		fmt.Println(ticker.GetMidPrice())
		fmt.Println(ticker.DateTime())
		fmt.Println(ticker.TruncateDateTime(time.Second))
		fmt.Println(ticker.TruncateDateTime(time.Minute))
		fmt.Println(ticker.TruncateDateTime(time.Hour))
	}
}
