import { Component, inject, OnDestroy, OnInit, signal } from '@angular/core';
import { MatTabsModule } from '@angular/material/tabs';
import { MatTableModule } from '@angular/material/table';
import { TradingService } from '../../services';
import { IHolding, ITradingSnapshot } from '../../types';
import { CurrencyPipe } from '@angular/common';

@Component({
  selector: 'portfolio',
  templateUrl: 'portfolio.html',
  imports: [MatTabsModule, MatTableModule, CurrencyPipe],
})
export class Portfolio implements OnInit, OnDestroy {
  private tradingService = inject(TradingService);
  private unsubscribeSnapshot?: () => void;

  snapshot = signal<ITradingSnapshot | null>(null);
  holdingCols = ['symbol', 'qty', 'avg', 'ltp', 'total_pnl', 'value'];
  positionCols = ['symbol', 'qty', 'avg', 'ltp', 'day_pnl', 'value'];

  ngOnInit(): void {
    this.tradingService.startPolling();
    this.unsubscribeSnapshot = this.tradingService.onSnapshotUpdate((s) => this.snapshot.set(s));
    void this.tradingService.getSnapshot().then((s) => this.snapshot.set(s));
  }

  ngOnDestroy(): void {
    this.unsubscribeSnapshot?.();
    this.tradingService.stopPolling();
  }

  pnlClass(value: number): string {
    return value >= 0 ? 'text-green-600' : 'text-red-600';
  }
}
