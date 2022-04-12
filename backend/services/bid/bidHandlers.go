package bid

import (
	"auctionsPlatform/errors"
	"auctionsPlatform/models"
	"context"
)

func (s *service) handleAbsoluteBid(ctx context.Context, auctionID string, bid models.Bid) error {
	latestBid, err := s.bidRepository.GetLatestAuctionBids(ctx, auctionID)
	if err != nil {
		return err
	}

	if latestBid != nil && bid.Value <= latestBid[0].Value {
		return errors.ErrBidNotHigher
	}

	return nil
}
