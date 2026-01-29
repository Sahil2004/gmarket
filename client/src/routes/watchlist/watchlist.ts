import { Component, computed, inject, OnInit, signal } from '@angular/core';
import { ActivatedRoute, RouterOutlet } from '@angular/router';
import { MatSidenavModule } from '@angular/material/sidenav';
import { Search, WatchlistTabs } from '../../components';
import { StocksService, WatchlistService } from '../../services';
import { IStock } from '../../types/stocks.types';
import { firstValueFrom, type Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { MatSnackBar } from '@angular/material/snack-bar';
import { toSignal } from '@angular/core/rxjs-interop';
import { IWatchlist, IWatchlistSymbol } from '../../types';

@Component({
  selector: 'watchlist',
  templateUrl: 'watchlist.html',
  imports: [RouterOutlet, MatSidenavModule, Search, WatchlistTabs, AsyncPipe],
})
export class Watchlist {
  private route = inject(ActivatedRoute);
  private data = toSignal(this.route.data);
  stockList = computed(() => this.data()?.['stockData'] as IStock[]);
  private stockService = inject(StocksService);
  private watchlistService = inject(WatchlistService);
  private _snackBar = inject(MatSnackBar);

  currentWatchlistIdx = signal<number>(0);
  watchlistUpdatedAt = signal<Date | null>(null);
  watchlist = computed(async () => {
    const update = this.watchlistUpdatedAt();
    try {
      const watchlist = (await firstValueFrom(
        this.watchlistService.getWatchlist(this.currentWatchlistIdx() + 1),
      )) as IWatchlist;
      return watchlist;
    } catch (err) {
      return { index: this.currentWatchlistIdx() + 1, symbols: [] } as IWatchlist;
    }
  });

  getStockSymbols(stocks: IStock[]): string[] {
    return stocks.map((stock) => stock.symbol);
  }

  updateWatchlists() {
    this.watchlistUpdatedAt.set(new Date());
  }

  addStockToWatchlist() {
    return (symbol: string, exchange: string) => {
      return this.stockService.isAStockSymbol(symbol).subscribe((decision) => {
        if (decision) {
          this.watchlistService
            .addStockToWatchlist(this.currentWatchlistIdx() + 1, symbol, exchange)
            .subscribe((res) => {
              this.updateWatchlists();
            });
        } else {
          let snackBarRef = this._snackBar.open('Invalid stock symbol', 'Close', {
            duration: 3000,
          });
          snackBarRef.onAction().subscribe(() => {
            snackBarRef.dismiss();
          });
        }
      });
    };
  }

  buyHandler() {
    return (stock: string) => {
      alert(`Buying stock: ${stock}`);
    };
  }

  sellHandler() {
    return (stock: string) => {
      alert(`Selling stock: ${stock}`);
    };
  }

  removeFromWatchlistHandler() {
    return (symbol: string, exchange: string) => {
      this.watchlistService
        .removeStockFromWatchlist(this.currentWatchlistIdx() + 1, symbol, exchange)
        .subscribe({
          next: (res) => {
            this.updateWatchlists();
          },
        });
    };
  }

  watchlistChangeHandler() {
    return (event: { index: number }) => this.currentWatchlistIdx.set(event.index);
  }
}
