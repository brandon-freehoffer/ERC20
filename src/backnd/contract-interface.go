package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/brandon-freehoffer/ERC20/src/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/labstack/echo/v4"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	e := echo.New()
	tokenAddress := common.HexToAddress("0xb055e1b310c07c0825945C6e2e324304F0579D12")
	instance, err := api.NewApi(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0xA6FB0Be1a7fD6291414e465CA549e5bbbd477068")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("558a0233c8963a9a37ede845745290d9fda2c78cc2997a5cd758a1039a3bc843")
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

	// estimate gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice

	recipient := common.HexToAddress("0xabdd7608DcBb7AD30Dc0043AC382844F406f649b")

	amount := new(big.Int)
	amount.SetString("100", 3) // sets the value to 1000 tokens, in the token denomination
	app, err := instance.Approve(&bind.TransactOpts{}, recipient, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("approved: %s", app)

	tx, err := instance.Mint(auth, recipient, amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	fmt.Printf("name: %s\n", name)
	fmt.Printf("symbol: %s\n", symbol)
	fmt.Printf("decimals: %v\n", decimals)

	fmt.Printf("wei: %s\n", bal)

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	val2 := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", val2)

	e.Logger.Fatal(e.Start(":1331"))
}
