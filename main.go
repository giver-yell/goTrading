package main

import (
	"fmt"

	"github.com/giver-yell/goTrading/bitflyer"
	"github.com/giver-yell/goTrading/config"
	"github.com/giver-yell/goTrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
