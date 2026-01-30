export interface IWatchlistSymbol {
  exchange: string;
  symbol: string;
}

export interface IWatchlist {
  index: number;
  symbols: IWatchlistSymbol[];
}
