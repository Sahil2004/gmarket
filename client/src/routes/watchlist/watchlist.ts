import { Component, inject, OnInit, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MatSidenavModule } from '@angular/material/sidenav';
import { Search, WatchlistTabs } from '../../components';
import { StocksService, WatchlistService } from '../../services';
import { IStock } from '../../types/stocks.types';
import type { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'watchlist',
  templateUrl: 'watchlist.html',
  imports: [RouterOutlet, MatSidenavModule, Search, WatchlistTabs, AsyncPipe],
})
export class Watchlist implements OnInit {
  private stockService = inject(StocksService);
  private watchlistService = inject(WatchlistService);
  private _snackBar = inject(MatSnackBar);

  currentWatchlistIdx = signal<number>(0);
  // watchlists = signal<string[][]>(this.currentUser?.watchlists ?? []);
  watchlists = signal<string[][]>([]);

  stockList$!: Observable<IStock[] | undefined>;

  ngOnInit(): void {
    this.stockList$ = this.stockService.getAllStocks();
  }

  getStockSymbols(stocks: IStock[]): string[] {
    return stocks.map((stock) => stock.symbol);
  }

  updateWatchlists() {
    // this.watchlists.set(this.currentUser?.watchlists ?? []);
  }

  addStockToWatchlist() {
    return (query: string) => {
      return this.stockService.isAStockSymbol(query).subscribe((decision) => {
        // if (decision) {
        //   const res = this.userService.addStockToWatchlist(this.currentWatchlistIdx(), query);
        //   if (!res) {
        //     let snackBarRef = this._snackBar.open('Failed to add stock to watchlist', 'Close', {
        //       duration: 3000,
        //     });
        //     snackBarRef.onAction().subscribe(() => {
        //       snackBarRef.dismiss();
        //     });
        //     return false;
        //   }
        //   this.updateWatchlists();
        //   return true;
        // } else {
        //   let snackBarRef = this._snackBar.open('Stock symbol not found', 'Close', {
        //     duration: 3000,
        //   });
        //   snackBarRef.onAction().subscribe(() => {
        //     snackBarRef.dismiss();
        //   });
        //   return false;
        // }
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
    return (stock: string) => {
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
