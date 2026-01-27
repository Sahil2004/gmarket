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
    userService.currentUser.subscribe((res) => {
      if (!res.hasOwnProperty('id')) {
        router.navigate(['/']);
      }
    });
    return true;
  }
  userService.currentUser.subscribe((res) => {
    if (!res.hasOwnProperty('id')) {
      router.navigate(['/login']);
    }
  });
  return true;
};
