package main

import (
	"fmt"

	"github.com/giver-yell/goTrading/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
