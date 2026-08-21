#!/bin/bash

# Create validator account and transactions
# Usage: ./scripts/validator-operations.sh [operation] [args]

set -e

echo "Validator Operations Helper"
echo ""

OPERATION=${1:-help}
CLI="docker-compose exec -T cli myblockchainicli"

case $OPERATION in
    create-key)
        KEY_NAME=${2:-my-validator}
        echo "Creating validator key: $KEY_NAME"
        $CLI keys add $KEY_NAME --keyring-backend test
        echo ""
        echo "✓ Key created"
        echo "Save your mnemonic in a safe place!"
        echo ""
        echo "Your address:"
        $CLI keys show $KEY_NAME -a
        ;;
    
    show-key)
        KEY_NAME=${2:-my-validator}
        echo "Showing key: $KEY_NAME"
        echo ""
        $CLI keys show $KEY_NAME -a
        ;;
    
    list-keys)
        echo "All keys:"
        $CLI keys list
        ;;
    
    balance)
        ADDRESS=${2}
        if [ -z "$ADDRESS" ]; then
            echo "Usage: ./scripts/validator-operations.sh balance <address>"
            echo "Example: ./scripts/validator-operations.sh balance cosmos1abc123..."
            exit 1
        fi
        echo "Balance for $ADDRESS:"
        $CLI query bank balances $ADDRESS
        ;;
    
    gentx)
        KEY_NAME=${2:-my-validator}
        AMOUNT=${3:-1000000stake}
        CHAIN_ID=${4:-mychain-1}
        echo "Generating gentx..."
        echo "Key:      $KEY_NAME"
        echo "Amount:   $AMOUNT"
        echo "Chain:    $CHAIN_ID"
        echo ""
        docker-compose exec -T validator myblockaind gentx $KEY_NAME $AMOUNT --chain-id=$CHAIN_ID
        echo ""
        echo "✓ Gentx generated"
        ;;
    
    create-validator)
        KEY_NAME=${2:-my-validator}
        AMOUNT=${3:-1000000stake}
        CHAIN_ID=${4:-mychain-1}
        echo "Creating validator..."
        echo "Key:      $KEY_NAME"
        echo "Amount:   $AMOUNT"
        echo "Chain:    $CHAIN_ID"
        echo ""
        
        # Get pubkey
        PUBKEY=$(docker-compose exec -T validator myblockaind tendermint show-validator)
        
        echo "Public Key: $PUBKEY"
        echo ""
        
        $CLI tx staking create-validator \
            --amount=$AMOUNT \
            --pubkey=$PUBKEY \
            --moniker="MyValidator" \
            --chain-id=$CHAIN_ID \
            --commission-rate="0.10" \
            --commission-max-rate="0.20" \
            --commission-max-change-rate="0.01" \
            --min-self-delegation="1" \
            --from=$KEY_NAME \
            --fees=5000stake \
            --keyring-backend test
        ;;
    
    query-validator)
        ADDRESS=${2}
        if [ -z "$ADDRESS" ]; then
            echo "Usage: ./scripts/validator-operations.sh query-validator <address>"
            exit 1
        fi
        echo "Validator info:"
        $CLI query staking validator $ADDRESS
        ;;
    
    list-validators)
        echo "All validators:"
        $CLI query staking validators
        ;;
    
    status)
        echo "Node status:"
        curl -s http://localhost:26657/status | jq
        ;;
    
    blocks)
        echo "Latest block:"
        curl -s http://localhost:26657/block | jq
        ;;
    
    peers)
        echo "Connected peers:"
        curl -s http://localhost:26657/net_info | jq .result.peers
        ;;
    
    help|*)
        echo "Usage: ./scripts/validator-operations.sh [operation] [args]"
        echo ""
        echo "Account Operations:"
        echo "  create-key [name]              Create new validator key"
        echo "  show-key [name]                Show specific key"
        echo "  list-keys                      List all keys"
        echo "  balance <address>              Check account balance"
        echo ""
        echo "Validator Operations:"
        echo "  gentx [name] [amount] [chain]  Generate gentx"
        echo "  create-validator [name] [amount] [chain]  Create validator"
        echo "  query-validator <address>      Get validator info"
        echo "  list-validators                List all validators"
        echo ""
        echo "Network Operations:"
        echo "  status                         Show node status"
        echo "  blocks                         Show latest block"
        echo "  peers                          Show connected peers"
        echo ""
        echo "Examples:"
        echo "  ./scripts/validator-operations.sh create-key myvalidator"
        echo "  ./scripts/validator-operations.sh balance cosmos1abc123..."
        echo "  ./scripts/validator-operations.sh list-validators"
        echo "  ./scripts/validator-operations.sh status"
        ;;
esac
