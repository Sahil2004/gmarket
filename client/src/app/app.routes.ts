import { Routes } from '@angular/router';

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
  },
];
