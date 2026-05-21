import { Component, inject, OnDestroy, OnInit, signal } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { TradingService } from '../../services';
import { ITradingSnapshot } from '../../types';
import { CurrencyPipe } from '@angular/common';

@Component({
  selector: 'funds',
  templateUrl: 'funds.html',
  imports: [
    ReactiveFormsModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    CurrencyPipe,
  ],
})
export class Funds implements OnInit, OnDestroy {
  private tradingService = inject(TradingService);
  private unsubscribeSnapshot?: () => void;

  snapshot = signal<ITradingSnapshot | null>(null);

  fundForm = new FormGroup({
    bank_account_id: new FormControl('', { nonNullable: true, validators: [Validators.required] }),
    amount: new FormControl(0, { nonNullable: true, validators: [Validators.required, Validators.min(1)] }),
    type: new FormControl<'deposit' | 'withdraw'>('deposit', { nonNullable: true }),
  });

  ngOnInit(): void {
    this.tradingService.startPolling();
    this.unsubscribeSnapshot = this.tradingService.onSnapshotUpdate((s) => {
      this.snapshot.set(s);
      if (s.bank_accounts.length > 0 && !this.fundForm.controls.bank_account_id.value) {
        this.fundForm.controls.bank_account_id.setValue(s.bank_accounts[0].id);
      }
    });
    void this.tradingService.getSnapshot().then((s) => this.snapshot.set(s));
  }

  ngOnDestroy(): void {
    this.unsubscribeSnapshot?.();
    this.tradingService.stopPolling();
  }

  async transferFunds(): Promise<void> {
    if (this.fundForm.invalid) return;
    const { bank_account_id, amount, type } = this.fundForm.getRawValue();
    if (type === 'deposit') {
      await this.tradingService.depositFunds(bank_account_id, amount);
    } else {
      await this.tradingService.withdrawFunds(bank_account_id, amount);
    }
  }
}
