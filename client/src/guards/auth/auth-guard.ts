import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { UserService } from '../../services/user.service';
import { Router } from '@angular/router';

export const authGuard: CanActivateFn = (route, state) => {
  const userService = inject(UserService);
  const router = inject(Router);
  const authRoutes = ['login', 'register'];
  const currentUrlInAuthRoutes = authRoutes.some((r) => route.url.toString().includes(r));
  if (currentUrlInAuthRoutes) {
    if (userService.isAuthenticated()) {
      router.navigate(['/']);
      return false;
    }
    return true;
  }
  if (!userService.isAuthenticated()) {
    router.navigate(['/login']);
    return false;
  }
  return true;
};
