package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/9b28r4Hz8M_nF7j7DdOvCRsvX96Eg-0c")
	if err != nil {
		log.Fatal(err)
	}
	//您可以调用客户端的 HeaderByNumber 来返回有关一个区块的头信息。若您传入 nil，它将返回最新的区块头。

	header, err := client.HeaderByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744

	blockNumber := big.NewInt(7315417)
	//调用客户端的 BlockByNumber 方法来获得完整区块。您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等。
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744   区块的高度
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065

	//区块的哈希
	fmt.Println(block.Hash().Hex())        // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions())) // 144

	count, err := client.TransactionCount(context.Background(), block.Hash()) //交易次数
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144
}
