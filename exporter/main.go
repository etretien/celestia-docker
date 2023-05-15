package main

import (
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("CELESTIA_APP_RPC_URL", "http://127.0.0.1:26657")
	viper.SetDefault("CELESTIA_NODE_RPC_URL", "http://127.0.0.1:26658")
	viper.SetDefault("PORT", "10000")
	exporterHTTP()
}
