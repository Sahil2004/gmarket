export interface IWatchlistSymbol {
  exchange: string;
  symbol: string;
}

export interface IWatchlist {
  index: number;
  symbols: IWatchlistSymbol[];
}

export interface IWatchlistSymbolInfo {
  symbol: string;
  exchange: string;
  ltp: number;
  last_close_price: number;
}
