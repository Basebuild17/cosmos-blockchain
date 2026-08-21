# Validator Setup Guide

## Overview

This guide will help you set up a validator node for your Cosmos blockchain. There are three main deployment scenarios covered:

1. **Local Single Node** - Test locally on your machine
2. **Local Testnet** - Multiple nodes on your machine
3. **Join Existing Cosmos Testnet** - Connect to an existing Cosmos Hub testnet

---

## Prerequisites

### System Requirements

- **OS**: Linux (Ubuntu 20.04+ recommended), macOS, or Windows (WSL2)
- **CPU**: 2+ cores
- **RAM**: 4GB minimum (8GB+ recommended)
- **Disk**: 20GB+ SSD
- **Network**: Stable internet connection, open ports 26656-26657

### Software Requirements

```bash
# Install Go 1.21+
sudo wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version  # Verify installation

# Install Git
sudo apt install git

# Install Make
sudo apt install make
```

---

## Part 1: Local Single Node Validator

Start here to test your blockchain locally.

### Step 1: Build the Blockchain

```bash
cd cosmos-blockchain
go mod download
go mod tidy
make build
make build-cli
```

You should now have two binaries:
- `myblockaind` - The blockchain daemon
- `myblockchainicli` - The CLI client

### Step 2: Initialize Your Node

```bash
# Create a new validator account
./myblockchainicli keys add validator-key

# Save your mnemonic phrase somewhere safe!
# This is your private key recovery phrase.
```

### Step 3: Initialize Blockchain Data

```bash
# Initialize the chain with your node name
./myblockaind init validator-node --chain-id=mychain-1

# This creates:
# ~/.myblockchain/config/  - Configuration files
# ~/.myblockchain/data/    - Blockchain data
```

### Step 4: Copy Genesis File

```bash
# Copy the template genesis file
cp genesis/genesis.json ~/.myblockchain/config/genesis.json

# Or use the chain's genesis file if available
```

### Step 5: Get Your Node Key

```bash
# Display your node's public key (needed for creating validator)
./myblockchainicli keys show validator-key --bech32-prefix cosmos

# Output example:
# address: cosmos1abc123...
# pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"..."}'
```

### Step 6: Create Gentx (Genesis Transaction)

A gentx is a transaction that creates your validator and adds it to the genesis block.

```bash
# Create a self-delegation transaction
./myblockaind gentx validator-key 1000000stake --chain-id=mychain-1

# This creates a file in:
# ~/.myblockchain/config/gentx/
```

### Step 7: Collect Genesis Transactions (For Multi-node)

If this is the only node, skip to Step 8. For multi-node setup, collect all gentx files:

```bash
# Collect all gentx files into genesis
./myblockaind collect-gentxs
```

### Step 8: Validate Genesis File

```bash
# Check that your genesis file is valid
./myblockaind validate-genesis

# Should output: "genesis valid"
```

### Step 9: Start Your Validator

```bash
# Start the blockchain node
./myblockaind start

# You should see output like:
# 10:30AM INF Tendermint started height=1
# 10:30AM INF produced block
```

### Step 10: Verify in Another Terminal

```bash
# Check your account balance
./myblockchainicli query bank balances cosmos1abc123...

# Query registry module
./myblockchainicli query registry list-data

# Check validator status
./myblockchainicli query staking validators
```

---

## Part 2: Local Testnet (Multiple Nodes)

Set up multiple validator nodes on your machine.

### Architecture

We'll create 3 validators:
- `validator-1` (port 26656)
- `validator-2` (port 26666)
- `validator-3` (port 26676)

### Step 1: Create Data Directories

```bash
# Create separate homes for each node
mkdir -p ~/.myblockchain-testnet

for i in 1 2 3; do
  mkdir -p ~/.myblockchain-testnet/validator-$i
export TMHOME=~/.myblockchain-testnet/validator-$i
  ./myblockaind init validator-$i --chain-id=mychain-testnet
done
```

### Step 2: Create Keys for Each Validator

```bash
# Create keys (save mnemonics!)
for i in 1 2 3; do
  ./myblockchainicli keys add validator-$i --keyring-backend test --home ~/.myblockchain-testnet/validator-$i
done
```

### Step 3: Generate Gentx Files

```bash
# Create gentx for each validator
for i in 1 2 3; do
  export TMHOME=~/.myblockchain-testnet/validator-$i
  ./myblockaind gentx validator-$i 1000000stake --chain-id=mychain-testnet --keyring-backend test --home ~/.myblockchain-testnet/validator-$i
done
```

### Step 4: Copy Gentx Files to Node 1

```bash
# Copy all gentx to node-1's config
cp ~/.myblockchain-testnet/validator-2/config/gentx/* ~/.myblockchain-testnet/validator-1/config/gentx/
cp ~/.myblockchain-testnet/validator-3/config/gentx/* ~/.myblockchain-testnet/validator-1/config/gentx/
```

### Step 5: Collect Gentx on Node 1

```bash
export TMHOME=~/.myblockchain-testnet/validator-1
./myblockaind collect-gentxs
```

### Step 6: Copy Genesis to All Nodes

```bash
# Copy the combined genesis to all nodes
for i in 2 3; do
  cp ~/.myblockchain-testnet/validator-1/config/genesis.json ~/.myblockchain-testnet/validator-$i/config/
done
```

### Step 7: Configure Persistent Peers

Get the node IDs:

```bash
# Get node IDs
for i in 1 2 3; do
  echo "Validator $i:"
  TMHOME=~/.myblockchain-testnet/validator-$i ./myblockaind tendermint show-node-id
done
```

Update each node's `config.toml`:

```bash
# Edit validator-1 config
nano ~/.myblockchain-testnet/validator-1/config/config.toml

# Find the `persistent_peers` line and add (with your actual node IDs):
persistent_peers = "node-id-2@127.0.0.1:26666,node-id-3@127.0.0.1:26676"
```

Repeat for validator-2 and validator-3 (pointing to the others).

### Step 8: Update Ports in config.toml

For each node, update the ports to avoid conflicts:

```bash
# validator-1 (default)
laddr = "tcp://127.0.0.1:26656"
rpc_laddr = "tcp://127.0.0.1:26657"

# validator-2
laddr = "tcp://127.0.0.1:26666"
rpc_laddr = "tcp://127.0.0.1:26667"

# validator-3
laddr = "tcp://127.0.0.1:26676"
rpc_laddr = "tcp://127.0.0.1:26677"
```

### Step 9: Start All Validators

Open three separate terminals:

```bash
# Terminal 1
export TMHOME=~/.myblockchain-testnet/validator-1
./myblockaind start

# Terminal 2
export TMHOME=~/.myblockchain-testnet/validator-2
./myblockaind start

# Terminal 3
export TMHOME=~/.myblockchain-testnet/validator-3
./myblockaind start
```

### Step 10: Verify Network

```bash
# Check validator set
export TMHOME=~/.myblockchain-testnet/validator-1
./myblockchainicli query staking validators

# Check network status
curl http://localhost:26657/status
```

---

## Part 3: Join Existing Cosmos Testnet

Connect your validator to an existing Cosmos testnet.

### Available Testnets

**Cosmos Hub Testnets:**
- [Cosmos Hub Testnet List](https://github.com/cosmos/testnets)
- Official: Check the repo for active testnets

**Other Popular Testnets:**
- Osmosis testnet
- Akash testnet
- Juno testnet

### Step 1: Get Testnet Information

For this example, we'll use a generic Cosmos testnet. Replace values with actual testnet info:

```bash
# Download the testnet's genesis file
wget https://example-testnet.com/genesis.json -O ~/.myblockchain/config/genesis.json

# Or from GitHub
wget https://raw.githubusercontent.com/cosmos/testnets/master/v12-testnet/genesis.json \
  -O ~/.myblockchain/config/genesis.json

# Verify the chain ID and genesis hash
sha256sum ~/.myblockchain/config/genesis.json
# Compare with official testnet documentation
```

### Step 2: Initialize Your Node

```bash
# Initialize with testnet chain ID
./myblockaind init validator-testnet --chain-id=<testnet-chain-id>

# Replace <testnet-chain-id> with actual testnet ID (e.g., theta-testnet-001)
```

### Step 3: Replace Genesis File

```bash
# Replace with testnet genesis
cp ~/path/to/testnet-genesis.json ~/.myblockchain/config/genesis.json

# Validate
./myblockaind validate-genesis
```

### Step 4: Configure Seeds and Persistent Peers

Edit `~/.myblockchain/config/config.toml`:

```toml
# Add seed nodes (from testnet docs)
seeds = "seed1@seed1.example.com:26656,seed2@seed2.example.com:26656"

# Add persistent peers (from testnet docs)
persistent_peers = "peer1@peer1.example.com:26656,peer2@peer2.example.com:26656"

# Enable state sync (optional, faster syncing)
[statesync]
enable = true
rpc_servers = "rpc1.example.com:26657,rpc2.example.com:26657"
trust_height = 12345
trust_hash = "ABCDEF123456..."
trust_period = "168h"
```

### Step 5: Configure App Settings

Edit `~/.myblockchain/config/app.toml`:

```toml
# Enable API
[api]
enable = true
address = "tcp://0.0.0.0:1317"

# Enable gRPC
[grpc]
enable = true
address = "0.0.0.0:9090"

# Minimum gas price
minimum-gas-prices = "0.025stake"
```

### Step 6: Start Syncing

```bash
# Start the node (this will sync from peers)
./myblockaind start

# Monitor sync progress
watch -n 2 'curl -s http://localhost:26657/status | jq .result.sync_info'
```

Wait for `"catching_up": false` to indicate full sync.

### Step 7: Create Validator Account

Once synced:

```bash
# Create or import validator key
./myblockchainicli keys add validator-key

# Get testnet tokens (from faucet)
# Visit the testnet's faucet website or use:
./myblockchainicli query bank balances cosmos1...
```

### Step 8: Create Validator

```bash
# Create validator transaction
./myblockchainicli tx staking create-validator \
  --amount=1000000stake \
  --pubkey=$(./myblochaind tendermint show-validator) \
  --moniker="my-validator" \
  --chain-id=<testnet-chain-id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --from=validator-key \
  --fees=1000stake
```

### Step 9: Verify Validator Status

```bash
# Check validator set
./myblockchainicli query staking validators

# Check your specific validator
./myblockchainicli query staking validator <your-validator-address>

# Check voting power
./myblockchainicli query staking validator --output=json | jq .tokens
```

---

## Common Tasks

### Check Node Status

```bash
# Status
curl http://localhost:26657/status | jq

# Network info
curl http://localhost:26657/net_info | jq

# Latest block
curl http://localhost:26657/block | jq
```

### Unjail Validator (if slashed)

```bash
./myblockchainicli tx slashing unjail \
  --from=validator-key \
  --chain-id=mychain-1 \
  --fees=1000stake
```

### Increase Validator Stake

```bash
# Self-delegate more tokens
./myblockchainicli tx staking delegate \
  <validator-address> \
  1000000stake \
  --from=validator-key \
  --chain-id=mychain-1 \
  --fees=1000stake
```

### Query Validator Rewards

```bash
./myblockchainicli query distribution rewards <your-address> \
  --chain-id=mychain-1
```

### Withdraw Rewards

```bash
./myblockchainicli tx distribution withdraw-all-rewards \
  --from=validator-key \
  --chain-id=mychain-1 \
  --fees=1000stake
```

---

## Troubleshooting

### Node won't start

```bash
# Check genesis validity
./myblockaind validate-genesis

# Check config
./myblochaind config get validator

# Check logs
tail -f ~/.myblockchain/logs/node.log
```

### Not syncing

```bash
# Check peers
curl http://localhost:26657/net_info | jq '.result.peers | length'

# If 0 peers, check seeds in config.toml
# Restart with:
./myblockaind unsafe-reset-all
./myblockaind start
```

### Out of disk space

```bash
# Prune old blocks (keep only recent)
./myblochaind prune default

# Or reset and resync
./myblochaind unsafe-reset-all
./myblochaind start
```

### High memory usage

```bash
# Edit app.toml
# Reduce: halt-height, halt-time, iavl-cache-size
# Enable pruning
```

---

## Next Steps

1. **Set up monitoring** with Prometheus + Grafana
2. **Enable snapshots** for faster syncing
3. **Configure sentry architecture** for production
4. **Set up key management** with Horcrux or similar
5. **Join governance** - participate in chain decisions

---

## Resources

- [Cosmos SDK Docs](https://docs.cosmos.network/)
- [Tendermint Docs](https://docs.tendermint.com/)
- [Validator Best Practices](https://github.com/cosmos/cosmos-sdk/blob/main/docs/guide/run_node.md)
- [Testnet Info](https://github.com/cosmos/testnets)

---

**Need help?** Check GETTING_STARTED.md or open an issue on GitHub!
