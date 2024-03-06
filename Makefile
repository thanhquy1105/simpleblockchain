blockchain_server:
	go run ./blockchain_server/main.go ./blockchain_server/blockchain_server.go -port 5001

wallet_server:
	go run ./wallet_server/main.go ./wallet_server/wallet_server.go -port 8080

.PHONY: blockchain_server wallet_server