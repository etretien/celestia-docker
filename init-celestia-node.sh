#!/bin/bash

# set debug, exit on error and no unset variables
set -xeu

mkdir -p volumes/celestia-node-light
sudo chown 10001:10001 volumes/celestia-node-light -R
docker compose -f compose-node-light.yaml up -d
