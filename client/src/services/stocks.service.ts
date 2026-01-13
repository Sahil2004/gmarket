import { Injectable } from '@angular/core';
import { IStock } from '../types/stocks.types';

@Injectable({
  providedIn: 'root',
})
export class StocksService {
  private stocksData: IStock[] = [];

  constructor() {}
}
