package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {

	value := flag.String("v", "", "uint256")

	flag.Parse()

	bn := new(big.Int)
	bn.SetString(*value, 10)
	result := abi.U256(bn)
	fmt.Printf("result=%v\n", result)
	fmt.Printf("result=0x%02x\n", result)

	// result1 := bn.Bytes()
	// fmt.Printf("result1=%v\n", result1)
	// fmt.Printf("result1=0x%02x\n", result1)

	result2 := reverseByteSlice(result)
	fmt.Printf("result2=%v\n", result2)
	fmt.Printf("result2=0x%02x\n", result2)
}

func reverseByteSlice(a []byte) []byte {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
