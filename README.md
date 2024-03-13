# A simplified blockchain implementation in Golang
This repo simulate how blockchain operates with PoW
## What did I do in this repo:
- Create Blockchain
    - Create Block
    - Create Blockchain
    - Calculate the Hash of a Block
    - Add a transaction
    - Derive the Nonce
    - Calculate the transaction total
- Create Wallet
    - Private Key, Public Key, and ECDSA
    - Create wallet
    - Create Addresses for Blockchain
    - Signatures for Transactions
    - Transaction Verification
- Blockchain Server API
    - Create Blockchain API
    - Create Wallet Server API
    - Send Transaction fron UI
    - Create Mining API
    - Automate mining execution
    - Create currency total API
- Structure of Blockchain Network
    - Search for other Blockchain Nodes
    - Automatic registration of Blockchain Nodes
    - Synchronizing transactions
    - Blockchain verification
    - The longest-chain rule
    - Create consensus API

## How to run this repo
1. run server blockchain node 0

    ```make blockchain_server0```
2. run server blockchain node 1

    ```make blockchain_server1```
3. run server blockchain node 1

    ```make blockchain_server2```

4. run wallet server

    ```make wallet_server```
