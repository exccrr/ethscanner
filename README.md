# Ethereum Block Scanner (Go + Infura)

A lightweight CLI tool written in Go that connects to the Ethereum blockchain and retrieves the latest block information using Infura RPC.

It fetches:
- Latest block number
- Block hash
- Block timestamp (converted to human-readable time)
- Sender and receiver (From / To) for first 5 transactions
- ETH value transferred in each transaction
- Total ETH transferred in the block
- First 5 transaction hashes
- ERC-20 Transfer events from logs (token transfers)
- Number of transactions in the block

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
git clone https://github.com/exccrr/ethscanner.git
cd ethscanner
```

### 2. Create .env file with your INFURA_URL
```
INFURA_URL=https://mainnet.infura.io/v3/YOUR_KEY
```
### 3. Run main
```bash
go run cmd/main.go
```
