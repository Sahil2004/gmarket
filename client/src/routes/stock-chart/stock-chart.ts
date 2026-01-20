import { Component, computed, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { toSignal } from '@angular/core/rxjs-interop';
import { IChartApiResponse } from '../../types/stocks.types';

@Component({
  selector: 'stock-chart',
  templateUrl: 'stock-chart.html',
})
export class StockChart {
  private route = inject(ActivatedRoute);
  private data = toSignal(this.route.data);
  stockSymbol = computed(() => this.route.snapshot.paramMap.get('symbol'));
  chartData = computed(() => this.data()?.['stockData'] as IChartApiResponse);
  ltp = computed(() => this.chartData()?.chart.result?.[0].meta.regularMarketPrice);
}
