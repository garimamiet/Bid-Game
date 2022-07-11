package handlers

import (
	"encoding/json"
<<<<<<< HEAD
	"math/big"
	"os"
=======
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878

	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

<<<<<<< HEAD
	"github.com/joho/godotenv"
	"gitlab.com/garima.miet/bidgame/bidding"
	"gitlab.com/garima.miet/bidgame/biddingFactory"
=======
	//"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gitlab.com/garima.miet/bidgame/bidding"
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
)

func ExpiryHandler(w http.ResponseWriter, r *http.Request) {
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

	address := biddingContractAddress
=======
	address := common.HexToAddress("0x5F3D6b53978057c655574a2d84EEFcC49a2C1CE8")
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
	instance, err := bidding.NewBidding(address, client)

	if err != nil {
		log.Fatal(err)
	}
	expiry, err := instance.Expiry(nil)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(expiry)

}
