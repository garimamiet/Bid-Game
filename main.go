package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gitlab.com/garima.miet/bidgame/handlers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()

<<<<<<< HEAD
	r.HandleFunc("/createGame", handlers.CreateGameHandler)
	r.HandleFunc("/start", handlers.StartBidHandler)
	r.HandleFunc("/bid/{id:[0-9]+}", handlers.NewBidHandler)
	r.HandleFunc("/lastBid", handlers.LastBidHandler)
	r.HandleFunc("/expiry", handlers.ExpiryHandler)
	r.HandleFunc("/winner", handlers.WinnerHandler)
	r.HandleFunc("/withdrawCommission", handlers.CommissionHandler)
	r.HandleFunc("/reward", handlers.RewardsBidHandler)

=======
	r.HandleFunc("/bid/{id}", handlers.NewBidHandler)
	r.HandleFunc("/lastBid", handlers.LastBidHandler)
	r.HandleFunc("/reward", handlers.RewardsBidHandler)
	r.HandleFunc("/start", handlers.StartBidHandler)
	r.HandleFunc("/expiry", handlers.ExpiryHandler)
	r.HandleFunc("/withdrawCommission", handlers.CommissionHandler)
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
	log.Fatal(http.ListenAndServe(":3000", r))

}
