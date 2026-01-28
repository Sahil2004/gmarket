import { inject, Injectable } from '@angular/core';
import { IChartApiResponse, IStock } from '../types/stocks.types';
import { HttpClient } from '@angular/common/http';
import { map, shareReplay, type Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class StocksService {
  private http = inject(HttpClient);
  private stocks$!: Observable<IStock[]>;

  getAllStocks(): Observable<IStock[]> {
    this.stocks$ = this.http.get<IStock[]>('/market/symbols').pipe(shareReplay(1));
    return this.stocks$;
  }

  isAStockSymbol(symbol: string): Observable<boolean> {
    return this.getAllStocks().pipe(
      shareReplay(1),
      map((stocks) => stocks.some((stock) => stock.symbol === symbol)),
    );
  }

  getChartData(symbol: string, interval: string, range: string): Observable<IChartApiResponse> {
    const url = `https://query1.finance.yahoo.com/v8/finance/chart/${symbol}?range=${range}&interval=${interval}`;
    const proxiedUrl = `https://api.allorigins.win/raw?url=${encodeURIComponent(url)}`;
    return this.http.get<IChartApiResponse>(proxiedUrl);
  }
}
