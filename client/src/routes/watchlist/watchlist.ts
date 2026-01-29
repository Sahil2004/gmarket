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
  watchlists = computed(async () => {
    try {
      const watchlist = (await firstValueFrom(
        this.watchlistService.getWatchlist(this.currentWatchlistIdx() + 1),
      )) as IWatchlist;
      let watchlists: IWatchlistSymbol[][] = [];
      for (let i = 0; i < 10; i++) {
        if (i != this.currentWatchlistIdx()) {
          watchlists.push([]);
        } else {
          watchlists.push(watchlist.symbols);
        }
      }
      return watchlists;
    } catch (err) {
      return new Array(10).fill([]) as IWatchlistSymbol[][];
    }
  });

  getStockSymbols(stocks: IStock[]): string[] {
    return stocks.map((stock) => stock.symbol);
  }

  updateWatchlists() {
    // this.watchlists.set(this.currentUser?.watchlists ?? []);
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
      // const res = this.userService.removeStockFromWatchlist(this.currentWatchlistIdx(), stock);
      // if (!res) {
      //   let snackBarRef = this._snackBar.open('Failed to remove stock from watchlist', 'Close', {
      //     duration: 3000,
      //   });
      //   snackBarRef.onAction().subscribe(() => {
      //     snackBarRef.dismiss();
      //   });
      //   return false;
      // }
      // this.updateWatchlists();
      // return true;
    };
  }

  watchlistChangeHandler() {
    return (event: { index: number }) => this.currentWatchlistIdx.set(event.index);
  }
}
