package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/VictoriaMetrics/metrics"
)

type celestiaAppStruct struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		NodeInfo struct {
			ProtocolVersion struct {
				P2P   string `json:"p2p"`
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"protocol_version"`
			ID         string `json:"id"`
			ListenAddr string `json:"listen_addr"`
			Network    string `json:"network"`
			Version    string `json:"version"`
			Channels   string `json:"channels"`
			Moniker    string `json:"moniker"`
			Other      struct {
				TxIndex    string `json:"tx_index"`
				RPCAddress string `json:"rpc_address"`
			} `json:"other"`
		} `json:"node_info"`
		SyncInfo struct {
			LatestBlockHash     string    `json:"latest_block_hash"`
			LatestAppHash       string    `json:"latest_app_hash"`
			LatestBlockHeight   string    `json:"latest_block_height"`
			LatestBlockTime     time.Time `json:"latest_block_time"`
			EarliestBlockHash   string    `json:"earliest_block_hash"`
			EarliestAppHash     string    `json:"earliest_app_hash"`
			EarliestBlockHeight string    `json:"earliest_block_height"`
			EarliestBlockTime   time.Time `json:"earliest_block_time"`
			CatchingUp          bool      `json:"catching_up"`
		} `json:"sync_info"`
		ValidatorInfo struct {
			Address string `json:"address"`
			PubKey  struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"pub_key"`
			VotingPower string `json:"voting_power"`
		} `json:"validator_info"`
	} `json:"result"`
}

func getCelestiaAppRPC(url string) (celestiaAppStruct, error) {
	var celestia celestiaAppStruct

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return celestia, fmt.Errorf("celestia: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return celestia, fmt.Errorf("celestia: %v", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return celestia, fmt.Errorf("celestia: %v", err)
	}

	err = json.Unmarshal(data, &celestia)
	if err != nil {
		return celestia, fmt.Errorf("celestia: %v", err)
	}

	return celestia, err
}

// SetCelestiaMetrics returns metrics set
func SetCelestiaAppMetrics(m *metrics.Set, url string) error {
	celestia, err := getCelestiaAppRPC(url)
	if err != nil {
		return err
	}

	m.GetOrCreateGauge(
		fmt.Sprintf(`celestia_app_sync_info_latest_block_height{moniker="%v"}`, celestia.Result.NodeInfo.Moniker), func() float64 {
			f, err := strconv.ParseFloat(celestia.Result.SyncInfo.LatestBlockHeight, 64)
			if err != nil {
				log.Println(err)
			}
			return f
		})

	m.GetOrCreateGauge(
		fmt.Sprintf(`celestia_app_sync_info_latest_block_time{moniker="%v"}`, celestia.Result.NodeInfo.Moniker), func() float64 {
			return float64(celestia.Result.SyncInfo.LatestBlockTime.Unix())
		})

	m.GetOrCreateGauge(
		fmt.Sprintf(`celestia_app_sync_info_catchingup{moniker="%v"}`, celestia.Result.NodeInfo.Moniker), func() float64 {
			if celestia.Result.SyncInfo.CatchingUp {
				return float64(1)
			}
			return float64(1)
		})

	m.GetOrCreateGauge(
		fmt.Sprintf(`celestia_app_validator_info_voting_power{moniker="%v"}`, celestia.Result.NodeInfo.Moniker), func() float64 {
			f, err := strconv.ParseFloat(celestia.Result.ValidatorInfo.VotingPower, 64)
			if err != nil {
				log.Println(err)
			}
			return f
		})

	return err
}
