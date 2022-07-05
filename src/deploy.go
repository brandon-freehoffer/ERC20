package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/brandon-freehoffer/ERC20/src/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println(err)
		panic(err)

	}

	privateKey, err := crypto.HexToECDSA("70e8a81413b2b7082eea34ffd227f42bb3e2dadc4ce90a8e82f06d716d103fc9")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {

		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = big.NewInt(875625000)

	address, tx, instance, err := api.DeployApi(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println("Contract address: " + address.Hex())

	_, _ = instance, tx
}
