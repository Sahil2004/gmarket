import { ResolveFn } from '@angular/router';
import { IStock } from '../types';
import { inject } from '@angular/core';
import { StocksService } from '../services';

export const stockResolver: ResolveFn<IStock[]> = () => {
  const stockService = inject(StocksService);

  return stockService.getAllStocks();
};
