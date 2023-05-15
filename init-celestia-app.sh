#!/bin/bash

# set debug, exit on error and no unset variables
set -xeu

echo "Moniker: $MONIKER"
echo "Chain: blockspacerace"

mkdir -p volumes/celestia-app
# docker compose -f compose-app.yaml build

echo
echo getting seeds, peers and genesis from celestiaorg/networks repo
echo

rm -rf networks || true
git clone https://github.com/celestiaorg/networks.git

PEERS=`cat networks/blockspacerace/peers.txt | tr '\n' ',' | sed 's/,$//'`
SEEDS=`cat networks/blockspacerace/seeds.txt | tr '\n' ',' | sed 's/,$//'`

docker compose -f compose-app.yaml run --rm celestia-app /root/celestia-appd init $MONIKER --chain-id blockspacerace-0 2>&1 | tee celestia-appd-init.log

# copy genesis to container volume
cat networks/blockspacerace/genesis.json | docker compose -f compose-app.yaml run --rm -T celestia-app /bin/sh -c "cat > /root/.celestia-app/config/genesis.json"

# using dasel to edit config.toml
# https://github.com/TomWright/dasel
docker compose -f compose-tools.yaml run --rm dasel put -t string -r toml -f .celestia-app/config/config.toml -v "$SEEDS" .p2p.seeds
docker compose -f compose-tools.yaml run --rm dasel put -t string -r toml -f .celestia-app/config/config.toml -v "$PEERS" .p2p.peers
docker compose -f compose-tools.yaml run --rm dasel put -t string -r toml -f .celestia-app/config/config.toml -v "tcp://0.0.0.0:26657" .rpc.laddr

docker compose -f compose-app.yaml up -d
