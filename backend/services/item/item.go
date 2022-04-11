package item

import (
	"auctionsPlatform/models"
	"context"
)

type itemRepository interface {
	CreateItem(ctx context.Context, item models.Item) (models.Item, error)
	GetItemByID(ctx context.Context, itemID string) (models.Item, error)
	AssignItem(ctx context.Context, auctionID, itemID string) error
	GetItemsByUserName(ctx context.Context, userName string) ([]models.Item, error)
}

type auctionRepository interface {
	GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error)
}
type service struct {
	itemRepository    itemRepository
	auctionRepository auctionRepository
}

func New(itemRepository itemRepository, auctionRepository auctionRepository) *service {
	return &service{
		itemRepository:    itemRepository,
		auctionRepository: auctionRepository,
	}
}

func (s *service) CreateItem(ctx context.Context, item models.Item) (models.Item, error) {
	return s.itemRepository.CreateItem(ctx, item)
}

func (s *service) GetItemByID(ctx context.Context, itemID string) (models.Item, error) {
	return s.itemRepository.GetItemByID(ctx, itemID)
}

func (s *service) AssignItem(ctx context.Context, auctionID, itemID string) error {
	item, err := s.itemRepository.GetItemByID(ctx, itemID)
	if err != nil {
		return err
	}

	auction, err := s.auctionRepository.GetAuctionByID(ctx, auctionID)
	if err != nil {
		return err
	}

	return s.itemRepository.AssignItem(ctx, auction.ID, item.ID)
}

func (s *service) GetItemsByUserName(ctx context.Context, userName string) ([]models.Item, error) {
	return s.itemRepository.GetItemsByUserName(ctx, userName)
}
