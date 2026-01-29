import { Routes } from '@angular/router';
import { authGuard } from '../guards';
import { chartResolver, stockResolver, userResolver } from '../resolvers';

export const routes: Routes = [
  {
    path: '',
    loadComponent: () => import('../routes').then((m) => m.Dashboard),
    canActivate: [authGuard],
    canActivateChild: [authGuard],
    children: [
      {
        path: 'watchlist',
        loadComponent: () => import('../routes').then((m) => m.Watchlist),
        resolve: {
          stockData: stockResolver,
        },
        children: [
          {
            path: ':exchange/:symbol',
            loadComponent: () => import('../routes').then((m) => m.StockChart),
            resolve: {
              chartData: chartResolver,
            },
            runGuardsAndResolvers: 'paramsOrQueryParamsChange',
          },
        ],
      },
      {
        path: 'profile',
        loadComponent: () => import('../routes').then((m) => m.Profile),
        resolve: {
          userData: userResolver,
        },
        runGuardsAndResolvers: 'always',
      },
    ],
  },
  {
    path: 'login',
    loadComponent: () => import('../routes').then((m) => m.Login),
    canActivate: [authGuard],
  },
  {
    path: 'register',
    loadComponent: () => import('../routes').then((m) => m.Register),
    canActivate: [authGuard],
  },
];
