import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { UserService } from '../../services/user.service';
import { Router } from '@angular/router';

export const authGuard: CanActivateFn = (route, state) => {
  const userService = inject(UserService);
  const router = inject(Router);
  if (!userService.isAuthenticated()) {
    const loginPath = '/login';
    router.navigate([loginPath]);
    return false;
  }
  return true;
};
