//go:build !test
// +build !test

package models

import "time"

type Auction struct {
	ID           string
	Type         AuctionType
	BuyoutPrice  *float64
	StartDate    time.Time
	Participants []string /// ID's of Auction participants
	BidIncrement float64
	CreatorID    string
	ItemID       string
	EndDate      time.Time
	Stage        string
	// Category string Prideti kategorija (nes rodome korteleje)
}

type AuctionType string

const (
	AuctionTypeAbsolute AuctionType = "AbsoluteAuction"
)

type AuctionListView struct {
	Auction Auction
	Item    Item
}

type AuctionEvent struct {
	AuctionID string    `json:"id"`
	Stage     string    `json:"stage"`
	EndDate   time.Time `json:"endDate"`
}
