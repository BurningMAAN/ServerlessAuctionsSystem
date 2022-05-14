package auction

import (
	errs "auctionsPlatform/errors"
	"auctionsPlatform/models"
	"auctionsPlatform/repositories/auction"
	"context"
	"log"
	"time"
)

type auctionRepository interface {
	CreateAuction(ctx context.Context, auction models.Auction, category models.ItemCategory) (models.Auction, error)
	GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error)
	GetAllAuctions(ctx context.Context, optFns ...func(*auction.OptionalGetParameters)) ([]models.Auction, error)
	SearchAuctions(ctx context.Context, searchParams models.AuctionSearchParams) ([]models.Auction, error)
	UpdateAuction(ctx context.Context, auctionID string, update models.AuctionUpdate) error
	DeleteAuction(ctx context.Context, auctionID string) error
}

type itemRepository interface {
	AssignItem(ctx context.Context, auctionID, itemID string) error
	GetItemByID(ctx context.Context, itemID string) (models.Item, error)
	UpdateItem(ctx context.Context, itemID string, update models.ItemUpdate) error
}

type eventRepository interface {
	CreateEventRule(ctx context.Context, auctionID string, startDate time.Time) error
	UpdateEventRule(ctx context.Context, auctionID string, newDate time.Time) error
	DeleteEventRule(ctx context.Context, auctionID string) error
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

	auction.PhotoURL = item.PhotoURLs[0]
	auction, err = s.auctionRepository.CreateAuction(ctx, auction, item.Category)
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

func (s *service) SearchAuctions(ctx context.Context, searchParams models.AuctionSearchParams) ([]models.AuctionListView, error) {
	view := []models.AuctionListView{}
	auctions, err := s.auctionRepository.SearchAuctions(ctx, searchParams)
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

func (s *service) UpdateAuction(ctx context.Context, auctionID string, update models.AuctionUpdate) error {
	auction, err := s.auctionRepository.GetAuctionByID(ctx, auctionID)
	if err != nil {
		return err
	}

	if auction.Stage == "STAGE_AUCTION_ONGOING" || auction.Stage == "STAGE_AUCTION_FINISHED" {
		return errs.ErrAuctionCannotBeUpdate
	}

	if update.StartDate != nil && update.StartDate.Before(time.Now()) {
		return errs.ErrAuctionInvalidDateUpdate
	}

	err = s.auctionRepository.UpdateAuction(ctx, auction.ID, update)
	if err != nil {
		return err
	}

	return s.eventRepository.UpdateEventRule(ctx, auction.ID, *update.StartDate)
}

func (s *service) DeleteAuction(ctx context.Context, auctionID string) error {
	auction, err := s.auctionRepository.GetAuctionByID(ctx, auctionID)
	if err != nil {
		return err
	}

	if auction.Stage == "STAGE_AUCTION_ONGOING" || auction.Stage == "STAGE_AUCTION_FINISHED" {
		return errs.ErrAuctionCannotBeUpdate
	}

	item, err := s.itemRepository.GetItemByID(ctx, auction.ItemID)
	if err != nil {
		return err
	}

	err = s.auctionRepository.DeleteAuction(ctx, auctionID)
	if err != nil {
		return err
	}

	err = s.itemRepository.UpdateItem(ctx, item.ID, models.ItemUpdate{
		AuctionID: nil,
	})
	if err != nil {
		return err
	}

	return s.eventRepository.DeleteEventRule(ctx, auction.ID)
}
