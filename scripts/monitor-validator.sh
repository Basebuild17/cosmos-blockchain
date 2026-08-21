#!/bin/bash

# Monitor validator node status
# Continuously displays validator status and key metrics

set -e

echo "================================"
echo "Validator Monitor"
echo "================================"
echo ""

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# Check if validator is running
if ! docker-compose ps validator | grep -q 'Up'; then
    echo -e "${RED}✗ Validator is not running${NC}"
    echo "Start with: docker-compose up -d"
    exit 1
fi

echo -e "${GREEN}✓ Validator is running${NC}"
echo ""

# Function to get validator status
get_status() {
    curl -s http://localhost:26657/status 2>/dev/null || echo "{}"
}

# Function to get validators
get_validators() {
    docker-compose exec -T cli myblockchainicli query staking validators --output=json 2>/dev/null || echo "{}"
}

# Function to format time
format_time() {
    date '+%H:%M:%S'
}

# Monitor loop
while true; do
    clear
    echo -e "${BLUE}Validator Monitor - $(format_time)${NC}"
    echo "================================"
    echo ""
    
    # Get status
    STATUS=$(get_status)
    
    # Extract key info
    BLOCK_HEIGHT=$(echo $STATUS | jq -r '.result.sync_info.latest_block_height // "N/A"')
    CATCHING_UP=$(echo $STATUS | jq -r '.result.sync_info.catching_up // "N/A"')
    BLOCK_TIME=$(echo $STATUS | jq -r '.result.sync_info.latest_block_time // "N/A"')
    MONIKER=$(echo $STATUS | jq -r '.result.node_info.moniker // "N/A"')
    NETWORK=$(echo $STATUS | jq -r '.result.node_info.network // "N/A"')
    PEERS=$(echo $STATUS | jq -r '.result.net_info.n_peers // "0"')
    
    # Display info
    echo -e "${GREEN}Node Information:${NC}"
    echo "  Moniker:        $MONIKER"
    echo "  Chain ID:       $NETWORK"
    echo "  Version:        $(docker-compose exec -T validator myblockaind version 2>/dev/null | head -1 || echo 'N/A')"
    echo ""
    
    echo -e "${GREEN}Sync Status:${NC}"
    if [ "$CATCHING_UP" = "false" ]; then
        echo -e "  Status:         ${GREEN}✓ Synced${NC}"
    else
        echo -e "  Status:         ${YELLOW}⟳ Syncing${NC}"
    fi
    echo "  Block Height:   $BLOCK_HEIGHT"
    echo "  Block Time:     $BLOCK_TIME"
    echo "  Connected Peers: $PEERS"
    echo ""
    
    # Get validator status
    VALIDATORS=$(get_validators)
    VAL_COUNT=$(echo $VALIDATORS | jq '.validators | length' 2>/dev/null || echo "0")
    
    echo -e "${GREEN}Validator Status:${NC}"
    echo "  Total Validators: $VAL_COUNT"
    echo ""
    
    # Display ports
    echo -e "${GREEN}Ports:${NC}"
    echo "  P2P:     http://localhost:26656"
    echo "  RPC:     http://localhost:26657"
    echo "  REST:    http://localhost:1317"
    echo "  gRPC:    http://localhost:9090"
    echo ""
    
    echo -e "${BLUE}Updated every 5 seconds (press Ctrl+C to exit)${NC}"
    
    sleep 5
done
