blockchain_server0:
	go run ./blockchain_server/main.go ./blockchain_server/blockchain_server.go -port 5000

blockchain_server1:
	go run ./blockchain_server/main.go ./blockchain_server/blockchain_server.go -port 5001

blockchain_server2:
	go run ./blockchain_server/main.go ./blockchain_server/blockchain_server.go -port 5002

wallet_server:
	go run ./wallet_server/main.go ./wallet_server/wallet_server.go -port 8080

.PHONY: blockchain_server0 wallet_server blockchain_server1 blockchain_server2