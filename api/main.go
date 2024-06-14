package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// ------------------------------------------------------------

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{
		port: port,
	}
}

func (bcs *BlockchainServer) PrintChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:

		w.Header().Add("Content-Type", "application/json")

		fmt.Println("Printing chain")

		// ----------------------------------------------------------
		w.Write([]byte("Printing chain"))

	default:
		http.Error(w, "ERROR: Invalid HTTP Method", http.StatusBadRequest)
	}
}

func (bcs *BlockchainServer) Run() {

	http.HandleFunc("/printchain", bcs.PrintChain)

	hostURL := fmt.Sprintf("0.0.0.0:%d", bcs.port)
	fmt.Println("Blockchain HTTP Server is live @:", hostURL)
	log.Fatal(http.ListenAndServe(hostURL, nil))
}

// ------------------------------------------------------------
func init() {
	log.SetPrefix("Blockchain API: ")
}

func main() {

	defer os.Exit(0)

	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()

	app := NewBlockchainServer(uint16(*port))

	app.Run()
}
