package handlers

import (
	"encoding/json"
	"math/big"
	"os"

	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
	"gitlab.com/garima.miet/bidgame/bidding"
	"gitlab.com/garima.miet/bidgame/biddingFactory"
)

func WinnerHandler(w http.ResponseWriter, r *http.Request) {
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
	biddingContractAddress, _ := factoryInstance.GetActiveGames(nil, big.NewInt(int64(0)))

	instance, err := bidding.NewBidding(biddingContractAddress, client)

	if err != nil {
		log.Fatal(err)
	}
	winner, err := instance.CheckWinner(nil)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(winner)

}
