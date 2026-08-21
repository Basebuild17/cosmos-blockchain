# Docker Deployment Guide

Complete guide to deploying Cosmos blockchain with Docker.

## Prerequisites

- **Docker** 20.10+
  ```bash
  curl -fsSL https://get.docker.com -o get-docker.sh
  sudo sh get-docker.sh
  ```

- **Docker Compose** 1.29+
  ```bash
  sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
  sudo chmod +x /usr/local/bin/docker-compose
  ```

- **Git** for cloning
  ```bash
  sudo apt install git
  ```

## Architecture Overview

### Single Node Setup
```
┌─────────────────┐
│   Validator     │
│  (Docker)       │
│                 │
│ P2P: 26656      │
│ RPC: 26657      │
│ REST: 1317      │
│ gRPC: 9090      │
└─────────────────┘
```

### Testnet Setup (3 Nodes)
```
┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│ Validator-1 │  │ Validator-2 │  │ Validator-3 │
│  (Docker)   │  │  (Docker)   │  │  (Docker)   │
│             │  │             │  │             │
│ P2P: 26656  │  │ P2P: 26666  │  │ P2P: 26676  │
│ RPC: 26657  │  │ RPC: 26667  │  │ RPC: 26677  │
└──────┬──────┘  └──────┬──────┘  └──────┬──────┘
       │                │               │
       └────────────────┴───────────────┘
              Docker Network
```

### Production Setup with Monitoring
```
┌────────────────┐
│   Validator    │
│    (Docker)    │
│ Metrics↓       │
└────────┬───────┘
         │
┌────────▼───────────────────┐
│   Prometheus (Scraper)     │
│   Metrics Collection       │
└────────┬───────────────────┘
         │
┌────────▼───────────────────┐
│   Grafana (Dashboards)     │
│   Visualization & Alerts   │
│   http://localhost:3000    │
└────────────────────────────┘
```

## Part 1: Build Docker Images

### Build Locally

```bash
# Clone repository
git clone https://github.com/Basebuild17/cosmos-blockchain.git
cd cosmos-blockchain

# Build validator image
docker build -f Dockerfile -t myblockchain:latest .

# Build CLI image
docker build -f Dockerfile.cli -t myblockchain-cli:latest .

# Verify images
docker images | grep myblockchain
```

### Using Pre-built Images (If Published)

```bash
# Pull from registry
docker pull your-registry/myblockchain:latest
docker pull your-registry/myblockchain-cli:latest
```

## Part 2: Single Node Validator

### Step 1: Initialize Node

```bash
# Use script
bash scripts/init-docker.sh mychain-1 myvalidator

# Or manually:
docker-compose up -d validator
docker-compose exec validator myblockaind init myvalidator --chain-id=mychain-1
```

### Step 2: Configure Node

```bash
# Copy genesis file
cp genesis/genesis.json ./validator-data/config/

# Or download testnet genesis
docker-compose exec validator bash -c 'wget https://raw.githubusercontent.com/cosmos/testnets/master/theta-testnet-001/genesis.json -O ~/.myblockchain/config/genesis.json'
```

### Step 3: Start Validator

```bash
# Start in background
docker-compose up -d

# View logs
docker-compose logs -f validator

# Wait for sync (watch for catching_up: false)
watch -n 2 'curl -s http://localhost:26657/status | jq .result.sync_info'
```

### Step 4: Verify

```bash
# Check status
curl http://localhost:26657/status | jq

# Check blocks
curl http://localhost:26657/block | jq

# Check validators
docker-compose exec cli myblockchainicli query staking validators
```

### Step 5: Create Validator (on Testnet)

```bash
# Create key
docker-compose exec cli myblockchainicli keys add my-validator

# Get address
docker-compose exec cli myblockchainicli keys show my-validator -a

# Get tokens from faucet (if testnet)
# Visit faucet URL and paste address

# Create validator
docker-compose exec cli myblockchainicli tx staking create-validator \
  --amount=1000000stake \
  --pubkey=$(docker-compose exec -T validator myblockaind tendermint show-validator) \
  --moniker="MyValidator" \
  --chain-id=theta-testnet-001 \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --from=my-validator \
  --fees=5000stake
```

## Part 3: 3-Node Testnet

### Step 1: Initialize All Nodes

```bash
# Use script for automatic setup
bash scripts/init-testnet-docker.sh

# Or manual initialization:
for i in 1 2 3; do
  docker-compose -f docker-compose.testnet.yml exec -T validator-$i \
    myblockaind init validator-$i --chain-id=mychain-testnet
done
```

### Step 2: Create Keys

```bash
for i in 1 2 3; do
  docker-compose -f docker-compose.testnet.yml exec -T validator-$i \
    myblockchainicli keys add validator-$i --keyring-backend test
done
```

### Step 3: Generate Gentx Files

```bash
for i in 1 2 3; do
  docker-compose -f docker-compose.testnet.yml exec -T validator-$i \
    myblockaind gentx validator-$i 1000000stake \
    --chain-id=mychain-testnet --keyring-backend test
done
```

### Step 4: Collect Gentx

This requires copying gentx files between containers. See advanced section.

### Step 5: Start Testnet

```bash
# Start all validators
docker-compose -f docker-compose.testnet.yml up -d

# View logs
docker-compose -f docker-compose.testnet.yml logs -f

# Check each validator
for i in 1 2 3; do
  echo "Validator $i:"
  PORT=$((26656 + (i-1)*10))
  curl -s http://localhost:$((PORT+1))/status | jq .result.sync_info.catching_up
done
```

## Part 4: With Monitoring (Prometheus + Grafana)

### Step 1: Configure Prometheus

Create `monitoring/prometheus.yml`:

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'myblockchain'
    static_configs:
      - targets: ['validator:26660']
```

### Step 2: Start Monitoring Stack

```bash
# Start validator, Prometheus, and Grafana
docker-compose -f docker-compose.monitoring.yml up -d

# Access Grafana
# http://localhost:3000
# Username: admin
# Password: admin
```

### Step 3: Add Datasource in Grafana

1. Go to http://localhost:3000
2. Login (admin/admin)
3. Configuration → Data Sources
4. Add Prometheus
5. URL: http://prometheus:9090
6. Save & Test

### Step 4: Create Dashboard

1. Create → Dashboard
2. Add Prometheus panel
3. Query: `rate(tendermint_consensus_height[5m])`
4. Save

## Volume Management

### Persist Data

```bash
# Data is stored in Docker volumes
# View volumes
docker volume ls | grep myblockchain

# Inspect volume
docker volume inspect myblockchain-validator-data

# Backup volume
docker run --rm -v myblockchain-validator-data:/data \
  -v $(pwd):/backup alpine tar czf /backup/validator-backup.tar.gz /data

# Restore volume
docker run --rm -v myblockchain-validator-data:/data \
  -v $(pwd):/backup alpine tar xzf /backup/validator-backup.tar.gz -C /
```

## Network Configuration

### Custom Network

```yaml
networks:
  myblockchain:
    driver: bridge
    ipam:
      config:
        - subnet: 172.20.0.0/16
```

### Connect to External Network

```yaml
services:
  validator:
    networks:
      - myblockchain
      - external_network

networks:
  external_network:
    external: true
```

## Environment Variables

### Customize in docker-compose.yml

```yaml
environment:
  - CHAIN_ID=mychain-1
  - MONIKER=myvalidator
  - LOG_LEVEL=info
  - TENDERMINT_TIMEOUT_COMMIT=1s
```

### Or via .env file

Create `.env`:

```env
CHAIN_ID=mychain-1
MONIKER=myvalidator
LOG_LEVEL=info
```

Reference in compose:

```yaml
environment:
  - CHAIN_ID=${CHAIN_ID}
  - MONIKER=${MONIKER}
```

## Logging and Debugging

### View Logs

```bash
# All services
docker-compose logs

# Specific service
docker-compose logs validator

# Follow logs (tail -f)
docker-compose logs -f validator

# Last 100 lines
docker-compose logs --tail=100 validator
```

### Debug Container

```bash
# Execute bash
docker-compose exec validator bash

# Check status
curl http://localhost:26657/status

# View config
cat ~/.myblockchain/config/config.toml

# Check logs
tail -f ~/.myblockchain/logs/node.log
```

## Troubleshooting

### Container Won't Start

```bash
# Check logs
docker-compose logs validator

# Check resource usage
docker stats

# Rebuild image
docker-compose build --no-cache
docker-compose up -d
```

### Sync Issues

```bash
# Check peers
curl http://localhost:26657/net_info | jq .result.peers

# Reset chain (loses data)
docker-compose exec validator myblockaind unsafe-reset-all
docker-compose restart validator
```

### High Memory Usage

```bash
# Limit resources in docker-compose.yml
services:
  validator:
    deploy:
      resources:
        limits:
          cpus: '2'
          memory: 4G
        reservations:
          cpus: '1'
          memory: 2G
```

### Network Issues

```bash
# Check Docker network
docker network ls
docker network inspect myblockchain-myblockchain

# Test connectivity
docker-compose exec validator ping validator-2
```

## Production Considerations

### Security

- Run as non-root user ✓ (Done in Dockerfile)
- Use secrets for private keys
- Restrict port access with firewall
- Use reverse proxy (nginx) for RPC
- Enable authentication

### High Availability

```yaml
# Use Kubernetes instead
# Or Docker Swarm with multiple replicas
# Or run multiple separate instances
```

### Backup Strategy

```bash
# Regular backups
docker volume ls | grep myblockchain | while read vol; do
  docker run --rm -v $vol:/data -v /backups:/backup \
    alpine tar czf /backup/$vol.tar.gz /data
done

# Cron job
0 2 * * * /home/user/backup-volumes.sh
```

### Monitoring

- Use docker-compose.monitoring.yml for Prometheus + Grafana
- Set up alerts for:
  - High memory usage
  - Node sync stopped
  - Validator slashed
  - Missing blocks

## Advanced Usage

### Custom Image Registry

```bash
# Tag image
docker tag myblockchain:latest myregistry/myblockchain:v1.0

# Push to registry
docker push myregistry/myblockchain:v1.0

# Use in docker-compose.yml
image: myregistry/myblockchain:v1.0
```

### Multi-stage Build Optimization

Already implemented in Dockerfile:
- Builder stage (go build)
- Final stage (alpine, ~50MB vs ~800MB)

### Health Checks

Already implemented:

```dockerfile
HEALTHCHECK --interval=30s --timeout=10s --start-period=40s --retries=3 \
    CMD curl -f http://localhost:26657/status || exit 1
```

## Next Steps

1. **Start single validator**: `docker-compose up -d`
2. **Test with testnet**: `docker-compose -f docker-compose.testnet.yml up -d`
3. **Add monitoring**: `docker-compose -f docker-compose.monitoring.yml up -d`
4. **Deploy to cloud** (AWS, GCP, DigitalOcean)
5. **Set up CI/CD** for automatic builds

## Quick Commands Reference

```bash
# Build
docker-compose build

# Start
docker-compose up -d

# Status
docker-compose ps

# Logs
docker-compose logs -f validator

# Execute
docker-compose exec validator bash

# Query blockchain
docker-compose exec cli myblockchainicli query bank balances

# Stop
docker-compose down

# Clean up
docker-compose down -v
```

---

For more details, see `docker/README.md`
