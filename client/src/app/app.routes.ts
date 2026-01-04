import { Routes } from '@angular/router';
import { authGuard } from '../guards';

export const routes: Routes = [
  {
    path: 'login',
    loadComponent: () => import('../routes').then((m) => m.Login),
  },
  {
    path: 'register',
    loadComponent: () => import('../routes').then((m) => m.Register),
  },
  {
    path: 'dashboard',
    loadComponent: () => import('../routes').then((m) => m.Dashboard),
    canActivate: [authGuard],
  },
];
