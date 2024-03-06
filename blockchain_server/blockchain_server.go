package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/thanhquy1105/simpleblockchain/blockchain"
	"github.com/thanhquy1105/simpleblockchain/wallet"
)

var cache map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func (bcs *BlockchainServer) GetBlockchain() *blockchain.Blockchain {
	bc, ok := cache["Blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = blockchain.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		cache["Blockchain"] = bc
		log.Printf("private_key	%v", minersWallet.PrivateKeyStr())
		log.Printf("public_key	%v", minersWallet.PublicKeyStr())
		log.Printf("blockchain_address	%v", minersWallet.BlockchainAddress())
	}
	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		{
			w.Header().Add("Content-Type", "application/json")
			bc := bcs.GetBlockchain()
			m, _ := bc.MarshalJSON()
			io.WriteString(w, string(m[:]))
		}
	default:
		log.Println("ERROR: Invalid HTTP method")
	}
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
