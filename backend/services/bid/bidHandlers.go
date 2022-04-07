package bid

import (
	"auctionsPlatform/errors"
	"auctionsPlatform/models"
	"context"
)

func (s *service) handleAbsoluteBid(ctx context.Context, auctionID string, bid models.Bid) error {
	latestBid, err := s.bidRepository.GetLatestAuctionBid(ctx, auctionID)
	if err != nil {
		return err
	}

	if latestBid != nil && bid.Value <= latestBid.Value {
		return errors.ErrBidNotHigher
	}

	return nil
}
