package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// "github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func init() {
	log.SetPrefix("Go-Eth: ")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func main() {

	// infuraURL := os.Getenv("INFURA_URL")
	hardhatURL := os.Getenv("HARDHAT_URL")

	client, err := ethclient.DialContext(context.Background(), hardhatURL)
	if err != nil {
		log.Fatal("ERROR: failed to load Ethereum client")
	}

	defer client.Close()

	// Using `nil` will fetch the most recent block?
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("ERROR: failed to load most recent block")
	}

	fmt.Printf("\n*** >>> Chain URL: %s", hardhatURL)
	fmt.Printf("\n*** >>> Block Number: %d\n", block.Number())
	fmt.Println()
}
