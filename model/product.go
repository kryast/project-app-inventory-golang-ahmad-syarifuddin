package model

type Item struct {
	ID         uint16
	ItemCode   string
	Name       string
	CategoryId int
	LocationId int
	Price      int
	Stock      int
}

type Category struct {
	ID   uint16
	Name string
}

type Location struct {
	ID           uint16
	Address      string
	City         string
	Province     string
	ItemLocation string
}
