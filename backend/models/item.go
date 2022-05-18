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
	ItemCategoryTransport   ItemCategory = "Transport"
	ItemCategoryElectronics ItemCategory = "Appliance"
	ItemCategoryHome        ItemCategory = "Furniture"
	ItemCategoryClothing    ItemCategory = "Clothing"
	ItemCategoryCarParts    ItemCategory = "CarPart"
	ItemCategoryPainting    ItemCategory = "Painting"
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

func (i ItemCategory) String() string {
	return string(i)
}
