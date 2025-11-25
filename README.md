<div align="center">
  <img src="assets/logo.png" alt="XMTP Logo" width="200"/>
  
  # XMTP Network Debugging Environment
  
  A local development and debugging environment for XMTP daemon (xmtpd) nodes
  
  [![XMTP](https://img.shields.io/badge/XMTP-Network-blueviolet)](https://xmtp.org)
</div>

## 📋 Overview

This repository provides a complete Docker Compose setup for running and debugging a local XMTP network. It includes:

- **Blockchain**: Local Anvil chain with XMTP contracts
- **Validation Service**: MLS validation service
- **XMTP Nodes**: Two xmtpd nodes (node-1 and node-2)
- **Gateway**: API gateway for network access
- **CLI Tools**: Command-line tools for local, testnet-dev, and testnet-staging environments
- **Go Client**: Example Go client implementation

## 🚀 Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for client development)

### Starting the Network

```bash
./start
```

This command will:

- Start the local blockchain (Anvil)
- Deploy XMTP contracts
- Launch the MLS validation service
- Start two XMTP nodes
- Start the API gateway

### Stopping the Network

```bash
./stop
```

## 📁 Project Structure

```
.
├── main.yml              # Core services (chain, validation)
├── node-1.yml            # First XMTP node configuration
├── node-2.yml            # Second XMTP node configuration
├── gateway.yml           # Gateway configuration
├── main.env              # Main environment variables
├── node-1.env            # Node 1 environment variables
├── node-2.env            # Node 2 environment variables
├── gateway.env           # Gateway environment variables
├── start                 # Start script
├── stop                  # Stop script
├── environments/         # Environment configurations
│   ├── anvil.json       # Local Anvil configuration
│   ├── testnet-dev.json
│   ├── testnet-staging.json
│   └── testnet.json
├── cli/                  # CLI binaries
│   └── local
└── xmtpd-client-go/     # Go client example
    ├── main.go
    ├── go.mod
    └── go.sum
```

## 🔑 Anvil Accounts

The local development environment uses Anvil's default test accounts. Each service is assigned a specific account:

| Service     | Account # | Address                                      | Private Key                                                          | Balance    |
| ----------- | --------- | -------------------------------------------- | -------------------------------------------------------------------- | ---------- |
| **Gateway** | 0         | `0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266` | `0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80` | 10,000 ETH |
| **Node 1**  | 1         | `0x70997970C51812dc3A010C7d01b50e0d17dc79C8` | `0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d` | 10,000 ETH |
| **Node 2**  | 2         | `0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC` | `0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a` | 10,000 ETH |

**Account Roles:**

- **Account 0 (Gateway)**: Used as the payer account (`XMTPD_PAYER_PRIVATE_KEY`) for transaction fees
- **Account 1 (Node 1)**: Used as the signer account (`XMTPD_SIGNER_PRIVATE_KEY`) for node 1 operations
- **Account 2 (Node 2)**: Used as the signer account (`XMTPD_SIGNER_PRIVATE_KEY`) for node 2 operations

> ⚠️ **Security Warning**: These are well-known test accounts. Never use them on public networks or with real funds!

The complete list of Anvil accounts is available in `environments/anvil-accounts`.

## 🔧 Services

### Chain (Anvil)

- Local Ethereum development node
- Pre-configured with XMTP contracts
- Network mode: host

### Validation Service

- MLS (Messaging Layer Security) validation
- Port: 60051 (mapped from 50051)
- Platform: linux/amd64

### XMTP Nodes

- **Node 1**: Primary XMTP daemon instance
- **Node 2**: Secondary XMTP daemon instance
- Both configured for local development and testing

### Gateway

- API gateway for accessing the XMTP network
- Provides HTTP/gRPC interfaces

## 🛠️ Development

### Using the CLI

The repository includes the xmtpd CLI:

```bash
./cli/local [command]
```

### Go Client Example

A sample Go client is provided in `xmtpd-client-go/`:

```bash
cd xmtpd-client-go
go run main.go
```

## 🌍 Environments

The project supports multiple environments:

- **Local**: Full local development environment with Anvil
- **Testnet Dev**: Development testnet configuration
- **Testnet Staging**: Staging testnet configuration
- **Testnet**: Production testnet configuration

Environment configurations are stored in the `environments/` directory.

## 📝 Configuration

Each component can be configured through its respective environment file:

- `main.env`: Core services configuration
- `node-1.env`: Node 1 specific settings
- `node-2.env`: Node 2 specific settings
- `gateway.env`: Gateway specific settings

## 🐛 Debugging

The Docker Compose setup is configured for easy debugging:

1. All services run with proper networking
2. Logs are accessible via `docker compose logs -f`
3. Individual services can be restarted without affecting others

```bash
# View all logs
docker compose -p xmtp-network logs -f

# View specific service logs
docker compose -p xmtp-network logs -f node-1

# Restart a specific service
docker compose -p xmtp-network restart node-1
```

## 📚 Resources

- [XMTP Documentation](https://xmtp.org/docs)
- [XMTP GitHub](https://github.com/xmtp)
- [MLS Protocol](https://messaginglayersecurity.rocks/)

## 🤝 Contributing

This is a debugging and development environment. Feel free to modify configurations to suit your needs.

## 📄 License

See the XMTP project repositories for license information.

---

<div align="center">
  Made with ❤️ for XMTP debugging and development
</div>
