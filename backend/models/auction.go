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
	Category     string
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

type AuctionSearchParams struct {
	Category    *string
	AuctionType *string
	OwnerID     *string
	Status      *string
}

type AuctionUpdate struct {
	BuyoutPrice  *float64
	StartDate    *time.Time
	BidIncrement *float64
	Type         *string
}
