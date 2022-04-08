//go:build !test
// +build !test

package models

import "time"

type Auction struct {
	ID           string
	IsFinished   bool
	Type         AuctionType
	BuyNowPrice  float32
	StartDate    time.Time
	EndTime      time.Time
	Participants []string /// ID's of Auction participants
	MinimalBid   float32
	CreatorID    string
}

type AuctionType string

const (
	AuctionTypeAbsolute AuctionType = "AbsoluteAuction"
)
