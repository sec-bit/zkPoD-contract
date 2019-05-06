package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/miguelmota/go-solidity-sha3"
)

func main() {

	/* Example Batch1Receipt */
	// uint256 sessionId;
	// address from;
	// bytes32 seed2;
	// bytes32 k_mkl_root;
	// uint64 count;
	// uint256 price;
	// uint256 expireAt;

	// ["0x00","0x7dB68eac733eB62C9f416a32F6B1d3b42179a53D","0x10","0x20",3,4,5]

	addr := common.HexToAddress("0x7dB68eac733eB62C9f416a32F6B1d3b42179a53D")

	receiptHash := solsha3.SoliditySHA3( // types
		[]string{"uint256", "address", "bytes32", "bytes32", "uint64", "uint256", "uint256"},

		// values
		[]interface{}{
			big.NewInt(0),
			addr,
			"0x10",
			"0x20",
			uint64(3),
			big.NewInt(4),
			big.NewInt(5),
		})

	log.Printf("receiptHash=0x%02x\n", receiptHash)

	privateKey, err := crypto.HexToECDSA("6ec15b813bc3ff3b61ed167ef42166cc66d7a04b5aa2decdde6ca9b0392a3306")
	if err != nil {
		log.Fatal(err)
	}

	// Buyer sign the receipt
	sig, err := crypto.Sign(receiptHash, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Buyer got signature, return it to seller along with the receipt
	log.Printf("sig=%v", hexutil.Encode(sig))

	// Seller split (r, s, v) from signature for contract
	r := fmt.Sprintf("%x", sig[:32])
	s := fmt.Sprintf("%x", sig[32:64])
	v := int(sig[64]) + 27 // add 27 for Ethereum contract
	log.Printf("r=0x%v, s=0x%v, v=0x%v", r, s, v)

	// Seller verify receipt signature
	sigPublicKeyECDSA, err := crypto.SigToPub(receiptHash, sig)
	if err != nil {
		log.Fatal(err)
	}

	// In case we do not know signer's publicKey
	rAddr := crypto.PubkeyToAddress(*sigPublicKeyECDSA)
	log.Printf("rAddr is equal to addr (%v)", rAddr == addr)
}
