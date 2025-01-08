package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {

	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/2890a3a5dca74346a55ecd0671521cd1")
	if err != nil {
		log.Fatal(err)
	}

	rawTx := "f8700785079091e7b182520894cfe2c40aa55ce799b55697870920484ddf651377880de0b6b3a7640000808401546d71a054d3202e35208b4d360e236b5a2ef002643c6e00f8793c6f49ba0a7ca3a3ed2ea05aa09f6449a92b371b2e07cb214079385ed465e8d537f3148b4daf2bf19cafb7"

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
