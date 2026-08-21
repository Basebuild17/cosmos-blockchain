# Docker Setup for Cosmos Blockchain

This directory contains Docker configurations for easy blockchain deployment.

## Files

- **Dockerfile** - Main blockchain validator node image
- **Dockerfile.cli** - CLI client image
- **docker-compose.yml** - Single validator setup
- **docker-compose.testnet.yml** - 3-node testnet setup
- **docker-compose.monitoring.yml** - Validator + Prometheus + Grafana

## Quick Start

### Option 1: Single Validator (Easiest)

```bash
# Build and run
docker-compose up -d

# Check status
curl http://localhost:26657/status | jq

# Use CLI
docker-compose exec cli myblockchainicli query bank balances

# View logs
docker-compose logs -f validator

# Stop
docker-compose down
```

### Option 2: 3-Node Testnet

```bash
# Use testnet compose file
docker-compose -f docker-compose.testnet.yml up -d

# Check validators
curl http://localhost:26657/status | jq  # validator-1
curl http://localhost:26667/status | jq  # validator-2
curl http://localhost:26677/status | jq  # validator-3

# Stop
docker-compose -f docker-compose.testnet.yml down
```

### Option 3: With Monitoring (Prometheus + Grafana)

```bash
# Build and run with monitoring
docker-compose -f docker-compose.monitoring.yml up -d

# Access Grafana
# http://localhost:3000 (admin/admin)

# Access Prometheus
# http://localhost:9000

# Check validator
curl http://localhost:26657/status | jq
```

## Detailed Usage

See **docs/DOCKER.md** for comprehensive guide including:
- Building images
- Configuration
- Volume management
- Network setup
- Troubleshooting

## Common Commands

```bash
# Build images
docker-compose build

# Start services
docker-compose up -d

# View logs
docker-compose logs -f validator

# Execute commands
docker-compose exec validator myblockaind status

# Initialize node
docker-compose exec validator myblockaind init mynode --chain-id=mychain-1

# Stop services
docker-compose down

# Clean up volumes
docker-compose down -v
```

## Port Mappings

| Service | Port | Purpose |
|---------|------|----------|
| P2P | 26656 | Peer-to-peer communication |
| RPC | 26657 | JSON-RPC endpoint |
| REST API | 1317 | REST API endpoint |
| gRPC | 9090 | gRPC endpoint |
| Grafana | 3000 | Dashboard (monitoring) |
| Prometheus | 9000 | Metrics (monitoring) |

## Scripts

- **scripts/init-docker.sh** - Initialize single node
- **scripts/init-testnet-docker.sh** - Initialize testnet
- **scripts/setup-testnet-docker.sh** - Complete testnet setup

```bash
# Make executable
chmod +x scripts/*.sh

# Use
./scripts/init-docker.sh mychain-1 myvalidator
```
