import { inject, Injectable } from '@angular/core';

import { IUserDataClient } from '../types/user-data.types';
import { UserDataStore } from '../stores/userData.store';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private user: IUserDataClient | null = null;
  private ds = inject(UserDataStore);

  login(email: string, password: string): boolean {
    const res = this.ds.findUserByEmail(email);
    if (res && res.password === password) {
      this.user = res;
      return true;
    }
    return false;
  }
}
