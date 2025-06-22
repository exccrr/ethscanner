package scanner

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
	}

	fmt.Printf("\nTotal ETH transferred in block: %.6f ETH\n", weiToEth(totalValue))
}

func weiToEth(wei *big.Int) float64 {
	f := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(f, big.NewFloat(math.Pow10(18)))
	val, _ := ethValue.Float64()
	return val
}
