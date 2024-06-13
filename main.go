package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	log.SetPrefix("Go-Eth: ")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func main() {
	log.Println("online and working fine")
}
