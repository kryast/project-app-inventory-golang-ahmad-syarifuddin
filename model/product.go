package model

type Item struct {
	ID         int
	ItemCode   string
	Name       string
	CategoryId int
	LocationId int
	Price      int
	Stock      int
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Location struct {
	ID           int    `json:"id"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Province     string `json:"province"`
	ItemPosition string `json:"item_position"`
}
