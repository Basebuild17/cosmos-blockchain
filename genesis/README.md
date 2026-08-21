# Genesis Files

This directory contains genesis file templates and configurations.

## Files

- **genesis.json** - Template genesis file with all required modules

## Using Genesis Files

### Option 1: Use Template Genesis

```bash
# Copy template to your node
cp genesis/genesis.json ~/.myblockchain/config/genesis.json

# Validate
./myblochaind validate-genesis
```

### Option 2: Download Testnet Genesis

```bash
# For Cosmos Hub testnet
wget https://raw.githubusercontent.com/cosmos/testnets/master/theta-testnet-001/genesis.json \
  -O ~/.myblockchain/config/genesis.json
```

### Option 3: Create Custom Genesis

Edit genesis.json before using:

```bash
# Update chain ID
jq '.chain_id = "my-custom-chain"' genesis.json > genesis-new.json

# Update initial parameters
jq '.consensus_params.block.max_gas = "100000000"' genesis.json > genesis-new.json

# Copy to node
cp genesis-new.json ~/.myblockchain/config/genesis.json
```

## Genesis File Structure

### Key Sections

1. **genesis_time** - Network start time
2. **chain_id** - Unique chain identifier
3. **consensus_params** - Tendermint consensus settings
4. **app_state** - Module initial state:
   - **auth** - Account settings
   - **bank** - Token balances and metadata
   - **registry** - Registry module data
   - **staking** - Validator settings
   - **slashing** - Slash penalties
   - **mint** - Inflation parameters
   - **distribution** - Reward distribution
   - **gov** - Governance settings
   - **genutil** - Genesis utilities (gentx collection)

## Customizing for Your Network

```bash
# Update chain ID
jq '.chain_id = "mynetwork-1"' genesis.json > genesis-custom.json

# Update min stake for validators
jq '.app_state.staking.params.bond_denom = "customcoin"' genesis-custom.json > genesis.json

# Update inflation
jq '.app_state.mint.params.inflation_max = "0.10"' genesis.json > genesis-custom.json

# Update governance parameters
jq '.app_state.gov.deposit_params.min_deposit[0].amount = "50000000"' genesis-custom.json > genesis.json
```

## Validating Genesis

```bash
# Full validation
./myblochaind validate-genesis

# JSON schema validation
jq . genesis.json > /dev/null && echo "Valid JSON"
```

## Next Steps

1. See VALIDATOR_SETUP.md for creating validators
2. See TESTNET_CONNECTION.md for joining existing testnets
3. Customize genesis parameters as needed
4. Collect gentx files from all validators
5. Start the network
