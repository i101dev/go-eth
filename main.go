package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var (
	INFURA_URL  string
	HARDHAT_URL string
	SEPOLIA_URL string

	PASSWORD string
)

const (
	WALLET_DIR = "./wallet"

	KEY_1 = "./wallet/key1"
	KEY_2 = "./wallet/key2"
)

func init() {
	log.SetPrefix("Go-Eth: ")
	log.SetFlags(0)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	PASSWORD = os.Getenv("PASSWORD")

	INFURA_URL = os.Getenv("INFURA_URL")
	HARDHAT_URL = os.Getenv("HARDHAT_URL")
	SEPOLIA_URL = os.Getenv("SEPOLIA_URL")
}
func main() {

	client, err := ethclient.DialContext(context.Background(), HARDHAT_URL)
	if err != nil {
		log.Fatal("ERROR: failed to load Ethereum client")
	}

	defer client.Close()

	// getLastBlock(client)
	// getBalance(client)
	// generateWallet()
	// generateEncryptedWallet()

	derivePrivateKey(KEY_1, PASSWORD)
	fmt.Println(strings.Repeat("-", 40))
	derivePrivateKey(KEY_2, "123")
}

func getLastBlock(client *ethclient.Client) {

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("ERROR: failed to load most recent block")
	}

	fmt.Printf("\n*** >>> Chain URL: %s", HARDHAT_URL)
	fmt.Printf("\n*** >>> Block Number: %d\n", block.Number())
	fmt.Println()
}

func getBalance(client *ethclient.Client) {

	addr := "0xcd3B766CCDd6AE721141F452C550Ca635964ce71"
	acct := common.HexToAddress(addr)
	balance, err := client.BalanceAt(context.Background(), acct, nil)
	if err != nil {
		log.Fatal("ERROR: failed to fetch account balance")
	}

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	etherBalance := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))

	x, _ := etherBalance.Float64()

	fmt.Printf("\n*** >>> Balance: %.3f\n", x)
	fmt.Println()
}

func generateWallet() {
	//
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("ERROR: failed to generate private key")
	}
	privKey := crypto.FromECDSA(pvk)
	fmt.Println("Private Key: ", hexutil.Encode(privKey))

	pubKey := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println("Private Key: ", hexutil.Encode(pubKey))

	addr := crypto.PubkeyToAddress(pvk.PublicKey).Hex()
	fmt.Println("Address: ", addr)
}

func generateEncryptedWallet() {

	keyStore := keystore.NewKeyStore(WALLET_DIR, keystore.StandardScryptN, keystore.StandardScryptP)

	acct, err := keyStore.NewAccount(PASSWORD)
	if err != nil {
		log.Fatal("ERROR: failed to generate private key")
	}

	fmt.Println(acct.Address)
}

func derivePrivateKey(filePath string, password string) {

	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("ERROR: failed to read file from path")
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal("ERROR: failed to decrypt private key: \n", err)
	}

	privKey := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private Key: ", hexutil.Encode(privKey))

	pubKey := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public Key: ", hexutil.Encode(pubKey))

	addr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println("Address: ", addr)
}
