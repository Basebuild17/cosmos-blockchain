# Quick Start: Single Validator

This guide will get your validator running in **5 minutes**.

## Prerequisites

- Docker: https://docs.docker.com/get-docker/
- Docker Compose: https://docs.docker.com/compose/install/

## Step 1: Quick Start (Automated)

```bash
# Navigate to project directory
cd cosmos-blockchain

# Make scripts executable
chmod +x scripts/*.sh

# Start validator (builds, initializes, and starts)
bash scripts/start-validator.sh
```

This script will:
1. ✓ Check Docker and Docker Compose installation
2. ✓ Build Docker images
3. ✓ Initialize validator node
4. ✓ Configure node settings
5. ✓ Start the validator

**Expected output:**
```
================================
✓ Validator Started!
================================

Node Information:
  Chain ID: mychain-1
  Moniker: myvalidator

Ports:
  P2P:     http://localhost:26656
  RPC:     http://localhost:26657
  REST:    http://localhost:1317
  gRPC:    http://localhost:9090
```

## Step 2: Verify It's Running

In a new terminal:

```bash
# Check node status
curl http://localhost:26657/status | jq

# Expected: Shows node info with catching_up status
```

## Step 3: Monitor the Validator

```bash
# Use the monitor script
bash scripts/monitor-validator.sh

# Or manually view logs
docker-compose logs -f validator
```

## Step 4: Create Validator Account

```bash
# Create account key
bash scripts/validator-operations.sh create-key my-validator

# Save the mnemonic phrase somewhere safe!
```

## Step 5: Query Information

```bash
# Show your key
bash scripts/validator-operations.sh show-key my-validator

# List all keys
bash scripts/validator-operations.sh list-keys

# Check node status
bash scripts/validator-operations.sh status

# List validators
bash scripts/validator-operations.sh list-validators

# Check connected peers
bash scripts/validator-operations.sh peers
```

## Manual Commands (If Not Using Scripts)

```bash
# Build images
docker-compose build

# Start validator and CLI
docker-compose up -d

# Initialize blockchain
docker-compose exec validator myblockaind init myvalidator --chain-id=mychain-1

# Copy genesis file
cp genesis/genesis.json ./validator-data/config/

# Restart (applies genesis)
docker-compose restart validator

# View logs
docker-compose logs -f validator

# Wait for sync (check catching_up: false)
curl http://localhost:26657/status | jq .result.sync_info
```

## Common Operations

### Check Account Balance

```bash
bash scripts/validator-operations.sh balance cosmos1...
```

### Generate Gentx (Genesis Transaction)

```bash
bash scripts/validator-operations.sh gentx my-validator 1000000stake mychain-1
```

### Create Validator

```bash
bash scripts/validator-operations.sh create-validator my-validator 1000000stake mychain-1
```

### View Latest Block

```bash
bash scripts/validator-operations.sh blocks
```

## Port Access

| Port | Service | Access |
|------|---------|--------|
| 26656 | P2P | http://localhost:26656 |
| 26657 | RPC | http://localhost:26657 |
| 1317 | REST | http://localhost:1317 |
| 9090 | gRPC | http://localhost:9090 |

## API Endpoints

```bash
# JSON-RPC (Tendermint)
http://localhost:26657/

# REST API (Cosmos SDK)
http://localhost:1317/

# Examples:
curl http://localhost:26657/status | jq
curl http://localhost:1317/cosmos/bank/v1beta1/balances/cosmos1... | jq
```

## Stop the Validator

```bash
# Stop (keeps data)
docker-compose down

# Stop and remove data
docker-compose down -v
```

## Troubleshooting

### "Cannot connect to Docker daemon"
```bash
# Make sure Docker is running
sudo systemctl start docker

# Or start Docker Desktop (macOS/Windows)
```

### "Port 26657 already in use"
```bash
# Stop existing validator
docker-compose down

# Or use different port in docker-compose.yml
ports:
  - "26658:26657"  # Use 26658 instead of 26657
```

### "Validator won't start"
```bash
# Check logs
docker-compose logs validator

# Rebuild
docker-compose build --no-cache
docker-compose up -d
```

### "catching_up: true" for too long
```bash
# Wait longer (normal for large chains)
# Or reset and resync
docker-compose exec validator myblockaind unsafe-reset-all
docker-compose restart validator
```

## Next Steps

1. **For local testing:**
   - Create accounts
   - Send transactions
   - Query blockchain state

2. **For testnet:**
   - Download testnet genesis
   - Connect to testnet seeds
   - See docs/TESTNET_CONNECTION.md

3. **For monitoring:**
   - Use docker-compose.monitoring.yml
   - Access Grafana at http://localhost:3000
   - See docs/DOCKER.md

4. **For production:**
   - Deploy on cloud (AWS/GCP/DigitalOcean)
   - Set up sentry nodes
   - Configure monitoring
   - See docs/DOCKER.md for details

## Full Documentation

- Quick reference: `docker/README.md`
- Complete guide: `docs/DOCKER.md`
- Testnet guide: `docs/TESTNET_CONNECTION.md`
- Validator setup: `docs/VALIDATOR_SETUP.md`

---

**Your validator is now running! 🚀**
