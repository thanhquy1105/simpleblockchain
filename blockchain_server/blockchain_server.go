package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/thanhquy1105/simpleblockchain/blockchain"
	"github.com/thanhquy1105/simpleblockchain/utils"
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

func (bcs *BlockchainServer) Transactions(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		transactions := bc.TransactionPool()
		m, _ := json.Marshal(struct {
			Transactions []*blockchain.Transaction `json:"transactions"`
			Length       int                       `json:"length"`
		}{
			Transactions: transactions,
			Length:       len(transactions),
		})
		io.WriteString(w, string(m[:]))
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t blockchain.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromString(*t.Signature)
		bc := bcs.GetBlockchain()
		isCreated := bc.CreateTransaction(*t.SenderBlockchainAddress, *t.RecipientBlockchainAddress, *t.Value, publicKey, signature)

		w.Header().Add("Content-Type", "application/json")
		var m []byte
		if !isCreated {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("fail 2")
		} else {
			w.WriteHeader(http.StatusCreated)
			m = utils.JsonStatus("success")
		}
		io.WriteString(w, string(m))
	default:
		log.Println("ERROR: Invalid HTTP method")
	}
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	http.HandleFunc("/transactions", bcs.Transactions)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
