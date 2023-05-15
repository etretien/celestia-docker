# celestia docker stack
This repo provides easy way to deploy and monitor celestia nodes (celestia-app and light node) using docker compose and prometheus stack. Information about node state gathered over local rpc queries via exporter and displayed in Grafana dashboard.
Additionally it installs node exporter, cadvisor and corresponding Grafana dashboards.
Written and tested on blockspacerace chain.

### Getting started:

```
git clone https://github.com/etretien/celestia-docker
cd celestia-docker

```

Next, depending on what node you want to run, start init script:
* celestia-app
`MONIKER=<node name> ./init-celestia-app.sh`
* celestia-node
`./init-celestia-node.sh`
* or both

celestia-app is built from source (https://github.com/etretien/celestia-docker/blob/main/Dockerfile), while celestia-node is pulled from official celestiaorg registry.

Start monitoring services:
```
./init-monitoring.sh
```

Grafana should be accessible at http://localhost:3000, default login `admin` and password `admin`

### Dashboard

TODO screenshot
