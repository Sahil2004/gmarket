import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { UserService } from '../../services/user.service';
import { Router } from '@angular/router';
import { Location } from '@angular/common';

export const authGuard: CanActivateFn = (route, state) => {
  const userService = inject(UserService);
  const router = inject(Router);
  const location = inject(Location);
  const authRoutes = ['login', 'register'];
  const currentUrlInAuthRoutes = authRoutes.some((r) => route.url.toString().includes(r));
  if (!userService.isAuthenticated() && !currentUrlInAuthRoutes) {
    const loginPath = '/login';
    router.navigate([loginPath]);
    return false;
  }
  if (currentUrlInAuthRoutes) {
    location.back();
  }
  return true;
};
