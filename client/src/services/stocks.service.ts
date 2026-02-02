import { inject, Injectable } from '@angular/core';
import { IMarketDepth, IStock } from '../types/stocks.types';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom, map, shareReplay, type Observable } from 'rxjs';
import { IWatchlistSymbol, IWatchlistSymbolInfo } from '../types';
import e from 'express';

@Injectable({
  providedIn: 'root',
})
export class StocksService {
  private http = inject(HttpClient);
  private stocks$: Observable<IStock[]> = this.http
    .get<IStock[]>('/market/symbols')
    .pipe(shareReplay(1));
  private stocksListWithData: IWatchlistSymbolInfo[] = [];
  private fetchedData: { symbols: IWatchlistSymbolInfo[] } = { symbols: [] };
  private symbolsToFetchSet: IWatchlistSymbol[] = [];
  private depthStock: IWatchlistSymbol | null = null;
  private fetchedDepthData: IMarketDepth | null = null;

  constructor() {
    // Polling
    setInterval(async () => {
      if (this.stocksListWithData.length > 0) {
        this.fetchedData = await firstValueFrom(
          this.http.post<{ symbols: IWatchlistSymbolInfo[] }>('/market/symbols/status', {
            symbols: this.symbolsToFetchSet,
          }),
        );
      }
      if (this.depthStock !== null) {
        this.fetchedDepthData = await firstValueFrom(
          this.http.get<IMarketDepth>('/market/depth', {
            params: {
              symbol: this.depthStock.symbol,
              exchange: this.depthStock.exchange,
            },
          }),
        );
      }
    }, 500);
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

  async getStocksData(symbols: IWatchlistSymbol[]): Promise<IWatchlistSymbolInfo[]> {
    // process the fetched data
    const fetchedDataMap = new Map<string, IWatchlistSymbolInfo>();
    for (const item of this.fetchedData.symbols) {
      fetchedDataMap.set(item.symbol, item);
    }
    this.stocksListWithData = this.stocksListWithData.map((stock) => {
      const updatedData = fetchedDataMap.get(stock.symbol);
      return updatedData ? updatedData : stock;
    });
    // Update the stocks list to match the requested symbols
    const incomingSymbolSet = new Set(symbols.map((s) => s.symbol));
    this.stocksListWithData = this.stocksListWithData.filter((stock) =>
      incomingSymbolSet.has(stock.symbol),
    );
    const existingSymbolSet = new Set(this.stocksListWithData.map((stock) => stock.symbol));
    const newSymbols = symbols.filter((s) => !existingSymbolSet.has(s.symbol));

    // Fetch data for new symbols and add them
    if (newSymbols.length > 0) {
      const fetchedData = await firstValueFrom(
        this.http.post<{ symbols: IWatchlistSymbolInfo[] }>('/market/symbols/status', {
          symbols: newSymbols,
        }),
      );

      this.stocksListWithData.push(...fetchedData.symbols);
    }
    this.symbolsToFetchSet = symbols;

    return this.stocksListWithData;
  }

  async getDepthData(symbol: string, exchange: string): Promise<IMarketDepth> {
    if (
      this.depthStock === null ||
      this.depthStock.symbol !== symbol ||
      this.depthStock.exchange !== exchange
    ) {
      this.depthStock = { symbol, exchange };
      this.fetchedDepthData = await firstValueFrom(
        this.http.get<IMarketDepth>('/market/depth', {
          params: {
            exchange,
            symbol,
          },
        }),
      );
    }
    return this.fetchedDepthData as IMarketDepth;
  }
}
