package main

import (
	"fmt"
	"tradingBitCoin/bitflyer"
	"tradingBitCoin/config"
	"tradingBitCoin/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
