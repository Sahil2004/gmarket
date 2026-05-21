import { Component, inject, signal } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSnackBar } from '@angular/material/snack-bar';
import { TradingService } from '../../../services';

@Component({
  selector: 'bank-account-dialog',
  templateUrl: 'bank-account-dialog.html',
  imports: [MatDialogModule, MatButtonModule, MatFormFieldModule, MatInputModule, ReactiveFormsModule],
})
export class BankAccountDialog {
  private dialogRef = inject(MatDialogRef<BankAccountDialog>);
  private tradingService = inject(TradingService);
  private snackBar = inject(MatSnackBar);
  submitting = signal(false);

  form = new FormGroup({
    bank_name: new FormControl('', { nonNullable: true, validators: [Validators.required] }),
    account_number: new FormControl('', { nonNullable: true, validators: [Validators.required] }),
    ifsc: new FormControl('', { nonNullable: true, validators: [Validators.required] }),
    nickname: new FormControl('', { nonNullable: true }),
  });

  async submit(): Promise<void> {
    if (this.form.invalid) return;
    this.submitting.set(true);
    try {
      await this.tradingService.addBankAccount(this.form.getRawValue());
      this.snackBar.open('Bank account added', 'Close', { duration: 2500 });
      this.dialogRef.close(true);
    } finally {
      this.submitting.set(false);
    }
  }
}
