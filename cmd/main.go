package main

import (
	"ethscanner/config"
	"ethscanner/internal/client"
	"ethscanner/internal/scanner"
)

func main() {
	config.LoadEnv()
	url := config.GetEnv("INFURA_URL")

	eth := client.NewEthclient(url)
	defer eth.Close()

	scanner.ScanLatestBlock(eth)
}
