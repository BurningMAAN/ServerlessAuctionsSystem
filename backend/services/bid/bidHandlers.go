package bid

import (
	"auctionsPlatform/errors"
	"auctionsPlatform/models"
	"context"
	"fmt"
)

func (s *service) handleAbsoluteBid(ctx context.Context, auctionID string, bid models.Bid) error {
	latestBid, err := s.bidRepository.GetLatestAuctionBids(ctx, auctionID)
	if err != nil {
		return fmt.Errorf("get latest auction bids err: %s", err.Error())
	}

	if latestBid != nil && bid.Value <= latestBid[0].Value {
		return errors.ErrBidNotHigher
	}

	return nil
}
