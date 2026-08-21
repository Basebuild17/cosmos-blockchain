#!/bin/bash

# Initialize a blockchain node in Docker
# Usage: ./scripts/init-docker.sh [chain-id] [moniker]

set -e

CHAIN_ID=${1:-mychain-1}
MONIKER=${2:-validator}
HOME=/home/myblockchain/.myblockchain

echo "Initializing blockchain node..."
echo "Chain ID: $CHAIN_ID"
echo "Moniker: $MONIKER"

# Initialize the node
myblockaind init $MONIKER --chain-id=$CHAIN_ID --home=$HOME

echo "✓ Node initialized"
echo "Config location: $HOME/config/"
echo ""
echo "Next steps:"
echo "1. Copy genesis file to: $HOME/config/genesis.json"
echo "2. Update config.toml with seeds and persistent peers"
echo "3. Start the node with: docker-compose up"
