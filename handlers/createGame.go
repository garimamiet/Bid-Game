package handlers

import (
	"crypto/ecdsa"
	"encoding/json"

	"log"
	"math/big"
	"net/http"

	"context"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
	"gitlab.com/garima.miet/bidgame/biddingFactory"
)

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("https://nodeapi.test.energi.network/v1/jsonrpc")
	if err != nil {
		log.Fatal(err)
	}

	factoryInstance, err := biddingFactory.NewBiddingFactory(common.HexToAddress(os.Getenv("biddingFactory")), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("pvt_key"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(49797))
	if err != nil {
		log.Fatal(err)
	}
	privateKey = nil
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice

	auth.Nonce = big.NewInt(int64(nonce))

	auth.Value = big.NewInt(int64(0)) // in wei

	tx, err := factoryInstance.Create(auth)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(tx.Hash().Hex())

}
