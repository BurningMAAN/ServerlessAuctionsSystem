package models

type EntityType string

var (
	AuctionEntityType EntityType = "Auction"
	ItemEntityType    EntityType = "Item"
	UserEntityType    EntityType = "User"
	BidEntityType     EntityType = "Bid"
)
