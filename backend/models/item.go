//go:build !test
// +build !test

package models

type Item struct {
	ID          string
	Description string
	Category    ItemCategory
	OwnerID     string
	PhotoURLs   []string
	AuctionID   string
	Name        string
}

type ItemCategory string

var (
	ItemCategoryCar        ItemCategory = "Car"
	ItemCategoryAppliances ItemCategory = "Appliances"
	ItemCategoryHome       ItemCategory = "Home"
)

type ItemUpdate struct {
	AuctionID   *string
	OwnerID     *string
	Category    *string
	Description *string
	Name        *string
}

type ItemSearchParams struct {
	OwnerID  string
	Category *string
}
