//go:build !test
// +build !test

package models

import "time"

type Bid struct {
	ID        string
	Value     float64
	UserID    string
	AuctionID string
	Timestamp time.Time
}
