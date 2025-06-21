package client

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthclient(rpcUrl string) *ethclient.Client {
	client, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to Eth node: %v", err)
	}
	return client
}
