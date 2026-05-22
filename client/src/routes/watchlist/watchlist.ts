import { Component, computed, inject, OnDestroy, OnInit, signal } from '@angular/core';
import { ActivatedRoute, NavigationEnd, Router, RouterOutlet } from '@angular/router';
import { filter, Subscription } from 'rxjs';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatDialog } from '@angular/material/dialog';
import { Search, WatchlistTabs, OrderDialog } from '../../components';
import { StocksService, WatchlistService } from '../../services';
import { IStock } from '../../types/stocks.types';
import { MatSnackBar } from '@angular/material/snack-bar';
import { toSignal } from '@angular/core/rxjs-interop';
import { IOrderDialogData, IWatchlist, IWatchlistSymbol } from '../../types';
import { DESIGN_SYSTEM } from '../../config';

@Component({
  selector: 'watchlist',
  templateUrl: 'watchlist.html',
  imports: [RouterOutlet, MatSidenavModule, Search, WatchlistTabs],
})
export class Watchlist implements OnInit, OnDestroy {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private routerSub?: Subscription;
  private data = toSignal(this.route.data);
  stockList = computed(() => this.data()?.['stockData'] as IStock[]);
  private stockService = inject(StocksService);
  private watchlistService = inject(WatchlistService);
  private _snackBar = inject(MatSnackBar);
  private _dialog = inject(MatDialog);
  private ds = inject(DESIGN_SYSTEM);

  throttlingTimeMs = computed(() => this.ds.devConfig.throttlingTimeMs);

  currentWatchlistIdx = signal<number>(0);
  watchlistUpdatedAt = signal<Date | null>(null);
  watchlist = computed(async () => {
    const update = this.watchlistUpdatedAt();
    try {
      const watchlist = (await this.watchlistService.getWatchlist(
        this.currentWatchlistIdx() + 1,
      )) as IWatchlist;
      return watchlist;
    } catch (err) {
      return { index: this.currentWatchlistIdx() + 1, symbols: [] } as IWatchlist;
    }
  });

  addStockToWatchlist() {
    return (symbol: string, exchange: string) => {
      return this.stockService.isAStockSymbol(symbol).subscribe(async (decision) => {
        if (decision) {
          try {
            await this.watchlistService.addStockToWatchlist(
              this.currentWatchlistIdx() + 1,
              symbol,
              exchange,
            );
            this.updateWatchlists();
          } catch (err) {
            this._snackBar.open('Invalid stock symbol', 'Close', { duration: 3000 });
          }
        } else {
          this._snackBar.open('Invalid stock symbol', 'Close', { duration: 3000 });
        }
      });
    };
  }

  private openOrderDialog(symbol: string, exchange: string, side: 'buy' | 'sell') {
    const data: IOrderDialogData = { symbol, exchange, side };
    this._dialog.open(OrderDialog, { data, width: '400px' });
  }

  buyHandler() {
    return (symbol: string, exchange: string) => {
      this.openOrderDialog(symbol, exchange, 'buy');
    };
  }

  sellHandler() {
    return (symbol: string, exchange: string) => {
      this.openOrderDialog(symbol, exchange, 'sell');
    };
  }

  removeFromWatchlistHandler() {
    return async (symbol: string, exchange: string) => {
      try {
        await this.watchlistService.removeStockFromWatchlist(
          this.currentWatchlistIdx() + 1,
          symbol,
          exchange,
        );
        this.updateWatchlists();
      } catch (err) {}
    };
  }

  watchlistChangeHandler() {
    return (event: { index: number }) => {
      this.currentWatchlistIdx.set(event.index);
      this.updateWatchlists();
    };
  }

  getStockDataHandler() {
    return (symbols: IWatchlistSymbol[]) => {
      return this.stockService.getStocksData(symbols);
    };
  }

  updateWatchlists() {
    this.watchlistUpdatedAt.set(new Date());
  }

  ngOnInit(): void {
    this.syncWatchlistPolling();
    this.routerSub = this.router.events
      .pipe(filter((e) => e instanceof NavigationEnd))
      .subscribe(() => this.syncWatchlistPolling());
  }

  ngOnDestroy(): void {
    this.routerSub?.unsubscribe();
    this.stockService.setWatchlistPolling(false);
    this.stockService.clearMarketPolling();
  }

  private syncWatchlistPolling(): void {
    const onChart = /\/watchlist\/[^/]+\/[^/]+/.test(this.router.url);
    this.stockService.setWatchlistPolling(!onChart);
  }
}
