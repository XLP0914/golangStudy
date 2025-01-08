package main

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//用于生成签名的组件是：签名者私钥，以及将要签名的数据的哈希。
	privateKey, err := crypto.HexToECDSA("e817df5c1ce27474e791eb556f73624ff6ca9d57318e0daf9652da6c0cbb5fc5")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

	fmt.Println(publicKey)
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data) //使用 Keccak-256 生成哈希
	fmt.Println("hash", hash.Hex())    // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	//现在假设我们有字节格式的签名，我们可以从 go-ethereum crypto 包调用 Ecrecover（椭圆曲线签名恢复）来检索签名者的公钥。
	//此函数采用字节格式的哈希和签名。
	//我们使用私钥签名哈希，得到签名。
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("signature", hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	//SigToPub 是以太坊开发中常用的一个方法，主要用于从签名中恢复出公钥（Public Key）。
	//它是 Go Ethereum (go-ethereum) 库的一部分，属于 crypto 包。

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	//FromECDSAPub 是 Go Ethereum (go-ethereum) 库中一个方法，
	//主要用于将椭圆曲线的公钥（*ecdsa.PublicKey）转换为字节切片（[]byte）形式。
	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	//为方便起见，go-ethereum/crypto 包提供了 VerifySignature 函数，该函数接收原始数据的签名，哈希值和字节格式的公钥。
	//它返回一个布尔值，如果公钥与签名的签名者匹配，则为 true。
	//一个重要的问题是我们必须首先删除 signture 的最后一个字节，因为它是 ECDSA 恢复 ID，不能包含它。
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}
