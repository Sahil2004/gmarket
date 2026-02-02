export interface IStock {
  symbol: string;
  name: string;
  exchange: string;
}

interface IDepth {
  orders: number;
  price: number;
  qty: number;
}
export interface IMarketDepth {
  asks: IDepth[];
  bids: IDepth[];
  exchange: string;
  symbol: string;
  ltp: number;
}
