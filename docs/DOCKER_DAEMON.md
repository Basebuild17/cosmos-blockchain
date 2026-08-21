# Docker Daemon Setup Guide

This guide helps you install and start the Docker daemon on your system.

## What is Docker Daemon?

The Docker daemon (`dockerd`) is a background service that manages Docker containers, images, networks, and volumes. You need it running to use Docker and Docker Compose.

---

## Installation by Operating System

### Linux (Ubuntu/Debian)

#### Option 1: Official Docker Installation

```bash
# 1. Remove old Docker versions
sudo apt-get remove docker docker-engine docker.io containerd runc

# 2. Update package index
sudo apt-get update

# 3. Install Docker
sudo apt-get install -y docker.io docker-compose

# 4. Verify installation
docker --version
docker-compose --version

# Expected output:
# Docker version 20.10+
# Docker Compose version 1.29+
```

#### Option 2: Docker's Official Repository (Recommended)

```bash
# 1. Install dependencies
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl gnupg lsb-release

# 2. Add Docker GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# 3. Add Docker repository
echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# 4. Install Docker
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

# 5. Verify
docker --version
```

#### Option 3: Using Installation Script

```bash
# Download and run official installation script
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Verify
docker --version
```

---

### macOS

#### Option 1: Docker Desktop (Easiest)

1. **Download Docker Desktop:**
   - Go to: https://www.docker.com/products/docker-desktop
   - Click "Mac"
   - Choose your chip:
     - **Apple Silicon (M1/M2/M3)**: Download `Docker.dmg` (Apple Silicon)
     - **Intel**: Download `Docker.dmg` (Intel Chip)

2. **Install:**
   - Open downloaded `Docker.dmg`
   - Drag Docker icon to Applications folder
   - Wait for copy to complete

3. **Launch:**
   - Open Applications → Docker
   - Enter your password when prompted
   - Wait for "Docker is running" message

4. **Verify:**
   ```bash
   docker --version
   docker-compose --version
   ```

#### Option 2: Homebrew

```bash
# Install Homebrew (if needed)
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install Docker Desktop via Homebrew
brew install docker docker-compose

# Or use Homebrew Cask (includes GUI)
brew install --cask docker

# Verify
docker --version
```

---

### Windows

#### Option 1: Docker Desktop for Windows (Recommended)

1. **Check Windows Version:**
   - Open `Settings` → `System` → `About`
   - Note your OS build number
   - **Windows 11**: Any version works
   - **Windows 10**: Requires version 2004 or later

2. **Download Docker Desktop:**
   - Go to: https://www.docker.com/products/docker-desktop
   - Click "Windows"
   - Download `Docker Desktop Installer.exe`

3. **Install:**
   - Run `Docker Desktop Installer.exe`
   - Follow installation wizard
   - Enable "WSL 2" when prompted (recommended)
   - Restart computer when finished

4. **Launch:**
   - Open Windows Start Menu
   - Search for "Docker Desktop"
   - Click to start
   - Wait for system tray icon

5. **Verify:**
   ```powershell
   docker --version
   docker-compose --version
   ```

#### Option 2: WSL 2 + Docker

If you prefer command-line only:

```powershell
# PowerShell (as Administrator)

# 1. Enable WSL 2
wsl --install

# 2. Restart computer

# 3. In WSL terminal, install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# 4. Start Docker daemon
sudo service docker start

# 5. Verify
docker --version
```

---

## Starting Docker Daemon

### Linux

#### Automatic Start on Boot

```bash
# Enable Docker to start on boot
sudo systemctl enable docker

# Start Docker daemon
sudo systemctl start docker

# Verify it's running
sudo systemctl status docker

# Should show: "Active: active (running)"
```

#### Manual Start

```bash
# Start Docker daemon
sudo systemctl start docker

# Stop Docker daemon
sudo systemctl stop docker

# Restart Docker daemon
sudo systemctl restart docker

# View logs
sudo journalctl -u docker -f
```

### macOS

```bash
# Docker Desktop (GUI):
# 1. Open Applications → Docker
# 2. Wait for "Docker is running" in menu bar

# Or via Terminal:
open -a Docker

# Check if running
docker info

# Should show Docker daemon info
```

### Windows

```powershell
# Docker Desktop automatically starts with Windows
# (if auto-start is enabled)

# Or manually start:
# 1. Open Windows Start Menu
# 2. Search "Docker Desktop"
# 3. Click to launch

# Verify in PowerShell:
docker info
```

---

## Fixing "Cannot Connect to Docker Daemon"

### Problem: "Cannot connect to Docker daemon"

```
error during connect: This error may indicate the docker daemon is not running.
```

### Solution by OS

#### Linux

```bash
# 1. Check if Docker is installed
docker --version

# If not installed, see Installation section above

# 2. Start Docker daemon
sudo systemctl start docker

# 3. Verify it's running
sudo systemctl status docker

# 4. Enable auto-start
sudo systemctl enable docker

# 5. Test
docker ps
```

#### macOS

```bash
# 1. Check if installed
docker --version

# 2. Start Docker Desktop
open -a Docker

# 3. Wait 30 seconds for startup
sleep 30

# 4. Test
docker ps
```

#### Windows

```powershell
# 1. Check if installed
docker --version

# 2. Open Docker Desktop
# Applications → Docker Desktop
# or Search → "Docker Desktop"

# 3. Wait for system tray icon (whale icon)

# 4. Test in PowerShell
docker ps
```

---

## Non-Root Access (Linux Only)

By default, Docker requires `sudo`. To use Docker without `sudo`:

```bash
# 1. Create docker group (usually exists)
sudo groupadd docker

# 2. Add your user to docker group
sudo usermod -aG docker $USER

# 3. Apply new group membership
newgrp docker

# 4. Test without sudo
docker ps

# If still doesn't work, restart Docker
sudo systemctl restart docker
```

**Warning:** Users in the `docker` group can run containers as root. This is a security consideration for production systems.

---

## Docker Compose Installation

### Linux

```bash
# Official method (latest version)
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

# Make executable
sudo chmod +x /usr/local/bin/docker-compose

# Verify
docker-compose --version
```

### macOS

```bash
# If using Docker Desktop: already included
docker-compose --version

# If using Homebrew:
brew install docker-compose

# Or:
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

### Windows

```powershell
# If using Docker Desktop: already included
docker-compose --version

# If using WSL + Docker:
# Inside WSL terminal:
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
```

---

## Verify Installation

```bash
# Check Docker version
docker --version
# Expected: Docker version 20.10 or later

# Check Docker Compose version
docker-compose --version
# Expected: Docker Compose version 1.29 or later

# Check Docker daemon status
docker info
# Should show containers, images, storage info

# Test with hello-world
docker run hello-world
# Should show "Hello from Docker!"
```

---

## Troubleshooting

### Issue: "Cannot connect to Docker daemon"

```bash
# Step 1: Check if daemon is running

# Linux:
sudo systemctl status docker

# macOS:
docker info

# Windows:
Tasklist | findstr Docker  # Check running processes
```

### Issue: Permission Denied

```bash
# Linux (non-root access):
sudo usermod -aG docker $USER
newgrp docker

# Or use sudo:
sudo docker ps
```

### Issue: Port Already in Use

```bash
# If Docker ports are in use:

# Linux:
sudo lsof -i :26657  # Check what's using the port
sudo kill -9 <PID>   # Kill the process

# macOS:
lsof -i :26657
kill -9 <PID>

# Windows:
netstat -ano | findstr :26657
taskkill /PID <PID> /F
```

### Issue: Disk Space

```bash
# Check Docker disk usage
docker system df

# Clean up unused images
docker image prune

# Clean up unused containers
docker container prune

# Clean up everything
docker system prune -a
```

### Issue: Memory/CPU Issues

#### macOS/Windows Docker Desktop

1. Right-click Docker icon in system tray
2. Select "Preferences" (macOS) or "Settings" (Windows)
3. Go to "Resources"
4. Adjust CPU, Memory, Disk limits
5. Click "Apply & Restart"

#### Linux

```bash
# Check Docker daemon resource usage
docker stats

# Limit container resources in docker-compose.yml:
# services:
#   validator:
#     deploy:
#       resources:
#         limits:
#           cpus: '2'
#           memory: 4G
```

---

## Quick Commands

```bash
# Start Docker daemon
sudo systemctl start docker          # Linux
open -a Docker                       # macOS
# Windows: GUI or Taskbar icon

# Stop Docker daemon
sudo systemctl stop docker           # Linux
# macOS: Click Docker icon → Quit Docker Desktop
# Windows: Click Docker icon → Quit

# Check status
sudo systemctl status docker         # Linux
docker info                          # All OS

# Enable auto-start (Linux)
sudo systemctl enable docker

# View logs (Linux)
sudo journalctl -u docker -f

# Restart daemon
sudo systemctl restart docker        # Linux
```

---

## Next Steps

Once Docker is running:

```bash
# Verify with hello-world
docker run hello-world

# Start your validator
cd cosmos-blockchain
bash scripts/start-validator.sh

# Monitor
bash scripts/monitor-validator.sh
```

---

## Resources

- **Official Docker Docs:** https://docs.docker.com/
- **Install Guide:** https://docs.docker.com/get-docker/
- **Docker Compose:** https://docs.docker.com/compose/
- **Troubleshooting:** https://docs.docker.com/config/troubleshoot/

---

## Common Issues Summary

| Issue | Solution |
|-------|----------|
| "daemon not running" | Start Docker daemon (see above) |
| "permission denied" | Use `sudo` or add user to docker group |
| "port in use" | Change ports in docker-compose.yml |
| "out of memory" | Increase Docker resource limits |
| "image won't build" | Check disk space, rebuild with --no-cache |

---

**Docker daemon setup complete!** 🐳

Now run:
```bash
bash scripts/start-validator.sh
```
