package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const (
	envPath         = ".env,.env.local"
	contractAddress = "0xc6e865c213c89ca42a622c5572d19f00d84d7a16"
)

func main() {
	if err := godotenv.Overload(strings.Split(envPath, ",")...); err != nil {
		fmt.Println("Load env error", err.Error())
	}

	amount, err := strconv.Atoi(os.Getenv("AMOUNT_CALL"))
	if err != nil {
		log.Fatal(err)
	}

	currentNonce, chainId := GetCurrentNonceAndChainId()
	for i := 0; i < amount; i++ {
		log.Println(currentNonce)
		err := CallContract(currentNonce, chainId)
		if err != nil {
			log.Println(err)
			i = i - 1
			continue
		}
		currentNonce = currentNonce + 1
	}
}

func GetCurrentNonceAndChainId() (uint64, *big.Int) {
	client, err := ethclient.Dial(os.Getenv("POLYGON_RPC"))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
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

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return nonce, chainID
}

func CallContract(nonce uint64, chainId *big.Int) error {
	client, err := ethclient.Dial(os.Getenv("POLYGON_RPC"))
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(30000) // in units
	auth.GasPrice = gasPrice

	CFXContract, err := NewCFXsContract(common.HexToAddress(contractAddress), client)
	if err != nil {
		return err
	}

	tx, err := CFXContract.CreateCFXs(auth)
	if err != nil {
		return err
	}

	log.Println("Tx Hash", tx.Hash().String())
	return nil
}
