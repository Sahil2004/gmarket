import { Component, Input, ViewChildren, type QueryList } from '@angular/core';
import { MatTabsModule } from '@angular/material/tabs';
import { StocksList } from '../stocks-list/stocks-list';
import { IWatchlist, IWatchlistSymbol, IWatchlistSymbolInfo } from '../../../types';
import { AsyncPipe } from '@angular/common';

@Component({
  selector: 'watchlist-tabs',
  templateUrl: 'watchlist-tabs.html',
  imports: [MatTabsModule, StocksList, AsyncPipe],
})
export class WatchlistTabs {
  @Input() tabChangeHandler: (event: { index: number }) => void = (event) => {};
  @Input() selectedIdx = 0;
  @Input({ required: true }) watchlist!: Promise<IWatchlist>;
  @Input() buyHandler: (symbol: string, exchange: string) => void = () => {};
  @Input() sellHandler: (symbol: string, exchange: string) => void = () => {};
  @Input() removeFromWatchlistHandler: (symbol: string, exchange: string) => void = () => {};
  @Input() getStockData: (symbols: IWatchlistSymbol[]) => Promise<IWatchlistSymbolInfo[]> = () => {
    return Promise.resolve([]);
  };
  @Input() throttlingTimeMs: number = 2 * 1000; // 2 seconds

  numbers: number[] = Array.from({ length: 10 }, (_, i) => i + 1);

  @ViewChildren(StocksList) stocksLists!: QueryList<StocksList>;

  onTabAnimationDone(): void {
    this.stocksLists.first?.recheckViewport();
  }
}
