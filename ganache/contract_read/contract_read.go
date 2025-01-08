package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "./contracts" // for demo
)

func main() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/2890a3a5dca74346a55ecd0671521cd1")

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("e817df5c1ce27474e791eb556f73624ff6ca9d57318e0daf9652da6c0cbb5fc5")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
