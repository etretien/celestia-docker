package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/VictoriaMetrics/metrics"
	"github.com/spf13/viper"
)

func exporterHTTP() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		m := metrics.NewSet()

		uApp, err := url.Parse(viper.GetString("CELESTIA_APP_RPC_URL"))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		uApp.Path = "/status"

		uNode, err := url.Parse(viper.GetString("CELESTIA_NODE_RPC_URL"))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		uNode.Path = "/status"

		err = SetCelestiaAppMetrics(m, uApp.String())
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		err = SetCelestiaNodeMetrics(m, uNode.String())
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		m.WritePrometheus(w)
	})

	addr := fmt.Sprintf(":%d", viper.GetInt("PORT"))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("http", err)
	}
}
