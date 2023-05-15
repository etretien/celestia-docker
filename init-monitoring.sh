#!/bin/sh

# set debug, exit on error and no unset variables
set -xeu
mkdir -p volumes/monitoring-prometheus
mkdir -p volumes/monitoring-grafana
sudo chown 65534 volumes/monitoring-prometheus
sudo chown 472 volumes/monitoring-grafan

stat exporter.env || touch exporter.env
docker compose -f compose-monitoring.yaml up -d
