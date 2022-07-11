# Bidding Game:

## Factory contract address:
0x9365b77d6071910f9195d6606df1cf0a125bc19f

## Environment variable to be set:
* pvt_key = YOUR PRIVATE KEY
* biddingFactory = 0x9365b77d6071910f9195d6606df1cf0a125bc19f

## End points:
* "localhost:3000/ceateGame"
: to deploy a new bidding contract
* "localhost:3000/start": to start the bidding
* "localhost:3000/bid/{id:[0-9]+}": to bid
* "localhost:3000/lastBid": to check the last valid bid
* "localhost:3000/expiry": to check the game end time
* "localhost:3000/winner": to check the winner 
* "localhost:3000/withdrawCommission": to withdraw the commissions earned
* "localhost:3000/reward": to claim the rewards
