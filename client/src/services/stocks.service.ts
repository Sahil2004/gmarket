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
