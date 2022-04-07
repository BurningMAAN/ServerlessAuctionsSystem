//go:build !test
// +build !test

package models

import "time"

type Auction struct {
	ID          string
	IsFinished  bool
	Type        AuctionType
	BuyNowPrice float32
	StartDate   time.Time
	EndTime     time.Time
}

type AuctionType string

const (
	AuctionTypeAbsolute AuctionType = "AbsoluteAuction"
)
