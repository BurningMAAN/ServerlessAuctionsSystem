package bid

import (
	"auctionsPlatform/errors"
	"auctionsPlatform/models"
	"context"
	"fmt"
)

func (s *service) handleAbsoluteBid(ctx context.Context, auctionID string, bid models.Bid, userBalance float64, bidIncrement float64) error {
	latestBids, err := s.bidRepository.GetLatestAuctionBids(ctx, auctionID)
	if err != nil {
		return fmt.Errorf("get latest auction bids err: %s", err.Error())
	}

	var latestBid float64
	if len(latestBids) == 0 {
		latestBid = 0
	} else {
		if userBalance < latestBids[0].Value+bidIncrement && len(latestBids) > 0 {
			return errors.ErrUnsufficientCreditBalance
		}
		latestBid = latestBids[0].Value
	}

	if bid.Value <= latestBid {
		return errors.ErrBidNotHigher
	}

	return nil
}
