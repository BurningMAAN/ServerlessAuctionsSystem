//go:build !test
// +build !test

package models

type Item struct {
	ID          string
	Description string
	Category    ItemCategory
	OwnerID     string
	PhotoURLs   []string
}

type ItemCategory string

var (
	ItemCategoryCar        ItemCategory = "Car"
	ItemCategoryAppliances ItemCategory = "Appliances"
	ItemCategoryHome       ItemCategory = "Home"
)
