import { Component, Input, output } from '@angular/core';
import { MatTabsModule } from '@angular/material/tabs';
import { StocksList } from '../stocks-list/stocks-list';
import { IWatchlist } from '../../../types';
import { AsyncPipe } from '@angular/common';

@Component({
  selector: 'watchlist-tabs',
  templateUrl: 'watchlist-tabs.html',
  imports: [MatTabsModule, StocksList, AsyncPipe],
})
export class WatchlistTabs {
  @Input() tabChangeHandler: (event: { index: number }) => void = (event) => {};
  @Input({ required: true }) watchlist!: Promise<IWatchlist>;
  @Input() buyHandler: (stock: string) => void = () => {};
  @Input() sellHandler: (stock: string) => void = () => {};
  @Input() removeFromWatchlistHandler: (symbol: string, exchange: string) => void = () => {};

  numbers: number[] = Array.from({ length: 10 }, (_, i) => i + 1);
}
