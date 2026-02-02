import { Component, computed, inject, signal } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { toSignal } from '@angular/core/rxjs-interop';
import { IMarketDepth } from '../../types';
import { MatTableModule } from '@angular/material/table';
import { StocksService } from '../../services';
import { DESIGN_SYSTEM } from '../../config';

@Component({
  selector: 'stock-chart',
  templateUrl: 'stock-chart.html',
  imports: [MatTableModule],
})
export class StockChart {
  private ds = inject(DESIGN_SYSTEM);
  private stockService = inject(StocksService);
  private route = inject(ActivatedRoute);
  private data = toSignal(this.route.data);
  private params = toSignal(this.route.paramMap);
  stockSymbol = computed(() => this.params()?.get('symbol'));
  exchange = computed(() => this.params()?.get('exchange'));
  depthData = signal<IMarketDepth>({
    bids: [],
    asks: [],
    ltp: 0,
    exchange: '',
    symbol: '',
  });
  displayedColumns: string[] = ['price', 'orders', 'qty'];

  constructor() {
    this.depthData.set(this.data()?.['depthData'] as IMarketDepth);
    setInterval(async () => {
      const depthData = await this.stockService.getDepthData(
        this.stockSymbol() as string,
        this.exchange() as string,
      );
      this.depthData.set(depthData);
    }, this.ds.devConfig.throttlingTimeMs);
  }
}
