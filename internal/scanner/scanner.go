package scanner

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ScanLatestBlock(client *ethclient.Client) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block number: ", header.Number.String())

	block, err := client.BlockByNumber(context.Background(), header.Number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Block hash: %s\n", block.Hash().Hex())
	fmt.Printf("Transactions: %d\n", len(block.Transactions()))

	for i, tx := range block.Transactions() {
		printTx(tx, i)
		if i >= 4 {
			break
		}
	}
}

func printTx(tx *types.Transaction, index int) {
	fmt.Printf("TX #%d: %s\n", index+1, tx.Hash().Hex())
}
