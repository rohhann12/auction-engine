package auction

import "time"

type AuctionSymbols string

const (
	Waiting AuctionSymbols = "WAITING"
	Live    AuctionSymbols = "LIVE"
	Ended   AuctionSymbols = "ENDED"
)

type Auction struct {
	Id         string
	Item       string
	CurrentBid int
	EndsAt     int64
	Status     AuctionSymbols
	Bids       []Bid
}

type Bid struct {
	UserId    string
	Amount    int
	Timestamp time.Duration
	Accepted  bool
	Topic     string
}

type BidRequest struct {
	// Client  *websocket.Conn
	Topic  string
	UserId string
	Amount int
}
