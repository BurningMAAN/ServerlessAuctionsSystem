//go:build !test
// +build !test

package models

import "time"

type Auction struct {
	ID           string
	IsFinished   bool
	Type         AuctionType
	BuyoutPrice  *float64
	StartDate    time.Time
	Participants []string /// ID's of Auction participants
	BidIncrement float64
	CreatorID    string
}

type AuctionType string

const (
	AuctionTypeAbsolute AuctionType = "AbsoluteAuction"
)
