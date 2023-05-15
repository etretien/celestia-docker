package main

import (
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("CELESITA_RPC_URL", "http://127.0.0.1:26657")
	viper.SetDefault("PORT", "10000")
	exporterHTTP()
}
