package auction

import (
	"auctionsPlatform/models"
	"auctionsPlatform/repositories/auction"
	"context"
	"log"
	"time"
)

type auctionRepository interface {
	CreateAuction(ctx context.Context, auction models.Auction) (models.Auction, error)
	GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error)
	GetAllAuctions(ctx context.Context, optFns ...func(*auction.OptionalGetParameters)) ([]models.Auction, error)
}

type itemRepository interface {
	AssignItem(ctx context.Context, auctionID, itemID string) error
	GetItemByID(ctx context.Context, itemID string) (models.Item, error)
}

type eventRepository interface {
	CreateEventRule(ctx context.Context, auctionID string, startDate time.Time) error
}

type service struct {
	itemRepository    itemRepository
	auctionRepository auctionRepository
	eventRepository   eventRepository
}

func New(auctionRepository auctionRepository, itemRepository itemRepository, eventRepository eventRepository) *service {
	return &service{
		auctionRepository: auctionRepository,
		itemRepository:    itemRepository,
		eventRepository:   eventRepository,
	}
}

func (s *service) CreateAuction(ctx context.Context, auction models.Auction, itemID string) (models.Auction, error) {
	item, err := s.itemRepository.GetItemByID(ctx, itemID)
	if err != nil {
		return models.Auction{}, err
	}

	auction, err = s.auctionRepository.CreateAuction(ctx, auction)
	if err != nil {
		return models.Auction{}, err
	}

	err = s.itemRepository.AssignItem(ctx, auction.ID, item.ID)
	if err != nil {
		return models.Auction{}, err
	}

	err = s.eventRepository.CreateEventRule(ctx, auction.ID, auction.StartDate)
	if err != nil {
		log.Printf("failed to publish event for auctionID: %s", auction.ID)
		log.Printf("got error: %s", err.Error())
		return auction, nil
	}

	return auction, err
}

func (s *service) GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error) {
	return s.auctionRepository.GetAuctionByID(ctx, auctionID)
}

func (s *service) GetAuctions(ctx context.Context) ([]models.AuctionListView, error) {
	view := []models.AuctionListView{}

	auctions, err := s.auctionRepository.GetAllAuctions(ctx)
	if err != nil {
		return []models.AuctionListView{}, err
	}

	for _, auction := range auctions {
		item, err := s.itemRepository.GetItemByID(ctx, auction.ItemID)
		if err != nil {
			return []models.AuctionListView{}, err
		}
		view = append(view, models.AuctionListView{
			Auction: auction,
			Item:    item,
		})
	}

	return view, err
}
