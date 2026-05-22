export interface IAlgoConfig {
  symbol: string;
  exchange: string;
  enabled: boolean;
  rsi_enabled: boolean;
  rsi_period: number;
  rsi_overbought: number;
  rsi_oversold: number;
  ma_enabled: boolean;
  ma_fast_period: number;
  ma_slow_period: number;
}

export interface IIndicatorPoint {
  time: number;
  value: number;
}

export interface IAlgoIndicators {
  symbol: string;
  exchange: string;
  range: string;
  config: IAlgoConfig;
  rsi: IIndicatorPoint[];
  ma_fast: IIndicatorPoint[];
  ma_slow: IIndicatorPoint[];
  current_rsi: number;
  current_ma_fast: number;
  current_ma_slow: number;
  rsi_signal: AlgoSignal;
  ma_signal: AlgoSignal;
  combined_signal: AlgoSignal;
}

export type AlgoSignal = 'buy' | 'sell' | 'hold' | 'bullish' | 'bearish' | 'disabled';
