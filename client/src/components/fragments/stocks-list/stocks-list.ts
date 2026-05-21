import {
  AfterViewInit,
  ChangeDetectorRef,
  Component,
  inject,
  Input,
  OnChanges,
  OnDestroy,
  SimpleChanges,
  ViewChild,
  ViewEncapsulation,
} from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { RouterLink } from '@angular/router';
import { IWatchlistSymbol, IWatchlistSymbolInfo } from '../../../types';
import { CdkVirtualScrollViewport, ScrollingModule } from '@angular/cdk/scrolling';

@Component({
  selector: 'stocks-list',
  templateUrl: 'stocks-list.html',
  styleUrl: 'stocks-list.css',
  encapsulation: ViewEncapsulation.None,
  imports: [MatListModule, RouterLink, MatButtonModule, MatIconModule, ScrollingModule],
})
export class StocksList implements OnChanges, AfterViewInit, OnDestroy {
  private cdr = inject(ChangeDetectorRef);
  @Input({ required: true }) stocks: IWatchlistSymbol[] = [];
  @Input() displayActions: boolean = true;
  @Input() buyHandler: (symbol: string, exchange: string) => void = () => {};
  @Input() sellHandler: (symbol: string, exchange: string) => void = () => {};
  @Input() removeFromWatchlistHandler: (symbol: string, exchange: string) => void = () => {};
  @Input() getStockData: (symbols: IWatchlistSymbol[]) => Promise<IWatchlistSymbolInfo[]> = () => {
    return Promise.resolve([]);
  };
  @Input() refreshAfterMs: number = 2 * 1000;

  @ViewChild(CdkVirtualScrollViewport) viewport!: CdkVirtualScrollViewport;

  readonly ITEM_SIZE = 48;

  stockDataMap = new Map<string, IWatchlistSymbolInfo>();
  intervalId: ReturnType<typeof setInterval> | null = null;
  private resizeObserver: ResizeObserver | null = null;

  trackBySymbol(_index: number, stock: IWatchlistSymbol) {
    return `${stock.exchange}:${stock.symbol}`;
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (!changes['stocks']) return;

    this.stockDataMap = new Map();
    if (this.intervalId !== null) {
      clearInterval(this.intervalId);
      this.intervalId = null;
    }
    if (this.stocks.length > 0) {
      void this.updateStockData(this.stocks);
    }
    queueMicrotask(() => this.recheckViewport());
  }

  ngAfterViewInit(): void {
    const el = this.viewport?.elementRef.nativeElement;
    if (el) {
      this.resizeObserver = new ResizeObserver(() => this.recheckViewport());
      this.resizeObserver.observe(el);
    }
    this.recheckViewport();
    if (this.stocks.length > 0) {
      this.onScroll();
    }
  }

  ngOnDestroy(): void {
    if (this.intervalId !== null) clearInterval(this.intervalId);
    this.resizeObserver?.disconnect();
  }

  recheckViewport(): void {
    requestAnimationFrame(() => {
      this.viewport?.checkViewportSize();
    });
  }

  async updateStockData(symbols: IWatchlistSymbol[]) {
    const data = await this.getStockData(symbols);
    for (const item of data) {
      this.stockDataMap.set(`${item.exchange}:${item.symbol}`, item);
    }
    this.cdr.detectChanges();
    this.recheckViewport();
  }

  onScroll() {
    if (!this.viewport) return;

    void this.updateStockData(this.stocks);
    if (this.intervalId !== null) clearInterval(this.intervalId);
    this.intervalId = setInterval(() => {
      void this.updateStockData(this.stocks);
    }, this.refreshAfterMs);
  }

  getStockInfo(stock: IWatchlistSymbol): IWatchlistSymbolInfo | undefined {
    return this.stockDataMap.get(`${stock.exchange}:${stock.symbol}`);
  }

  activated(stock: string): boolean {
    return false;
  }
}
