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
	ItemCategoryTransport   ItemCategory = "Transportas"
	ItemCategoryElectronics ItemCategory = "Elektronika"
	ItemCategoryHome        ItemCategory = "Baldas"
	ItemCategoryDrabužiai   ItemCategory = "Drabužiai"
	ItemCategoryDetalės     ItemCategory = "Detalė"
	ItemCategoryPaveikslas  ItemCategory = "Paveikslas"
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
