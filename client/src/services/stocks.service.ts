import { inject, Injectable } from '@angular/core';
import { IStock } from '../types/stocks.types';
import { HttpClient } from '@angular/common/http';
import type { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class StocksService {
  private http = inject(HttpClient);

  constructor() {}

  getAllStocks(): Observable<IStock[]> {
    return this.http.get<IStock[]>('/data/stocks.json');
  }
}
