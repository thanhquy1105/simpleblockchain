package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	port := flag.Uint("port", 5000, "TCP port Number for Blockchain Server ")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))
	app.Run()
	// walletMiner := wallet.NewWallet()
	// walletA := wallet.NewWallet()
	// walletB := wallet.NewWallet()

	// t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(),
	// 	walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	// blockchain := blockchain.NewBlockchain(walletMiner.BlockchainAddress())
	// isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(),
	// 	walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature())

	// fmt.Println("Added? ", isAdded)

	// blockchain.Mining()
	// blockchain.Print()

	// fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	// fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	// fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletMiner.BlockchainAddress()))
}
