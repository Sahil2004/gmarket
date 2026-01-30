package dtos

type SymbolInfo struct {
	Symbol   string `json:"symbol" example:"RELIANCE"`
	Exchange string `json:"exchange" example:"NSE"`
}

type WatchlistDTO struct {
	Index   int          `json:"index" example:"1"`
	Symbols []SymbolInfo `json:"symbols"`
}

type SymbolDTO struct {
	Symbol   string `json:"symbol" example:"RELIANCE"`
	Exchange string `json:"exchange" example:"NSE"`
}
