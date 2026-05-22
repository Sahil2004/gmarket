import { inject, Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import {
  IBankAccount,
  IOrder,
  IOrderPreview,
  IPlaceOrder,
  ITradingAccount,
  ITradingSnapshot,
} from '../types/trading.types';
import { PollingService } from './polling.service';

@Injectable({ providedIn: 'root' })
export class TradingService {
  private http = inject(HttpClient);
  private polling = inject(PollingService);
  private cachedSnapshot: ITradingSnapshot | null = null;
  private snapshotListeners = new Set<(s: ITradingSnapshot) => void>();

  constructor() {
    this.polling.register('trading', () => this.pollSnapshot());
  }

  onSnapshotUpdate(listener: (s: ITradingSnapshot) => void): () => void {
    this.snapshotListeners.add(listener);
    return () => this.snapshotListeners.delete(listener);
  }

  startPolling(): void {
    this.polling.start('trading');
  }

  stopPolling(): void {
    this.polling.stop('trading');
  }

  private async pollSnapshot(): Promise<void> {
    try {
      this.cachedSnapshot = await firstValueFrom(
        this.http.get<ITradingSnapshot>('/trading/snapshot'),
      );
      for (const listener of this.snapshotListeners) {
        listener(this.cachedSnapshot);
      }
    } catch {
      /* best-effort */
    }
  }

  async getSnapshot(): Promise<ITradingSnapshot> {
    this.cachedSnapshot = await firstValueFrom(this.http.get<ITradingSnapshot>('/trading/snapshot'));
    return this.cachedSnapshot;
  }

  getCachedSnapshot(): ITradingSnapshot | null {
    return this.cachedSnapshot;
  }

  async getBankAccounts(): Promise<IBankAccount[]> {
    return firstValueFrom(this.http.get<IBankAccount[]>('/trading/banks'));
  }

  async previewOrder(params: Partial<IPlaceOrder>): Promise<IOrderPreview> {
    let httpParams = new HttpParams();
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        httpParams = httpParams.set(key, String(value));
      }
    });
    return firstValueFrom(
      this.http.get<IOrderPreview>('/trading/orders/preview', { params: httpParams }),
    );
  }

  async placeOrder(order: IPlaceOrder): Promise<IOrder> {
    const placed = await firstValueFrom(this.http.post<IOrder>('/trading/orders', order));
    await this.getSnapshot();
    return placed;
  }

  async cancelOrder(orderId: string): Promise<void> {
    await firstValueFrom(this.http.delete(`/trading/orders/${orderId}`));
    await this.getSnapshot();
  }

  async addBankAccount(body: {
    bank_name: string;
    account_number: string;
    ifsc: string;
    nickname?: string;
  }): Promise<IBankAccount> {
    const account = await firstValueFrom(this.http.post<IBankAccount>('/trading/banks', body));
    await this.getSnapshot();
    return account;
  }

  async removeBankAccount(bankId: string): Promise<void> {
    await firstValueFrom(this.http.delete(`/trading/banks/${bankId}`));
    await this.getSnapshot();
  }

  async depositFunds(bankAccountId: string, amount: number): Promise<ITradingAccount> {
    const account = await firstValueFrom(
      this.http.post<ITradingAccount>('/trading/funds/deposit', {
        bank_account_id: bankAccountId,
        amount,
      }),
    );
    await this.getSnapshot();
    return account;
  }

  async withdrawFunds(bankAccountId: string, amount: number): Promise<ITradingAccount> {
    const account = await firstValueFrom(
      this.http.post<ITradingAccount>('/trading/funds/withdraw', {
        bank_account_id: bankAccountId,
        amount,
      }),
    );
    await this.getSnapshot();
    return account;
  }
}
