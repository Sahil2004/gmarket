import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { UserService } from '../../services/user.service';
import { Router } from '@angular/router';
import { firstValueFrom } from 'rxjs';

export const authGuard: CanActivateFn = async (route, state) => {
  const userService = inject(UserService);
  const router = inject(Router);
  const authRoutes = ['login', 'register'];
  const currentUrlInAuthRoutes = authRoutes.some((r) => route.url.toString().includes(r));
  if (currentUrlInAuthRoutes) {
    try {
      const user = await firstValueFrom(userService.isAuthenticated());
      router.navigate(['/']);
      return false;
    } catch (err) {
      return true;
    }
  }
  try {
    const user = await firstValueFrom(userService.isAuthenticated());
    return true;
  } catch (err) {
    router.navigate(['/login']);
    return false;
  }
};
