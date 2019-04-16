package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/miguelmota/go-solidity-sha3"
)

// function keccak256EncodeTest(uint64 _size, uint64 _s, uint64 _n, uint256 _sigma_mkl_root, uint256 _vrf_meta_digest, uint256 _blt_type)

// bytes32 _bltKey = keccak256(abi.encodePacked(_size, _s, _n, _sigma_mkl_root));

func main() {
	_size := uint64(1)
	_s := uint64(2)
	_n := uint64(3)
	_sigma_mkl_root := big.NewInt(4)

	// solidity hash result
	solHash := "8812e4c4163c9d5da0424869eac0f59b3ee986bbcc4b294145e0b25c0f3648b8"

	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint64", "uint64", "uint64", "uint256"},

		// values
		[]interface{}{
			_size,
			_s,
			_n,
			_sigma_mkl_root,
		},
	)
	goHash := hex.EncodeToString(hash)
	fmt.Printf("solHash=0x%v\n", solHash)
	fmt.Printf("goHash=0x%v\n", goHash)
	fmt.Printf("equal=%v\n", solHash == goHash)
}
