# Ethereum Block Scanner (Go + Infura)

A lightweight CLI tool written in Go that connects to the Ethereum blockchain and retrieves the latest block information using Infura RPC.

It fetches:
- latest block number
- block hash
- from and to
- timestamp
- total value
- number of transactions
- first 5 transaction hashes

---

## 🛠 Tech Stack

- **Golang** (1.20+)
- [go-ethereum](https://github.com/ethereum/go-ethereum)
- [Infura](https://infura.io) RPC endpoint
- [godotenv](https://github.com/joho/godotenv) for environment configuration

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/ethscanner.git
cd ethscanner
go run cmd/main.go