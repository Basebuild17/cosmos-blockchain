# Getting Started with Cosmos SDK Blockchain

## Welcome!

This guide will help you understand and run your first Cosmos blockchain. Since you're new to Go, I'll explain things step by step.

## Understanding the Structure

### `go.mod`
Think of this like a `package.json` (if you know JavaScript) or `requirements.txt` (Python). It lists all the dependencies your blockchain needs.

### `app/` - The Main Application
- `app.go`: This is like the "core" of your blockchain. It brings together all the modules and sets them up.
- `config.go`: Configuration settings for your chain.

### `x/registry/` - Your Custom Business Logic
This is where YOU add your business logic. Here's what each part does:

- **`types/`**: Defines the data structures
  - `messages.go`: The transactions users can send (SetData, GetData)
  - `keys.go`: Constants for storage
  - `errors.go`: Error definitions
  - `codec.go`: How data is serialized/deserialized
  - `genesis.go`: Initial state when the blockchain starts

- **`keeper/`**: The "manager" that handles state
  - `keeper.go`: Stores and retrieves data from the blockchain
  - `msg_server.go`: Processes incoming messages (transactions)
  - `query_server.go`: Handles read-only requests
  - `genesis.go`: Initializes blockchain state

## How It Works

### The Flow of a Transaction

1. **User creates a message**: "SetData with key=mykey, value=myvalue"
2. **Message gets validated**: Is it well-formed? Does the user have an account?
3. **Handler processes it**: `msg_server.go` calls keeper.SetData()
4. **Data is stored**: The keeper saves data to the blockchain state
5. **Event is emitted**: Other apps can listen to what happened
6. **Transaction is complete**: Recorded in a block forever

### The Registry Module

The Registry is your first business logic. It's very simple:

```
SetData(key, value)  → Stores a key-value pair
GetData(key)         → Retrieves the value for a key
ListData()           → Lists all key-value pairs
```

## Prerequisites

Before you start, install:

1. **Go** (version 1.21+): https://golang.org/dl/
   - Verify: `go version`

2. **Git**: https://git-scm.com/

3. **Make** (optional, but helpful)
   - On Mac: `brew install make`
   - On Linux: `sudo apt install make`

## First Run

### Step 1: Clone and Setup
```bash
git clone https://github.com/Basebuild17/cosmos-blockchain.git
cd cosmos-blockchain
```

### Step 2: Download Dependencies
```bash
go mod download
go mod tidy
```

This downloads all the Cosmos SDK code you need.

### Step 3: Build
```bash
make build
make build-cli
```

Or manually:
```bash
go build -o myblockaind ./cmd/myblockchain
go build -o myblockchainicli ./cmd/myblockchaincli
```

### Step 4: Initialize Your Node
```bash
./myblockaind init mynode --chain-id=mychain-1
```

This creates the configuration files for your blockchain node.

### Step 5: Start the Blockchain
```bash
./myblockaind start
```

You should see output like:
```
Started blockchain with ID: mychain-1
```

### Step 6: Interact (in a new terminal)

Create an account:
```bash
./myblockchainicli keys add myaccount
```

Store data:
```bash
./myblockchainicli tx registry set-data mykey myvalue \
  --from myaccount \
  --chain-id=mychain-1
```

Query data:
```bash
./myblockchainicli query registry get-data mykey
```

## What You Should Know (Go Basics)

### Variables and Types
```go
key := "mykey"           // String
value := "myvalue"       // String
count := 42              // Integer
success := true          // Boolean
```

### Functions
```go
// Function definition
func SetData(key string, value string) error {
    // ...
    return nil
}

// Function call
SetData("name", "Alice")
```

### Error Handling
```go
value, found := k.Keeper.GetData(ctx, msg.Key)
if !found {
    return nil, errors.New("key not found")
}
```

### Structs (like classes)
```go
type Data struct {
    Key   string
    Value string
}
```

## Next Steps

1. **Understand the Registry Module**: Read through `x/registry/types/` and `x/registry/keeper/`
2. **Try Modifying**: Add a new message type, like `DeleteData`
3. **Add Validation**: Make sure keys follow a certain pattern
4. **Add Events**: Emit custom events from your handlers

## Common Issues

### "Command not found: go"
Go isn't installed or not in PATH. Reinstall from golang.org

### "Module not found: github.com/cosmos/cosmos-sdk"
Run `go mod download`

### "Cannot find myblockaind"
The binary wasn't built. Run `make build`

## Testing Your Blockchain Locally

After starting the node:

```bash
# In a new terminal

# Create two accounts
./myblockchainicli keys add alice
./myblockchainicli keys add bob

# Alice sets data
./myblockchainicli tx registry set-data \
  name alice \
  --from alice \
  --chain-id=mychain-1 \
  --broadcast-mode=block

# Query the data
./myblockchainicli query registry get-data name

# List all data
./myblockchainicli query registry list-data
```

## Resources

- **Cosmos SDK Docs**: https://docs.cosmos.network/
- **Go Basics**: https://go.dev/tour/welcome/1
- **Tendermint (Consensus)**: https://docs.tendermint.com/
- **Protobuf**: https://developers.google.com/protocol-buffers

## Getting Help

- Check the `README.md` for quick reference
- Review the comments in the code
- Look at Cosmos documentation for similar examples

## Ready to Extend?

Once you're comfortable with the registry module, try:

1. **Add a Counter Module**: A simple counter that increments/decrements
2. **Add Permissions**: Only certain users can set data
3. **Add Expiration**: Data expires after a certain block height
4. **Add User Accounts**: Track who owns what data

Good luck! 🚀
