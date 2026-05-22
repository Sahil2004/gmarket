import { inject, Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';
import { ChartRange } from '../types/stocks.types';
import { IAlgoConfig, IAlgoIndicators } from '../types/algo.types';

@Injectable({ providedIn: 'root' })
export class AlgoService {
  private http = inject(HttpClient);

  getConfig(symbol: string, exchange: string): Promise<IAlgoConfig> {
    const params = new HttpParams().set('symbol', symbol).set('exchange', exchange);
    return firstValueFrom(this.http.get<IAlgoConfig>('/algo/config', { params }));
  }

  saveConfig(config: IAlgoConfig): Promise<IAlgoConfig> {
    return firstValueFrom(this.http.put<IAlgoConfig>('/algo/config', config));
  }

  getIndicators(symbol: string, exchange: string, range: ChartRange): Promise<IAlgoIndicators> {
    const params = new HttpParams()
      .set('symbol', symbol)
      .set('exchange', exchange)
      .set('range', range);
    return firstValueFrom(this.http.get<IAlgoIndicators>('/algo/indicators', { params }));
  }
}
