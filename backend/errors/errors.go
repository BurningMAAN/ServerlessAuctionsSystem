//go:build !test
// +build !test

package errors

import "errors"

var (
	ErrBidNotHigher             = errors.New("provided bid value is lower than the existing one")
	ErrAuctionAlreadyFinished   = errors.New("auction is already finished")
	ErrAuctionUserBidUserMatch  = errors.New("bid and auction user matches")
	ErrAuctionCannotBeUpdate    = errors.New("auction is not in updatable state")
	ErrAuctionInvalidDateUpdate = errors.New("given auction date is not in future")
)
