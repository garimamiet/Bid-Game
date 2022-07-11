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

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gitlab.com/garima.miet/bidgame/bidding"
<<<<<<< HEAD
	"gitlab.com/garima.miet/bidgame/biddingFactory"
=======
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
)

func NewBidHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("https://nodeapi.test.energi.network/v1/jsonrpc")
	if err != nil {
		log.Fatal(err)
	}

<<<<<<< HEAD
	factoryInstance, err := biddingFactory.NewBiddingFactory(common.HexToAddress(os.Getenv("biddingFactory")), client)
	if err != nil {
		log.Fatal(err)
	}
	biddingContractAddress, _ := factoryInstance.GetActiveGames(nil, big.NewInt(int64(0)))

=======
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
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
<<<<<<< HEAD

=======
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(49797))
	if err != nil {
		log.Fatal(err)
	}
<<<<<<< HEAD

	privateKey = nil

=======
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
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

	params := mux.Vars(r)

	value, _ := new(big.Int).SetString(params["id"], 10)

	auth.Value = value // in wei

<<<<<<< HEAD
	address := biddingContractAddress

=======
	auth.GasLimit = uint64(300000) // in units

	address := common.HexToAddress("0x5F3D6b53978057c655574a2d84EEFcC49a2C1CE8")
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
	instance, err := bidding.NewBidding(address, client)

	if err != nil {
		log.Fatal(err)
	}
	tx, err := instance.Bid(auth)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(tx.Hash().Hex())

}
