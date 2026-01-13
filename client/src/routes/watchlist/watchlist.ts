import { Component, inject, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MatSidenavModule } from '@angular/material/sidenav';
import { Search, WatchlistTabs } from '../../components';
import { StocksService } from '../../services';
import { IStock } from '../../types/stocks.types';
import type { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
  selector: 'watchlist',
  templateUrl: 'watchlist.html',
  imports: [RouterOutlet, MatSidenavModule, Search, WatchlistTabs, AsyncPipe],
})
export class Watchlist implements OnInit {
  private stockService = inject(StocksService);

  stockList$!: Observable<IStock[] | undefined>;

  ngOnInit(): void {
    this.stockList$ = this.stockService.getAllStocks();
  }

  getStockSymbols(stocks: IStock[]): string[] {
    return stocks.map((stock) => stock.symbol);
  }

  addStockToWatchlist() {
    return (query: string) => {
      alert(`Add ${query} to watchlist`);
    };
  }
}
