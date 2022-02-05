package main

import (
	"log"

	"github.com/giver-yell/goTrading/config"
	"github.com/giver-yell/goTrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	// fmt.Println(config.Config.ApiKey)
	// fmt.Println(config.Config.ApiSecret)
}
