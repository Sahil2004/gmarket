export interface IWatchlistSymbol {
  exchange: string;
  last_close_price: number;
  ltp: number;
  symbol: string;
}

export interface IWatchlist {
  index: number;
  symbols: IWatchlistSymbol[];
}
