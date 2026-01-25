package dtos

type SymbolInfo struct {
	Symbol         string  `json:"symbol" example:"RELIANCE"`
	Exchange       string  `json:"exchange" example:"NSE"`
	LTP            float64 `json:"ltp" example:"2450.75"`
	LastClosePrice float64 `json:"last_close_price" example:"2430.50"`
}

type WatchlistDTO struct {
	Index   int          `json:"index" example:"1"`
	Symbols []SymbolInfo `json:"symbols"`
}

type SymbolDTO struct {
	Symbol   string `json:"symbol" example:"RELIANCE"`
	Exchange string `json:"exchange" example:"NSE"`
}
