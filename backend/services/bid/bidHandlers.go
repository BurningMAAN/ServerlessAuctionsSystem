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

// Aukcionu statusu keitima gali implementinti su scheduled eventBridge eventu ar chronJob'u kas keleta sekundziu ir tikrinti aukcionu statusus
// galima pabandyti gal publishinti kazkoki tai job'a kuris b

/*
	Solution 1:
		1. Kas valanda yra prasukama lambda, kuri tikrina, ar yra ateinancia valanda ivykstanciu aukcionu
			* Jei yra:
				* Sukuriamas DB irasas, turintis duomenis apie aukciono pabaiga (timestamp ar dar kazka)
				*
*/
