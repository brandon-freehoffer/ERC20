package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/brandon-freehoffer/ERC20/src/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/sha3"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	tokenAddress := common.HexToAddress("0x5274222f6856F76C14d8046079DbA79466654c30")
	instance, err := api.NewApi(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	e.GET("/GetTokenInfo", func(c echo.Context) error {
		fmt.Println(c.Request())
		name, err := instance.Name(&bind.CallOpts{})
		if err != nil {
			fmt.Printf("Error")
			return err
		}
		symbol, err := instance.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		decimals, err := instance.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		t := &TokenInfo{
			Name:      name,
			Symbol:    symbol,
			Precision: decimals,
		}

		return c.JSON(http.StatusOK, t) // get name

	})
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
	privateKey, err := crypto.HexToECDSA("0a0895ec306ed938ac3eb29559bcc9d1a8df59afd909c60bb57f286547873392")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress("0x0C4fCeE187bB4889Ff27cDf9ba0D41665300550A")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Printf("\nMethod ID: %s\n", hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Printf("\nTo address: %s\n", hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Printf("\nToken amount: %s", hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	e.GET("/Sign", func(c echo.Context) error {

		add := c.QueryParam("adress")
		fmt.Println("add: " + add)
		toAddress := common.HexToAddress(add)
		transferFnSignature := []byte("transfer(address,uint256)")
		hash := sha3.NewLegacyKeccak256()
		hash.Write(transferFnSignature)
		methodID := hash.Sum(nil)[:4]
		fmt.Printf("\nMethod ID: %s\n", hexutil.Encode(methodID))

		paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
		fmt.Printf("\nTo address: %s\n", hexutil.Encode(paddedAddress))

		amount := new(big.Int)
		amount.SetString("1000000000000000000000", 10) // 1000 tokens
		paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		value := big.NewInt(0) // in wei (0 eth)
		gasl := uint64(5000000)
		fmt.Printf("\nToken amount: %s", hexutil.Encode(paddedAmount))
		tokenAddress := common.HexToAddress("0x5274222f6856F76C14d8046079DbA79466654c30")
		var data []byte
		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedAmount...)

		tx := types.NewTransaction(nonce, tokenAddress, value, gasl, gasPrice, data)
		signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
		if err != nil {
			log.Fatal(err)
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, tx)
	})

	e.GET("/Transfer", func(c echo.Context) error {
		trans := c.QueryParam("tx")
		raw, _ := hexutil.Decode(trans)
		var tx *types.Transaction
		fmt.Println(rlp.DecodeBytes(raw, &tx))

		return c.JSON(http.StatusOK, trans)
	})
	fmt.Printf("\nname: %s\n", name)
	fmt.Printf("\nsymbol: %s\n", symbol)
	fmt.Printf("\ndecimals: %v\n", decimals)

	e.Logger.Fatal(e.Start(":1348"))
}

type TokenInfo struct {
	Name      string
	Symbol    string
	Precision uint8
}
