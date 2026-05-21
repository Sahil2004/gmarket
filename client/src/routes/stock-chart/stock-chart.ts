import {
  afterNextRender,
  Component,
  computed,
  effect,
  ElementRef,
  inject,
  OnDestroy,
  signal,
  viewChild,
} from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { toSignal } from '@angular/core/rxjs-interop';
import { ChartRange, ICandle, IMarketDepth } from '../../types';
import { IAlgoIndicators } from '../../types/algo.types';
import { AlgoService, StocksService } from '../../services';
import { StockAlgoPanel } from '../../components';
import { DESIGN_SYSTEM } from '../../config';
import {
  CandlestickSeries,
  createChart,
  LineSeries,
  type IChartApi,
  type ISeriesApi,
  type UTCTimestamp,
} from 'lightweight-charts';

@Component({
  selector: 'stock-chart',
  templateUrl: 'stock-chart.html',
  imports: [StockAlgoPanel],
})
export class StockChart implements OnDestroy {
  private ds = inject(DESIGN_SYSTEM);
  private stockService = inject(StocksService);
  private algoService = inject(AlgoService);
  private route = inject(ActivatedRoute);
  private data = toSignal(this.route.data);
  private params = toSignal(this.route.paramMap);

  chartHost = viewChild.required<ElementRef<HTMLDivElement>>('chartHost');

  stockSymbol = computed(() => this.params()?.get('symbol') ?? '');
  exchange = computed(() => this.params()?.get('exchange') ?? '');
  depthData = signal<IMarketDepth>({
    bids: [],
    asks: [],
    ltp: 0,
    exchange: '',
    symbol: '',
  });

  ranges: ChartRange[] = ['1D', '1W', '1M', '1Y'];
  selectedRange = signal<ChartRange>('1D');
  currentCandle = signal<ICandle | null>(null);
  algoIndicators = signal<IAlgoIndicators | null>(null);

  maxBidQty = computed(() => Math.max(...this.depthData().bids.map((b) => b.qty), 1));
  maxAskQty = computed(() => Math.max(...this.depthData().asks.map((a) => a.qty), 1));

  private chart: IChartApi | null = null;
  private candleSeries: ISeriesApi<'Candlestick'> | null = null;
  private maFastSeries: ISeriesApi<'Line'> | null = null;
  private maSlowSeries: ISeriesApi<'Line'> | null = null;
  private lastCandleTime: UTCTimestamp | null = null;
  private uiIntervalId: ReturnType<typeof setInterval> | null = null;
  private resizeObserver: ResizeObserver | null = null;
  private chartReady = signal(false);

  qtyBarWidth(qty: number, maxQty: number): number {
    return Math.min(100, (qty / maxQty) * 100);
  }

  constructor() {
    this.depthData.set(this.data()?.['depthData'] as IMarketDepth);

    afterNextRender(() => {
      this.initChart();
      this.chartReady.set(true);
      this.activateMarketTargets();
      void this.loadCandles();
      void this.loadAlgoIndicators();
      this.startUiRefresh();
    });

    effect(() => {
      const symbol = this.stockSymbol();
      const exchange = this.exchange();
      const range = this.selectedRange();
      if (!this.chartReady() || !symbol || !exchange) return;
      this.activateMarketTargets();
      this.lastCandleTime = null;
      void this.loadCandles();
      void this.loadAlgoIndicators();
    });
  }

  ngOnDestroy(): void {
    if (this.uiIntervalId) clearInterval(this.uiIntervalId);
    this.resizeObserver?.disconnect();
    this.chart?.remove();
    this.stockService.setDepthActive(null, null);
    this.stockService.setChartActive(null, null, null);
  }

  setRange(range: ChartRange): void {
    if (this.selectedRange() === range) return;
    this.selectedRange.set(range);
    this.lastCandleTime = null;
    this.stockService.setChartActive(this.stockSymbol(), this.exchange(), range);
    void this.loadAlgoIndicators();
  }

  onAlgoSaved(): void {
    void this.loadAlgoIndicators();
  }

  private activateMarketTargets(): void {
    const symbol = this.stockSymbol();
    const exchange = this.exchange();
    if (!symbol || !exchange) return;
    this.stockService.setDepthActive(symbol, exchange);
    this.stockService.setChartActive(symbol, exchange, this.selectedRange());
  }

  private async loadAlgoIndicators(): Promise<void> {
    const symbol = this.stockSymbol();
    const exchange = this.exchange();
    if (!symbol || !exchange) return;
    try {
      const ind = await this.algoService.getIndicators(symbol, exchange, this.selectedRange());
      this.algoIndicators.set(ind);
      this.updateMaLines(ind);
    } catch {
      this.algoIndicators.set(null);
    }
  }

  private startUiRefresh(): void {
    this.uiIntervalId = setInterval(() => {
      const depth = this.stockService.peekDepthData();
      if (depth) this.depthData.set(depth);
      const candles = this.stockService.peekCandles();
      if (candles?.candles?.length) {
        this.applyRealtimeCandles(candles.candles);
      }
      void this.loadAlgoIndicators();
    }, this.ds.devConfig.throttlingTimeMs);
  }

  private initChart(): void {
    const host = this.chartHost().nativeElement;
    this.chart = createChart(host, {
      layout: { background: { color: 'transparent' }, textColor: '#9ca3af' },
      grid: {
        vertLines: { color: 'rgba(55, 65, 81, 0.4)' },
        horzLines: { color: 'rgba(55, 65, 81, 0.4)' },
      },
      rightPriceScale: { borderColor: 'rgba(75, 85, 99, 0.6)' },
      timeScale: { borderColor: 'rgba(75, 85, 99, 0.6)', timeVisible: true },
      crosshair: { mode: 1 },
    });

    this.candleSeries = this.chart.addSeries(CandlestickSeries, {
      upColor: '#22c55e',
      downColor: '#ef4444',
      borderUpColor: '#22c55e',
      borderDownColor: '#ef4444',
      wickUpColor: '#22c55e',
      wickDownColor: '#ef4444',
    });

    this.maFastSeries = this.chart.addSeries(LineSeries, {
      color: '#f59e0b',
      lineWidth: 2,
      title: 'MA Fast',
    });
    this.maSlowSeries = this.chart.addSeries(LineSeries, {
      color: '#3b82f6',
      lineWidth: 2,
      title: 'MA Slow',
    });

    this.resizeObserver = new ResizeObserver(() => {
      if (this.chart) {
        this.chart.applyOptions({ width: host.clientWidth, height: host.clientHeight });
      }
    });
    this.resizeObserver.observe(host);
    this.chart.applyOptions({ width: host.clientWidth, height: host.clientHeight });
  }

  private updateMaLines(ind: IAlgoIndicators): void {
    if (!this.maFastSeries || !this.maSlowSeries) return;
    const toLine = (pts: { time: number; value: number }[]) =>
      pts.map((p) => ({ time: p.time as UTCTimestamp, value: p.value }));

    if (ind.config.ma_enabled) {
      this.maFastSeries.setData(toLine(ind.ma_fast));
      this.maSlowSeries.setData(toLine(ind.ma_slow));
    } else {
      this.maFastSeries.setData([]);
      this.maSlowSeries.setData([]);
    }
  }

  private toChartCandles(candles: ICandle[]) {
    return candles.map((c) => ({
      time: c.time as UTCTimestamp,
      open: c.open,
      high: c.high,
      low: c.low,
      close: c.close,
    }));
  }

  private async loadCandles(): Promise<void> {
    const symbol = this.stockSymbol();
    const exchange = this.exchange();
    if (!symbol || !exchange || !this.candleSeries) return;

    const data = await this.stockService.getCandles(symbol, exchange, this.selectedRange());
    const formatted = this.toChartCandles(data.candles);
    this.candleSeries.setData(formatted);
    this.lastCandleTime = formatted.length ? formatted[formatted.length - 1].time : null;
    this.setCurrentCandle(data.candles);
    this.chart?.timeScale().fitContent();
  }

  private setCurrentCandle(candles: ICandle[]): void {
    this.currentCandle.set(candles.length > 0 ? candles[candles.length - 1] : null);
  }

  private applyRealtimeCandles(candles: ICandle[]): void {
    if (!this.candleSeries || candles.length === 0) return;

    const formatted = this.toChartCandles(candles);
    const last = formatted[formatted.length - 1];
    this.setCurrentCandle(candles);

    if (this.lastCandleTime === null) {
      this.candleSeries.setData(formatted);
      this.lastCandleTime = last.time;
      return;
    }

    if (last.time === this.lastCandleTime) {
      this.candleSeries.update(last);
      return;
    }

    this.candleSeries.update(last);
    this.lastCandleTime = last.time;
  }
}
