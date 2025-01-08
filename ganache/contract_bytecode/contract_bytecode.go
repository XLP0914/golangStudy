package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/2890a3a5dca74346a55ecd0671521cd1")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x4db51e8b6f65d63b5983edf3bb5833253c2485fa")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
}
