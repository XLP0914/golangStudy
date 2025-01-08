package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//通过合约地址查询钱包信息
	address := common.HexToAddress("0xcfe2c40aa55Ce799b55697870920484DDF651377")

	// 打印地址的十六进制格式
	fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F

	// 使用 Keccak256 哈希函数计算地址的哈希值
	addressHash := crypto.Keccak256Hash(address.Bytes())
	fmt.Println(addressHash.Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f

	// 打印地址的字节切片
	fmt.Println(address.Bytes()) // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
}
