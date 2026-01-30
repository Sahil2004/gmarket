import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { UserService } from '../../services/user.service';
import { Router } from '@angular/router';
import { firstValueFrom } from 'rxjs';

export const unauthGuard: CanActivateFn = async (route, state) => {
  const userService = inject(UserService);
  const router = inject(Router);
  try {
    const user = await firstValueFrom(userService.isAuthenticated());
    router.navigate(['/']);
    return false;
  } catch (err) {
    return true;
  }
};
