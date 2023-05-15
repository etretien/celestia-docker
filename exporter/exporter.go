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

		u, err := url.Parse(viper.GetString("CELESTIA_RPC_URL"))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		u.Path = "/status"

		err = SetCelestiaMetrics(m, u.String())
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
