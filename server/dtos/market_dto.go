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

type SymbolListDTO struct {
	Symbols []SymbolDTO `json:"symbols" example:"[{\"symbol\":\"RELIANCE\",\"exchange\":\"NSE\"},{\"symbol\":\"TCS\",\"exchange\":\"NSE\"}]"`
}

type SymbolWithStatusDTO struct {
	Symbol         string `json:"symbol" example:"RELIANCE"`
	Exchange       string `json:"exchange" example:"NSE"`
	LTP            string `json:"ltp" example:"2450.75"`
	LastClosePrice string `json:"last_close_price" example:"2430.50"`
}
type SymbolStatusDTO struct {
	Symbols []SymbolWithStatusDTO `json:"symbols"`
}

type CandleDTO struct {
	Time  int64   `json:"time"`
	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
	Close float64 `json:"close"`
}

type CandlesDTO struct {
	Symbol   string      `json:"symbol"`
	Exchange string      `json:"exchange"`
	Range    string      `json:"range"`
	Candles  []CandleDTO `json:"candles"`
}
