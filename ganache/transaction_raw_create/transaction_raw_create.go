package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/2890a3a5dca74346a55ecd0671521cd1")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("e817df5c1ce27474e791eb556f73624ff6ca9d57318e0daf9652da6c0cbb5fc5")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xcfe2c40aa55Ce799b55697870920484DDF651377")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 获取 RLP 编码
	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatal("Failed to marshal transaction:", err)
	}

	// 转换为十六进制字符串
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Printf("Raw transaction hex: %s\n", rawTxHex)
}
