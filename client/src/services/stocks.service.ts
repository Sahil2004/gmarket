import { inject, Injectable } from '@angular/core';
import { ChartRange, ICandles, IMarketDepth, IStock } from '../types/stocks.types';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom, map, shareReplay, type Observable } from 'rxjs';
import { IWatchlistSymbol, IWatchlistSymbolInfo } from '../types';
import { PollingService } from './polling.service';

@Injectable({
  providedIn: 'root',
})
export class StocksService {
  private http = inject(HttpClient);
  private polling = inject(PollingService);
  private stocks$: Observable<IStock[]> = this.http
    .get<IStock[]>('/market/symbols')
    .pipe(shareReplay(1));
  private stockDataCache = new Map<string, IWatchlistSymbolInfo>();
  private fetchedData: { symbols: IWatchlistSymbolInfo[] } = { symbols: [] };
  private symbolsToFetchSet: IWatchlistSymbol[] = [];
  private depthStock: IWatchlistSymbol | null = null;
  private fetchedDepthData: IMarketDepth | null = null;
  private chartStock: { symbol: string; exchange: string; range: ChartRange } | null = null;
  private fetchedCandles: ICandles | null = null;

  constructor() {
    this.polling.register('market', () => this.pollMarket());
  }

  setWatchlistPolling(active: boolean): void {
    if (active) {
      this.polling.start('market');
    } else if (!this.depthStock && !this.chartStock) {
      this.polling.stop('market');
    }
  }

  setDepthActive(symbol: string | null, exchange: string | null): void {
    if (symbol && exchange) {
      this.depthStock = { symbol, exchange };
      this.polling.start('market');
    } else {
      this.depthStock = null;
      this.fetchedDepthData = null;
      if (!this.chartStock && this.symbolsToFetchSet.length === 0) {
        this.polling.stop('market');
      }
    }
  }

  setChartActive(symbol: string | null, exchange: string | null, range: ChartRange | null): void {
    if (symbol && exchange && range) {
      this.chartStock = { symbol, exchange, range };
      this.polling.start('market');
    } else {
      this.chartStock = null;
      this.fetchedCandles = null;
      if (!this.depthStock && this.symbolsToFetchSet.length === 0) {
        this.polling.stop('market');
      }
    }
  }

  clearMarketPolling(): void {
    this.symbolsToFetchSet = [];
    this.depthStock = null;
    this.chartStock = null;
    this.fetchedDepthData = null;
    this.fetchedCandles = null;
    this.polling.stop('market');
  }

  private async pollMarket(): Promise<void> {
    if (this.symbolsToFetchSet.length > 0) {
      this.fetchedData = await firstValueFrom(
        this.http.post<{ symbols: IWatchlistSymbolInfo[] }>('/market/symbols/status', {
          symbols: this.symbolsToFetchSet,
        }),
      );
    }
    if (this.depthStock) {
      this.fetchedDepthData = await firstValueFrom(
        this.http.get<IMarketDepth>('/market/depth', {
          params: {
            symbol: this.depthStock.symbol,
            exchange: this.depthStock.exchange,
          },
        }),
      );
    }
    if (this.chartStock) {
      this.fetchedCandles = await firstValueFrom(
        this.http.get<ICandles>('/market/candles', {
          params: {
            symbol: this.chartStock.symbol,
            exchange: this.chartStock.exchange,
            range: this.chartStock.range,
          },
        }),
      );
    }
  }

  getAllStocks(): Observable<IStock[]> {
    return this.stocks$;
  }

  isAStockSymbol(symbol: string): Observable<boolean> {
    return this.getAllStocks().pipe(
      shareReplay(1),
      map((stocks) => stocks.some((stock) => stock.symbol === symbol)),
    );
  }

  private cacheKey(symbol: IWatchlistSymbol): string {
    return `${symbol.exchange}:${symbol.symbol}`;
  }

  async getStocksData(symbols: IWatchlistSymbol[]): Promise<IWatchlistSymbolInfo[]> {
    for (const item of this.fetchedData.symbols) {
      this.stockDataCache.set(this.cacheKey(item), item);
    }

    const missing = symbols.filter((s) => !this.stockDataCache.has(this.cacheKey(s)));
    if (missing.length > 0) {
      const fetchedData = await firstValueFrom(
        this.http.post<{ symbols: IWatchlistSymbolInfo[] }>('/market/symbols/status', {
          symbols: missing,
        }),
      );
      for (const item of fetchedData.symbols) {
        this.stockDataCache.set(this.cacheKey(item), item);
      }
    }
    this.symbolsToFetchSet = symbols;

    return symbols
      .map((s) => this.stockDataCache.get(this.cacheKey(s)))
      .filter((item): item is IWatchlistSymbolInfo => item !== undefined);
  }

  async getDepthData(symbol: string, exchange: string): Promise<IMarketDepth> {
    this.setDepthActive(symbol, exchange);
    if (!this.fetchedDepthData) {
      await this.pollMarket();
    }
    return this.fetchedDepthData as IMarketDepth;
  }

  peekDepthData(): IMarketDepth | null {
    return this.fetchedDepthData;
  }

  async getCandles(symbol: string, exchange: string, range: ChartRange): Promise<ICandles> {
    const changed =
      !this.chartStock ||
      this.chartStock.symbol !== symbol ||
      this.chartStock.exchange !== exchange ||
      this.chartStock.range !== range;

    this.setChartActive(symbol, exchange, range);
    if (changed || !this.fetchedCandles) {
      await this.pollMarket();
    }
    return this.fetchedCandles as ICandles;
  }

  peekCandles(): ICandles | null {
    return this.fetchedCandles;
  }
}
