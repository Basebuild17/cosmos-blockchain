#!/bin/bash

# Quick start script for single validator
# This script builds and starts a single validator node

set -e

echo "================================"
echo "Starting Single Validator"
echo "================================"
echo ""

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Step 1: Check prerequisites
echo "[1/5] Checking prerequisites..."
if ! command -v docker &> /dev/null; then
    echo -e "${YELLOW}Docker is not installed${NC}"
    echo "Install from: https://docs.docker.com/get-docker/"
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo -e "${YELLOW}Docker Compose is not installed${NC}"
    echo "Install from: https://docs.docker.com/compose/install/"
    exit 1
fi

echo -e "${GREEN}✓ Docker and Docker Compose installed${NC}"

# Step 2: Build Docker images
echo ""
echo "[2/5] Building Docker images..."
echo "This may take a few minutes on first run..."
docker-compose build --no-cache
echo -e "${GREEN}✓ Docker images built${NC}"

# Step 3: Initialize validator node
echo ""
echo "[3/5] Initializing validator node..."

# Remove old data if exists
if docker volume inspect myblockchain-validator-data &> /dev/null; then
    echo "Removing old validator data..."
    docker-compose down -v 2>/dev/null || true
fi

# Create and initialize
docker-compose up -d validator
echo "Waiting for container to start..."
sleep 5

# Initialize blockchain
echo "Initializing blockchain..."
docker-compose exec -T validator myblockaind init myvalidator --chain-id=mychain-1 --overwrite

echo -e "${GREEN}✓ Validator node initialized${NC}"

# Step 4: Configure node
echo ""
echo "[4/5] Configuring node..."

# Copy genesis file
echo "Copying genesis file..."
docker-compose exec -T validator bash -c 'cp /app/genesis/genesis.json ~/.myblockchain/config/genesis.json' 2>/dev/null || \
docker-compose exec -T validator bash -c 'echo "Genesis file needs to be configured. See docs/DOCKER.md"

echo -e "${GREEN}✓ Node configured${NC}"

# Step 5: Start validator
echo ""
echo "[5/5] Starting validator..."
docker-compose restart validator

echo ""
echo -e "${GREEN}================================${NC}"
echo -e "${GREEN}✓ Validator Started!${NC}"
echo -e "${GREEN}================================${NC}"
echo ""
echo "Node Information:"
echo "  Chain ID: mychain-1"
echo "  Moniker: myvalidator"
echo ""
echo "Ports:"
echo "  P2P:     http://localhost:26656"
echo "  RPC:     http://localhost:26657"
echo "  REST:    http://localhost:1317"
echo "  gRPC:    http://localhost:9090"
echo ""
echo "Commands:"
echo ""
echo "  Check status:"
echo "    curl http://localhost:26657/status | jq"
echo ""
echo "  View logs:"
echo "    docker-compose logs -f validator"
echo ""
echo "  Create account:"
echo "    docker-compose exec cli myblockchainicli keys add myaccount"
echo ""
echo "  Query account:"
echo "    docker-compose exec cli myblockchainicli query bank balances cosmos1..."
echo ""
echo "  Stop validator:"
echo "    docker-compose down"
echo ""
echo "  Stop validator and remove data:"
echo "    docker-compose down -v"
echo ""
echo "For more details, see: docs/DOCKER.md"
echo ""
