import { Component, Input, signal, ViewChild } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { RouterLink } from '@angular/router';
import { IWatchlistSymbol, IWatchlistSymbolInfo } from '../../../types';
import { CdkVirtualScrollViewport, ScrollingModule } from '@angular/cdk/scrolling';

@Component({
  selector: 'stocks-list',
  templateUrl: 'stocks-list.html',
  imports: [MatListModule, RouterLink, MatButtonModule, MatIconModule, ScrollingModule],
})
export class StocksList {
  @Input({ required: true }) stocks: IWatchlistSymbol[] = [];
  @Input() displayActions: boolean = true;
  @Input() buyHandler: (stock: string) => void = () => {};
  @Input() sellHandler: (stock: string) => void = () => {};
  @Input() removeFromWatchlistHandler: (symbol: string, exchange: string) => void = () => {};
  @Input() getStockData: (symbols: IWatchlistSymbol[]) => Promise<IWatchlistSymbolInfo[]> = () => {
    return Promise.resolve([]);
  };
  @Input() refreshAfterMs: number = 15 * 1000; // 15 seconds

  @ViewChild(CdkVirtualScrollViewport)
  viewport!: CdkVirtualScrollViewport;

  ITEM_SIZE = 48;

  stockList = signal<IWatchlistSymbolInfo[]>([]);
  intervalId: ReturnType<typeof setInterval> | null = null;

  trackBySymbol(index: number, stock: IWatchlistSymbol) {
    return stock.symbol;
  }

  onScroll() {
    if (!this.viewport) return;

    const scrollOffset = this.viewport.measureScrollOffset();
    const viewportHeight = this.viewport.getViewportSize();

    const startIndex = Math.floor(scrollOffset / this.ITEM_SIZE);
    const endIndex = Math.min(
      this.stocks.length,
      Math.ceil((scrollOffset + viewportHeight) / this.ITEM_SIZE),
    );
    const visibleStocks = this.stocks.slice(startIndex, endIndex);
    console.log(visibleStocks);
    (async () => {
      this.stockList.set(await this.getStockData(visibleStocks));
    })();
    if (this.intervalId !== null) clearInterval(this.intervalId);
    this.intervalId = setInterval(async () => {
      this.stockList.set(await this.getStockData(visibleStocks));
    }, this.refreshAfterMs);
  }

  activated(stock: string): boolean {
    return false;
  }
}
