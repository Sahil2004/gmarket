export interface IStock {
  symbol: string;
  name: string;
  exchange: string;
}

export interface IChartApiResponse {
  chart: IChart;
}

export interface IChart {
  result: IChartResult[] | null;
  error: unknown | null;
}

// Result
export interface IChartResult {
  meta: IChartMeta;
  timestamp: number[];
  indicators: IIndicators;
}

// Meta
export interface IChartMeta {
  currency: string;
  symbol: string;
  exchangeName: string;
  fullExchangeName: string;
  instrumentType: string;

  firstTradeDate: number;
  regularMarketTime: number;

  hasPrePostMarketData: boolean;
  gmtoffset: number;
  timezone: string;
  exchangeTimezoneName: string;

  regularMarketPrice: number;
  fiftyTwoWeekHigh: number;
  fiftyTwoWeekLow: number;
  regularMarketDayHigh: number;
  regularMarketDayLow: number;
  regularMarketVolume: number;

  longName: string;
  shortName: string;

  chartPreviousClose: number;
  previousClose: number;

  scale: number;
  priceHint: number;

  currentTradingPeriod: ICurrentTradingPeriod;
  tradingPeriods: ITradingPeriod[][];

  dataGranularity: string; // e.g. "1m"
  range: string; // e.g. "1d"

  validRanges: string[];
}

// Trading periods
export interface ICurrentTradingPeriod {
  pre: ITradingSession;
  regular: ITradingSession;
  post: ITradingSession;
}

export interface ITradingSession {
  timezone: string;
  start: number;
  end: number;
  gmtoffset: number;
}

export interface ITradingPeriod {
  timezone: string;
  start: number;
  end: number;
  gmtoffset: number;
}

// Indicators
export interface IIndicators {
  quote: IQuote[];
}

export interface IQuote {
  open: Array<number | null>;
  high: Array<number | null>;
  low: Array<number | null>;
  close: Array<number | null>;
  volume: Array<number | null>;
}
