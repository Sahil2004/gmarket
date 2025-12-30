import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: 'login',
    loadComponent: () => import('../routes/login').then((m) => m.Login),
  },
  {
    path: 'register',
    loadComponent: () => import('../routes/register').then((m) => m.Register),
  },
];
