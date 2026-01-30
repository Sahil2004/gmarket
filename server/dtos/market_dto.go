package dtos

type LevelDTO struct {
	Price  float64 `json:"price"`
	Qty    int     `json:"qty"`
	Orders int     `json:"orders"`
}

type MarketDepthDTO struct {
	Symbol   string     `json:"symbol" example:"RELIANCE"`
	Exchange string     `json:"exchange" example:"NSE"`
	LTP      float64    `json:"ltp" example:"2450.75"`
	Bids     []LevelDTO `json:"bids"`
	Asks     []LevelDTO `json:"asks"`
}
