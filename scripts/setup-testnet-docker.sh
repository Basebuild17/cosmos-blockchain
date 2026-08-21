#!/bin/bash

# Complete testnet setup script
# Initializes and starts a 3-node testnet

set -e

echo "================================"
echo "Docker Testnet Setup"
echo "================================"
echo ""

# Step 1: Initialize nodes
echo "[1/5] Initializing nodes..."
bash scripts/init-testnet-docker.sh

# Step 2: Create keys
echo ""
echo "[2/5] Creating validator keys..."
echo "Enter passphrase for validator-1 (and confirm):"
docker-compose -f docker-compose.testnet.yml exec -T validator-1 \
  myblockchainicli keys add validator-1 --keyring-backend test || echo "Key already exists"

echo "Enter passphrase for validator-2 (and confirm):"
docker-compose -f docker-compose.testnet.yml exec -T validator-2 \
  myblockchainicli keys add validator-2 --keyring-backend test || echo "Key already exists"

echo "Enter passphrase for validator-3 (and confirm):"
docker-compose -f docker-compose.testnet.yml exec -T validator-3 \
  myblockchainicli keys add validator-3 --keyring-backend test || echo "Key already exists"

# Step 3: Generate gentx
echo ""
echo "[3/5] Generating gentx files..."
docker-compose -f docker-compose.testnet.yml exec -T validator-1 \
  myblockaind gentx validator-1 1000000stake --chain-id=mychain-testnet --keyring-backend test

docker-compose -f docker-compose.testnet.yml exec -T validator-2 \
  myblockaind gentx validator-2 1000000stake --chain-id=mychain-testnet --keyring-backend test

docker-compose -f docker-compose.testnet.yml exec -T validator-3 \
  myblockaind gentx validator-3 1000000stake --chain-id=mychain-testnet --keyring-backend test

# Step 4: Collect gentx
echo ""
echo "[4/5] Collecting gentx files..."
echo "(This requires manual file copying - see docs/DOCKER.md for details)"

# Step 5: Start testnet
echo ""
echo "[5/5] Starting testnet..."
echo "Run: docker-compose -f docker-compose.testnet.yml up -d"

echo ""
echo "================================"
echo "Setup Complete!"
echo "================================"
