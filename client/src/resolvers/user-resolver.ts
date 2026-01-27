import { ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import type { IError, IUserData } from '../types';
import { inject } from '@angular/core';
import { UserService } from '../services';

export const userResolver: ResolveFn<IUserData | IError> = (
  route: ActivatedRouteSnapshot,
  state: RouterStateSnapshot,
) => {
  const userService = inject(UserService);

  return userService.currentUser;
};
