import { inject, Injectable } from '@angular/core';
import { IChartApiResponse, IStock } from '../types/stocks.types';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom, map, shareReplay, type Observable } from 'rxjs';
import { IWatchlistSymbol, IWatchlistSymbolInfo } from '../types';

@Injectable({
  providedIn: 'root',
})
export class StocksService {
  private http = inject(HttpClient);
  private stocks$: Observable<IStock[]> = this.http
    .get<IStock[]>('/market/symbols')
    .pipe(shareReplay(1));
  private stocksListWithData: IWatchlistSymbolInfo[] = [];

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

    return this.stocksListWithData;
  }

  getChartData(
    symbol: string,
    exchange: string,
    interval: string,
    range: string,
  ): Observable<IChartApiResponse> {
    return this.http.get<IChartApiResponse>('/market/chart', {
      params: {
        exchange,
        symbol,
        range,
        interval,
      },
    });
  }
}
