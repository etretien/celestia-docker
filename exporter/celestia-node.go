package main

import (
	"context"

	"github.com/VictoriaMetrics/metrics"
	"github.com/celestiaorg/celestia-node/api/rpc/client"
)

func SetCelestiaNodeMetrics(m *metrics.Set, url string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := client.NewPublicClient(ctx, url)
	if err != nil {
		return err
	}
	head, err := client.Header.NetworkHead(ctx)
	if err != nil {
		return err
	}
	
	m.GetOrCreateGauge("celestia_node_sync_info_latest_block_height", func() float64 {
			return float64(head.Height())
		})

	m.GetOrCreateGauge("celestia_node_sync_info_latest_block_time", func() float64 {
			return float64(head.Time().Unix())
		})
	
	client.Close()
	return err
}