import { Component, effect, inject, input, output, signal } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AlgoService } from '../../../services';
import { AlgoSignal, IAlgoConfig, IAlgoIndicators } from '../../../types';
import { ChartRange } from '../../../types/stocks.types';

@Component({
  selector: 'stock-algo-panel',
  templateUrl: 'stock-algo-panel.html',
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatSlideToggleModule,
    MatButtonToggleModule,
  ],
})
export class StockAlgoPanel {
  symbol = input.required<string>();
  exchange = input.required<string>();
  range = input.required<ChartRange>();
  indicators = input<IAlgoIndicators | null>(null);

  saved = output<IAlgoConfig>();

  private algoService = inject(AlgoService);
  private snackBar = inject(MatSnackBar);

  loading = signal(false);
  saving = signal(false);

  form = new FormGroup({
    enabled: new FormControl(false, { nonNullable: true }),
    rsi_enabled: new FormControl(true, { nonNullable: true }),
    rsi_period: new FormControl(14, { nonNullable: true, validators: [Validators.min(2), Validators.max(100)] }),
    rsi_overbought: new FormControl(70, { nonNullable: true }),
    rsi_oversold: new FormControl(30, { nonNullable: true }),
    ma_enabled: new FormControl(true, { nonNullable: true }),
    ma_fast_period: new FormControl(9, { nonNullable: true, validators: [Validators.min(2)] }),
    ma_slow_period: new FormControl(21, { nonNullable: true, validators: [Validators.min(3)] }),
  });

  constructor() {
    effect(() => {
      const sym = this.symbol();
      const ex = this.exchange();
      if (sym && ex) void this.loadConfig(sym, ex);
    });
  }

  async loadConfig(symbol: string, exchange: string): Promise<void> {
    this.loading.set(true);
    try {
      const cfg = await this.algoService.getConfig(symbol, exchange);
      this.form.patchValue(cfg, { emitEvent: false });
    } finally {
      this.loading.set(false);
    }
  }

  async save(): Promise<void> {
    if (this.form.invalid) return;
    this.saving.set(true);
    try {
      const cfg = await this.algoService.saveConfig({
        symbol: this.symbol(),
        exchange: this.exchange(),
        ...this.form.getRawValue(),
      });
      this.snackBar.open('Algo settings saved', 'Close', { duration: 2000 });
      this.saved.emit(cfg);
    } finally {
      this.saving.set(false);
    }
  }

  signalClass(sig: AlgoSignal | undefined): string {
    switch (sig) {
      case 'buy':
      case 'bullish':
        return 'text-emerald-400';
      case 'sell':
      case 'bearish':
        return 'text-red-400';
      default:
        return 'text-gray-400';
    }
  }

  rsiBarWidth(rsi: number): number {
    return Math.min(100, Math.max(0, rsi));
  }
}
