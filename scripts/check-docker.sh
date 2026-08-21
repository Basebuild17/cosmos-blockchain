#!/bin/bash

# Check and start Docker daemon
# This script verifies Docker is installed and running

set -e

echo "================================"
echo "Docker Daemon Checker"
echo "================================"
echo ""

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# Detect OS
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    OS="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
    OS="macos"
elif [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "cygwin" ]]; then
    OS="windows"
else
    OS="unknown"
fi

echo "Detected OS: $OS"
echo ""

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check Docker installation
echo "[1/3] Checking Docker installation..."
if command_exists docker; then
    DOCKER_VERSION=$(docker --version)
    echo -e "${GREEN}✓ Docker installed: $DOCKER_VERSION${NC}"
else
    echo -e "${RED}✗ Docker is not installed${NC}"
    echo ""
    echo "Installation instructions:"
    if [ "$OS" = "linux" ]; then
        echo "  Linux: https://docs.docker.com/engine/install/ubuntu/"
        echo "  Quick: curl -fsSL https://get.docker.com -o get-docker.sh && sudo sh get-docker.sh"
    elif [ "$OS" = "macos" ]; then
        echo "  macOS: https://docs.docker.com/desktop/install/mac-install/"
        echo "  Or: brew install --cask docker"
    elif [ "$OS" = "windows" ]; then
        echo "  Windows: https://docs.docker.com/desktop/install/windows-install/"
    fi
    exit 1
fi

# Check Docker Compose installation
echo ""
echo "[2/3] Checking Docker Compose installation..."
if command_exists docker-compose; then
    COMPOSE_VERSION=$(docker-compose --version)
    echo -e "${GREEN}✓ Docker Compose installed: $COMPOSE_VERSION${NC}"
else
    echo -e "${YELLOW}⚠ Docker Compose is not installed${NC}"
    echo "Installing Docker Compose..."
    if [ "$OS" = "linux" ]; then
        COMPOSE_URL="https://github.com/docker/compose/releases/latest/download/docker-compose-Linux-x86_64"
        sudo curl -L $COMPOSE_URL -o /usr/local/bin/docker-compose
        sudo chmod +x /usr/local/bin/docker-compose
        echo -e "${GREEN}✓ Docker Compose installed${NC}"
    elif [ "$OS" = "macos" ]; then
        echo "Please run: brew install docker-compose"
        exit 1
    fi
fi

# Check if Docker daemon is running
echo ""
echo "[3/3] Checking Docker daemon status..."
if docker info > /dev/null 2>&1; then
    echo -e "${GREEN}✓ Docker daemon is running${NC}"
else
    echo -e "${RED}✗ Docker daemon is not running${NC}"
    echo ""
    echo "Starting Docker daemon..."
    echo ""
    
    if [ "$OS" = "linux" ]; then
        echo -e "${BLUE}Command:${NC} sudo systemctl start docker"
        sudo systemctl start docker
        echo ""
        echo "To enable auto-start:"
        echo "  sudo systemctl enable docker"
        
    elif [ "$OS" = "macos" ]; then
        echo -e "${BLUE}Opening Docker Desktop...${NC}"
        open -a Docker
        echo ""
        echo "Please wait 30 seconds for Docker to start..."
        sleep 30
        
    elif [ "$OS" = "windows" ]; then
        echo -e "${BLUE}Please start Docker Desktop:${NC}"
        echo "  1. Click Windows Start Menu"
        echo "  2. Search 'Docker Desktop'"
        echo "  3. Click to launch"
        echo "  4. Wait for whale icon in system tray"
        echo ""
        read -p "Press Enter once Docker Desktop is running..."
    fi
fi

# Final verification
echo ""
echo "[VERIFICATION] Running docker info..."
echo ""
docker info | head -20

echo ""
echo -e "${GREEN}================================${NC}"
echo -e "${GREEN}✓ Docker is ready!${NC}"
echo -e "${GREEN}================================${NC}"
echo ""
echo "Next step: Run your validator"
echo "  cd cosmos-blockchain"
echo "  bash scripts/start-validator.sh"
echo ""
