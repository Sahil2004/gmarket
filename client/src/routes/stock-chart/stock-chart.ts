import { Component, computed, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { toSignal } from '@angular/core/rxjs-interop';
import { IMarketDepth } from '../../types';
import { MatTableModule } from '@angular/material/table';

@Component({
  selector: 'stock-chart',
  templateUrl: 'stock-chart.html',
  imports: [MatTableModule],
})
export class StockChart {
  private route = inject(ActivatedRoute);
  private data = toSignal(this.route.data);
  private params = toSignal(this.route.paramMap);
  stockSymbol = computed(() => this.params()?.get('symbol'));
  exchange = computed(() => this.params()?.get('exchange'));
  depthData = computed(() => this.data()?.['depthData'] as IMarketDepth);
  displayedColumns: string[] = ['price', 'orders', 'qty'];
}
