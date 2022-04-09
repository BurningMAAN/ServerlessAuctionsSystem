package item

import (
	"auctionsPlatform/models"
	"context"
)

type itemRepository interface {
	CreateItem(ctx context.Context, item models.Item) (models.Item, error)
	GetItemByID(ctx context.Context, itemID string) (models.Item, error)
}

type service struct {
	itemRepository itemRepository
}

func New(itemRepository itemRepository) *service {
	return &service{
		itemRepository: itemRepository,
	}
}

func (s *service) CreateItem(ctx context.Context, item models.Item) (models.Item, error) {
	return s.itemRepository.CreateItem(ctx, item)
}

func (s *service) GetItemByID(ctx context.Context, itemID string) (models.Item, error) {
	return s.itemRepository.GetItemByID(ctx, itemID)
}
