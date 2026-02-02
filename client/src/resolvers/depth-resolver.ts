import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { StocksService } from '../services';
import { IMarketDepth } from '../types';

export const depthResolver: ResolveFn<IMarketDepth> = (
  route: ActivatedRouteSnapshot,
  state: RouterStateSnapshot,
) => {
  let stockSymbol = route.paramMap.get('symbol')!;
  let exchange = route.paramMap.get('exchange')!;
  let stockService = inject(StocksService);
  return stockService.getDepthData(stockSymbol, exchange);
};
