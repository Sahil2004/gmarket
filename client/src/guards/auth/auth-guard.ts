import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { UserService } from '../../services/user.service';
import { Router } from '@angular/router';
import { Location } from '@angular/common';

export const authGuard: CanActivateFn = (route, state) => {
  const userService = inject(UserService);
  const router = inject(Router);
  const location = inject(Location);
  const authRoutes = ['/login', '/register'];
  if (!userService.isAuthenticated()) {
    const loginPath = '/login';
    router.navigate([loginPath]);
    return false;
  }
  if (authRoutes.some((r) => route.url.toString().includes(r))) {
    location.back();
  }
  return true;
};
