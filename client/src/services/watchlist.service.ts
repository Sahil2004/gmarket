import { HttpClient, HttpContext } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { IError, IWatchlist } from '../types';
import { Observable } from 'rxjs';
import { SKIP_TOAST } from '../contexts';

@Injectable({
  providedIn: 'root',
})
export class WatchlistService {
  private http = inject(HttpClient);

  getWatchlist(watchlistIndex: number): Observable<IWatchlist | IError> {
    return this.http.get<IWatchlist | IError>(`/watchlists/${watchlistIndex}`, {
      context: new HttpContext().set(SKIP_TOAST, true),
    });
  }

  addStockToWatchlist(
    watchlistIndex: number,
    stockSymbol: string,
    exchange: string,
  ): Observable<null | IError> {
    const stockDetails = {
      exchange,
      symbol: stockSymbol,
    };
    return this.http.post<null | IError>(`/watchlists/${watchlistIndex}/symbols`, stockDetails);
  }

  removeStockFromWatchlist(
    watchlistIndex: number,
    stockSymbol: string,
    exchange: string,
  ): Observable<null | IError> {
    const stockDetails = {
      exchange,
      symbol: stockSymbol,
    };
    return this.http.delete<null | IError>(`/watchlists/${watchlistIndex}/symbols`, {
      body: stockDetails,
    });
  }
}
