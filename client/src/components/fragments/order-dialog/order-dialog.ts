import { Component, inject, OnInit, signal } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatSelectModule } from '@angular/material/select';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IOrderDialogData } from '../../../types';
import { TradingService } from '../../../services';
import { debounceTime } from 'rxjs/operators';

@Component({
  selector: 'order-dialog',
  templateUrl: 'order-dialog.html',
  imports: [
    MatDialogModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonToggleModule,
    MatSelectModule,
    ReactiveFormsModule,
  ],
})
export class OrderDialog implements OnInit {
  data = inject<IOrderDialogData>(MAT_DIALOG_DATA);
  private dialogRef = inject(MatDialogRef<OrderDialog>);
  private tradingService = inject(TradingService);
  private snackBar = inject(MatSnackBar);

  preview = signal({
    ltp: this.data.ltp ?? 0,
    margin_required: 0,
    available: 0,
    order_value: 0,
  });
  submitting = signal(false);

  form = new FormGroup({
    product_type: new FormControl<'regular' | 'intraday'>('regular', { nonNullable: true }),
    order_type: new FormControl<'limit' | 'market'>('limit', { nonNullable: true }),
    quantity: new FormControl(1, { nonNullable: true, validators: [Validators.required, Validators.min(1)] }),
    price: new FormControl(this.data.ltp ?? 0, { nonNullable: true, validators: [Validators.required, Validators.min(0.01)] }),
    stop_loss: new FormControl<number | null>(null),
  });

  ngOnInit(): void {
    this.form.valueChanges.pipe(debounceTime(200)).subscribe(() => this.refreshPreview());
    this.refreshPreview();
  }

  get title(): string {
    return `${this.data.side === 'buy' ? 'Buy' : 'Sell'} ${this.data.symbol}`;
  }

  async refreshPreview(): Promise<void> {
    const v = this.form.getRawValue();
    try {
      const preview = await this.tradingService.previewOrder({
        symbol: this.data.symbol,
        exchange: this.data.exchange,
        side: this.data.side,
        product_type: v.product_type,
        order_type: v.order_type,
        quantity: v.quantity,
        price: v.order_type === 'market' ? 0 : v.price,
      });
      this.preview.set({
        ltp: preview.ltp,
        margin_required: preview.margin_required,
        available: preview.available,
        order_value: preview.order_value,
      });
      if (v.order_type === 'market') {
        this.form.controls.price.setValue(preview.ltp, { emitEvent: false });
      }
    } catch {
      /* preview optional while typing */
    }
  }

  async submit(): Promise<void> {
    if (this.form.invalid) return;
    this.submitting.set(true);
    const v = this.form.getRawValue();
    try {
      await this.tradingService.placeOrder({
        symbol: this.data.symbol,
        exchange: this.data.exchange,
        side: this.data.side,
        product_type: v.product_type,
        order_type: v.order_type,
        quantity: v.quantity,
        price: v.order_type === 'market' ? this.preview().ltp : v.price,
        stop_loss: v.stop_loss ?? undefined,
      });
      this.snackBar.open('Order placed', 'Close', { duration: 2500 });
      this.dialogRef.close(true);
    } catch {
      this.submitting.set(false);
    }
  }
}
