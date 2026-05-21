import { Component, inject, OnDestroy, OnInit, signal } from '@angular/core';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TradingService } from '../../services';
import { IOrder, ITradingSnapshot } from '../../types';
import { DatePipe } from '@angular/common';

@Component({
  selector: 'orders',
  templateUrl: 'orders.html',
  imports: [MatTableModule, MatButtonModule, MatIconModule, DatePipe],
})
export class Orders implements OnInit, OnDestroy {
  private tradingService = inject(TradingService);
  private unsubscribeSnapshot?: () => void;

  snapshot = signal<ITradingSnapshot | null>(null);
  orderCols = ['symbol', 'side', 'type', 'qty', 'price', 'status', 'action'];

  ngOnInit(): void {
    this.tradingService.startPolling();
    this.unsubscribeSnapshot = this.tradingService.onSnapshotUpdate((s) => this.snapshot.set(s));
    void this.tradingService.getSnapshot().then((s) => this.snapshot.set(s));
  }

  ngOnDestroy(): void {
    this.unsubscribeSnapshot?.();
    this.tradingService.stopPolling();
  }

  async cancelOrder(order: IOrder): Promise<void> {
    await this.tradingService.cancelOrder(order.id);
  }
}
