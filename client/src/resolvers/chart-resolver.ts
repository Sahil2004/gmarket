import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { StocksService } from '../services';
import { IChartApiResponse } from '../types/stocks.types';

export const chartResolver: ResolveFn<IChartApiResponse> = (
  route: ActivatedRouteSnapshot,
  state: RouterStateSnapshot,
) => {
  let stockSymbol = route.paramMap.get('symbol')!;
  let exchange = route.paramMap.get('exchange')!;
  let stockService = inject(StocksService);
  return stockService.getChartData(stockSymbol, exchange, '1d', '5d');
};
