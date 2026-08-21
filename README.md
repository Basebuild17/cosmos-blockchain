# Cosmos Blockchain - Beginner Edition

A simple, commercial-ready blockchain built on Cosmos SDK with beginner-friendly business logic.

## Features

- **Account Management**: Create and manage accounts with balances
- **Token Transfers**: Send tokens between accounts
- **Simple Registry**: Store and retrieve key-value data on-chain
- **Messages & Events**: Full event tracking for all transactions

## Project Structure

```
cosmos-blockchain/
├── app/                    # Application setup
│   ├── app.go             # Main app initialization
│   └── config.go          # Chain configuration
├── cmd/                   # CLI commands
│   ├── myblockchaind/     # Daemon (node)
│   └── myblockchaiencli/  # Client (CLI)
├── x/                     # Custom modules (business logic)
│   └── registry/          # Registry module for storing data
│       ├── keeper/        # State management
│       ├── types/         # Data structures & messages
│       ├── module.go      # Module setup
│       └── handler.go     # Message handlers
├── proto/                 # Protocol buffers (data serialization)
├── go.mod                 # Go dependencies
└── Makefile              # Build commands
```

## Getting Started

### Prerequisites
- Go 1.21+
- Cosmos SDK v0.47+

### Installation

```bash
# Clone the repo
git clone https://github.com/Basebuild17/cosmos-blockchain.git
cd cosmos-blockchain

# Install dependencies
go mod download

# Build the blockchain
make build
```

### Running a Node

```bash
# Initialize a new chain
./myblockchain init mynode --chain-id=mychain-1

# Start the node
./myblockchain start
```

### Interacting with the Blockchain

```bash
# Create an account
./myblockchaincli keys add myaccount

# Send a transaction
./myblockchaincli tx registry set-data \
  --from myaccount \
  --chain-id=mychain-1 \
  --key mykey \
  --value myvalue

# Query data
./myblockchaincli query registry get-data mykey
```

## Business Logic: Registry Module

The simple registry module allows you to:
1. **Set Data**: Store key-value pairs on-chain
2. **Get Data**: Retrieve stored data by key
3. **List Data**: View all stored data

Each operation is tracked as a transaction and generates events.

## Next Steps

- Add more complex business logic to the registry module
- Implement custom message types
- Add validation rules
- Create advanced queries

## Resources

- [Cosmos SDK Documentation](https://docs.cosmos.network/)
- [Tendermint Documentation](https://docs.tendermint.com/)
- [Protocol Buffers Guide](https://developers.google.com/protocol-buffers)

## License

MIT
