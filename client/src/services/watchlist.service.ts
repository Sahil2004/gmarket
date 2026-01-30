import { HttpClient, HttpContext } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { IError, IWatchlist } from '../types';
import { firstValueFrom, Observable } from 'rxjs';
import { SKIP_TOAST } from '../contexts';

@Injectable({
  providedIn: 'root',
})
export class WatchlistService {
  private http = inject(HttpClient);
  private watchlists = new Array<IWatchlist>(10);

  async getWatchlist(watchlistIndex: number) {
    if (this.watchlists[watchlistIndex]) {
      return this.watchlists[watchlistIndex];
    } else {
      try {
        const obs = this.http.get<IWatchlist | IError>(`/watchlists/${watchlistIndex}`, {
          context: new HttpContext().set(SKIP_TOAST, true),
        });
        const watchlist = await firstValueFrom(obs);
        this.watchlists[watchlistIndex] = watchlist as IWatchlist;
        return this.watchlists[watchlistIndex];
      } catch {
        this.watchlists[watchlistIndex] = { index: watchlistIndex, symbols: [] };
        return this.watchlists[watchlistIndex];
      }
    }
  }

  async addStockToWatchlist(watchlistIndex: number, stockSymbol: string, exchange: string) {
    const stockDetails = {
      exchange,
      symbol: stockSymbol,
    };
    const obs = this.http.post<null | IError>(
      `/watchlists/${watchlistIndex}/symbols`,
      stockDetails,
    );
    try {
      await firstValueFrom(obs);
      this.watchlists[watchlistIndex].symbols.push({
        symbol: stockSymbol,
        exchange: exchange,
      });
    } catch (err) {}
  }

  async removeStockFromWatchlist(watchlistIndex: number, stockSymbol: string, exchange: string) {
    const stockDetails = {
      exchange,
      symbol: stockSymbol,
    };
    try {
      const obs = this.http.delete<null | IError>(`/watchlists/${watchlistIndex}/symbols`, {
        body: stockDetails,
      });
      await firstValueFrom(obs);
      this.watchlists[watchlistIndex].symbols = this.watchlists[watchlistIndex].symbols.filter(
        (s) => !(s.symbol === stockSymbol && s.exchange === exchange),
      );
    } catch (err) {}
  }
}
