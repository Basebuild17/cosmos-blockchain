#!/usr/bin/env bash
# scripts/test-registry-cli.sh
# Simple smoke-test script for the registry CLI against a local node.
# Usage: ./scripts/test-registry-cli.sh [NODE_RPC]
# Example: ./scripts/test-registry-cli.sh tcp://127.0.0.1:26657

set -euo pipefail

NODE_RPC=${1:-"tcp://127.0.0.1:26657"}
CLI_BIN=./myblockchainicli

echo "Using node: ${NODE_RPC}"

# wait for node to respond
for i in {1..30}; do
  if ${CLI_BIN} --node "${NODE_RPC}" status >/dev/null 2>&1; then
    echo "node is responsive"
    break
  fi
  echo "waiting for node... (${i}/30)"
  sleep 1
done

# list all registry data
echo
echo "=== registry list ==="
${CLI_BIN} query registry list --node "${NODE_RPC}" || echo "list failed"

echo
# try to get a specific key (this may fail if key doesn't exist)
TEST_KEY="example"
echo "=== registry get ${TEST_KEY} ==="
${CLI_BIN} query registry get "${TEST_KEY}" --node "${NODE_RPC}" || echo "get failed or key not found"

echo
# done
echo "smoke test complete"
