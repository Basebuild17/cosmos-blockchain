# Joining Existing Cosmos Testnets

## Quick Reference

This guide explains how to join existing Cosmos Hub testnets or other Cosmos-based testnets.

---

## Popular Cosmos Testnets

### Cosmos Hub Official Testnet

**Repository:** https://github.com/cosmos/testnets

**Current Testnet:** Check the repo for active testnet (e.g., `theta-testnet-001`)

**Key Info:**
- **Chain ID:** See testnet directory
- **Genesis File:** Available in testnet repo
- **Seeds:** Listed in README
- **Block Explorer:** https://testnet-explorer.cosmos.network/

### Setup Steps

```bash
# 1. Clone testnet repo
git clone https://github.com/cosmos/testnets.git
cd testnets

# 2. Find active testnet
ls  # Look for latest theta-testnet or similar

# 3. Read testnet-specific instructions
cat theta-testnet-001/README.md
```

---

## Generic Testnet Setup

### Pre-requisites

```bash
# Ensure you have built the binary
cd cosmos-blockchain
make build

# You now have:
# - ./myblockaind (validator node)
# - ./myblockchainicli (CLI client)
```

### Step-by-Step Connection

#### 1. Download Testnet Genesis

```bash
# Create testnet directory
mkdir -p ~/.myblockchain-testnet

# Download genesis (example - replace with actual testnet)
wget https://raw.githubusercontent.com/cosmos/testnets/master/theta-testnet-001/genesis.json \
  -O ~/.myblockchain-testnet/genesis.json

# Verify downloaded file
sha256sum ~/.myblockchain-testnet/genesis.json
# Compare with official hash from testnet README
```

#### 2. Initialize Node

```bash
# Initialize with testnet chain ID
export TMHOME=~/.myblockchain-testnet
./myblockaind init my-validator --chain-id=theta-testnet-001
```

#### 3. Replace Genesis

```bash
# Use testnet genesis instead of default
cp ~/.myblockchain-testnet/genesis.json $TMHOME/config/genesis.json

# Validate
./myblockaind validate-genesis
```

#### 4. Configure Seeds

Edit `$TMHOME/config/config.toml`:

```bash
nano $TMHOME/config/config.toml
```

Find and update the `seeds` line (copy from testnet README):

```toml
# From: https://github.com/cosmos/testnets/blob/master/theta-testnet-001/README.md
seeds = "bc3442a47dd12dcd4a0e84f6b8c5c03abf1f82f2@seed-1.theta-testnet-001.cosmosnetwork.dev:26656,abc123def456@seed-2.theta-testnet-001.cosmosnetwork.dev:26656"

# Optional: Add persistent peers
persistent_peers = "peer1@peer1.example.com:26656,peer2@peer2.example.com:26656"
```

#### 5. (Optional) Enable State Sync

For faster sync, edit `config.toml`:

```toml
[statesync]
enable = true

# Get these from testnet docs or public RPC
rpc_servers = "https://rpc.testnet.cosmos.network:443,https://rpc2.testnet.cosmos.network:443"
trust_height = 12345  # Get from latest block - some buffer
trust_hash = "HASH_OF_TRUST_HEIGHT_BLOCK"
trust_period = "168h"
```

Get trust height and hash:

```bash
# Query a public RPC
curl https://rpc.testnet.cosmos.network/block?height=12345 | jq .result.block_id.hash
```

#### 6. Configure App Settings

Edit `$TMHOME/config/app.toml`:

```bash
nano $TMHOME/config/app.toml
```

Update:

```toml
# Minimum gas price (check testnet docs)
minimum-gas-prices = "0.025stake"

# Enable API
[api]
enable = true
address = "tcp://0.0.0.0:1317"

# Enable gRPC
[grpc]
enable = true
address = "0.0.0.0:9090"
```

#### 7. Start Syncing

```bash
# Start the node
export TMHOME=~/.myblockchain-testnet
./myblockaind start
```

You should see output like:

```
10:30AM INF Tendermint started height=1
10:30AM INF starting node with ABCI Tendermint client
10:30AM INF syncing
```

#### 8. Monitor Sync Progress

In another terminal:

```bash
# Check sync status
curl http://localhost:26657/status | jq '.result.sync_info'

# Expected output while syncing:
{
  "latest_block_height": "12345",
  "latest_block_time": "2024-01-15T10:30:00Z",
  "earliest_block_height": "1",
  "earliest_block_time": "2024-01-01T00:00:00Z",
  "catching_up": true
}

# When synced:
{
  "catching_up": false
}
```

#### 9. Create Validator Account

Once fully synced:

```bash
# Create new key (or import existing)
./myblockchainicli keys add my-validator-key

# Save the mnemonic!
```

#### 10. Get Testnet Tokens

Most testnets have a faucet:

```bash
# Check your balance
ADDRESS=$(./myblockchainicli keys show my-validator-key -a)
echo $ADDRESS

# Visit faucet URL or use:
# https://testnet-faucet.cosmos.network/
# Paste your address and request tokens

# Check balance
./myblockchainicli query bank balances $ADDRESS
```

#### 11. Create Validator

```bash
# Get your validator pubkey
VALKEY=$(./myblockaind tendermint show-validator)
echo $VALKEY

# Create validator (adjust amounts/rates as needed)
./myblockchainicli tx staking create-validator \
  --amount=1000000stake \
  --pubkey=$VALKEY \
  --moniker="MyValidator" \
  --identity="ABC123" \
  --website="https://example.com" \
  --details="My awesome validator" \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --from=my-validator-key \
  --chain-id=theta-testnet-001 \
  --fees=5000stake \
  -y
```

#### 12. Verify Validator Status

```bash
# List all validators
./myblockchainicli query staking validators

# Find your validator
./myblockchainicli query staking validators | grep moniker

# Check specific validator details
VALIDATOR_ADDR=$(./myblockchainicli query staking validators -o json | jq -r '.validators[0].operator_address')
./myblockchainicli query staking validator $VALIDATOR_ADDR
```

---

## Testnet Chain IDs and Info

### Cosmos Hub

| Network | Chain ID | Status | RPC | API |
|---------|----------|--------|-----|-----|
| Testnet | theta-testnet-001 | Active | https://rpc.testnet.cosmos.network | https://rest.testnet.cosmos.network |
| Devnet | devnet | Active | http://localhost:26657 | http://localhost:1317 |

### Query Network Info

```bash
# Get chain info
curl https://rpc.testnet.cosmos.network/status | jq .result.node_info.network

# Get all validators
curl https://rest.testnet.cosmos.network/cosmos/staking/v1beta1/validators

# Get blocks
curl https://rpc.testnet.cosmos.network/blockchain

# Get specific validator
curl https://rest.testnet.cosmos.network/cosmos/staking/v1beta1/validators/<address>
```

---

## Common Issues

### Can't Connect to Seeds

```bash
# Error: "failed to dial seed"

# Solution 1: Update seeds from testnet README
nano config.toml  # Update seeds line

# Solution 2: Restart
./myblockaind unsafe-reset-all
./myblockaind start
```

### Wrong Genesis Hash

```bash
# Error: "app hash mismatch"

# Solution: Download correct genesis
rm $TMHOME/config/genesis.json
wget https://raw.githubusercontent.com/cosmos/testnets/master/theta-testnet-001/genesis.json \
  -O $TMHOME/config/genesis.json

# Verify hash
sha256sum $TMHOME/config/genesis.json

# Reset and resync
./myblockaind unsafe-reset-all
./myblochaind start
```

### Validator Not in Set

```bash
# Ensure you have enough stake
./myblockchainicli query bank balances $ADDRESS

# Ensure validator created successfully
./myblockchainicli query staking validators | grep $(./myblockchainicli keys show my-validator-key -a)
```

### Slow Sync

```bash
# Enable state sync in config.toml (see Step 5 above)

# Or use prune to reduce disk space:
./myblochaind prune default
```

---

## Advanced Topics

### Running a Sentry Node Architecture

```bash
# Validator node (private, behind sentry)
# Only connects to sentry nodes

# Sentry node 1 (public)
# Connects to validator and other public nodes

# Sentry node 2 (public backup)
# Connects to validator and other public nodes
```

Benefits:
- DDoS protection
- Private validator node
- High availability

Setup:
1. Run separate validator and sentry nodes
2. Configure sentry persistent_peers to validator
3. Configure validator private_peer_ids to exclude sentries
4. Expose only sentry RPC/P2P ports

### Key Management

```bash
# Store keys safely
# Option 1: Hardware wallet integration
# Option 2: Encrypted keyring
# Option 3: Key management service (e.g., Horcrux)

# List all keys
./myblockchainicli keys list

# Export key
./myblockchainicli keys export my-validator-key > key-backup.json

# Import key
./myblockchainicli keys import my-validator-key key-backup.json
```

### Monitoring

```bash
# Monitor validator rewards
watch -n 10 './myblockchainicli query distribution rewards $(./myblockchainicli keys show my-validator-key -a)'

# Monitor validator status
watch -n 5 './myblockchainicli query staking validator <your-validator-addr>'

# Monitor node status
watch -n 2 'curl -s http://localhost:26657/status | jq .result.sync_info'
```

---

## Resources

- [Cosmos Testnets](https://github.com/cosmos/testnets)
- [Validator Guide](https://hub.cosmos.network/main/validators/overview.html)
- [Testnet Faucets](https://github.com/cosmos/testnet-faucet)
- [Block Explorers](https://cosmos.directory/)
- [Community Discord](https://discord.gg/cosmosnetwork)

---

## Support

For issues:
1. Check testnet README on GitHub
2. Search GitHub issues for your error
3. Ask in Cosmos Discord #validators channel
4. Check Cosmos docs at https://docs.cosmos.network/

Happy validating! 🚀
