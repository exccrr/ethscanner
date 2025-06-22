package scanner

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const erc20ABI = `[{
	"anonymous": false,
	"inputs": [
		{"indexed": true, "name": "from", "type": "address"},
		{"indexed": true, "name": "to", "type": "address"},
		{"indexed": false, "name": "value", "type": "uint256"}
	],
	"name": "Transfer",
	"type": "event"
}]`

func ScanLatestBlock(client *ethclient.Client) {
	ctx := context.Background()

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(ctx, header.Number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block number:", block.Number().Uint64())
	fmt.Println("Block hash: ", block.Hash().Hex())
	fmt.Println("Timestamp: ", time.Unix(int64(block.Time()), 0))
	fmt.Println("Transactions:", len(block.Transactions()))

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	signer := types.LatestSignerForChainID(chainID)

	totalValue := big.NewInt(0)

	for i, tx := range block.Transactions() {
		if i >= 5 {
			break
		}

		fmt.Printf("\nTX #%d: %s\n", i+1, tx.Hash().Hex())

		from, err := signer.Sender(tx)
		if err != nil {
			log.Printf("Could not decode sender: %v\n", err)
			continue
		}
		fmt.Printf("From: %s\n", from.Hex())

		to := "<contract creation>"
		if tx.To() != nil {
			to = tx.To().Hex()
		}
		fmt.Printf("To: %s\n", to)

		val := tx.Value()
		totalValue.Add(totalValue, val)
		fmt.Printf("Value: %.6f ETH\n", weiToEth(val))

		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Printf("Cannot get receipt for tx %s: %v\n", tx.Hash().Hex(), err)
			continue
		}

		parseTransferEvents(receipt)
	}

	fmt.Printf("\nTotal ETH transferred in block: %.6f ETH\n", weiToEth(totalValue))
}

func weiToEth(wei *big.Int) float64 {
	f := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(f, big.NewFloat(math.Pow10(18)))
	val, _ := ethValue.Float64()
	return val
}

func parseTransferEvents(receipt *types.Receipt) {
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		log.Println("Failed to parse ABI:", err)
		return
	}

	transferSig := []byte("Transfer(address,address,uint256)")
	transferHash := crypto.Keccak256Hash(transferSig)

	for _, vLog := range receipt.Logs {
		if len(vLog.Topics) == 0 || vLog.Topics[0] != transferHash {
			continue
		}

		from := common.HexToAddress(vLog.Topics[1].Hex())
		to := common.HexToAddress(vLog.Topics[2].Hex())

		decoded := make(map[string]interface{})
		err := parsedABI.UnpackIntoMap(decoded, "Transfer", vLog.Data)

		if err != nil {
			log.Println("Failed to decode event data:", err)
			continue
		}

		val := decoded["value"].(*big.Int)
		fmt.Printf("Token Transfer: %s → %s | %.6f\n", from.Hex(), to.Hex(), weiToEth(val))
	}
}
