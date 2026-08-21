#!/bin/bash

# Initialize testnet with multiple validators in Docker
# Usage: ./scripts/init-testnet-docker.sh

set -e

echo "Initializing testnet with 3 validators..."

# Create containers to initialize
docker-compose -f docker-compose.testnet.yml run --rm validator-1 true 2>/dev/null || true

echo "Initializing validator-1..."
docker-compose -f docker-compose.testnet.yml run --rm validator-1 \
  myblockaind init validator-1 --chain-id=mychain-testnet

echo "Initializing validator-2..."
docker-compose -f docker-compose.testnet.yml run --rm validator-2 \
  myblockaind init validator-2 --chain-id=mychain-testnet

echo "Initializing validator-3..."
docker-compose -f docker-compose.testnet.yml run --rm validator-3 \
  myblockaind init validator-3 --chain-id=mychain-testnet

echo "✓ All validators initialized"
echo ""
echo "Next steps:"
echo "1. Create validator keys in each container"
echo "2. Generate gentx files"
echo "3. Collect gentx files"
echo "4. Start the testnet with: docker-compose -f docker-compose.testnet.yml up"
echo ""
echo "Or use the guided setup: ./scripts/setup-testnet-docker.sh"
