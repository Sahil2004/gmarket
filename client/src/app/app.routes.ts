import { Routes } from '@angular/router';
import { authGuard } from '../guards';

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
      },
      {
        path: 'profile',
        loadComponent: () => import('../routes').then((m) => m.Profile),
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
