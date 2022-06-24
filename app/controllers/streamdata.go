package controllers

import (
	"log"

	"github.com/giver-yell/goTrading/app/models"
	"github.com/giver-yell/goTrading/bitflyer"
	"github.com/giver-yell/goTrading/config"
)

func StreamIngestionData() {
	var tickerChannel = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		log.Printf("action=StreamIngestionData, %v", ticker)
		for _, duration := range config.Config.Durations {
			isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
			if isCreated == true && duration == config.Config.TradeDuration {
				// TODO
			}
		}
	}
}
