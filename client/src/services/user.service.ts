import { inject, Injectable } from '@angular/core';

import { IUserData, IUserUpdateData, IError } from '../types';
import { HttpClient, HttpContext } from '@angular/common/http';
import { Observable } from 'rxjs';
import { SKIP_TOAST } from '../contexts';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private http = inject(HttpClient);

  get currentUser(): Observable<IUserData | IError> {
    return this.http.get<IUserData | IError>('/users');
  }

  isAuthenticated(): Observable<IUserData | IError> {
    return this.http.get<IUserData | IError>('/users/is-authenticated', {
      context: new HttpContext().set(SKIP_TOAST, true),
    });
  }

  login(email: string, password: string): Observable<IUserData | IError> {
    const loginDetails = {
      email,
      password,
    };
    return this.http.post<IUserData | IError>('/sessions', loginDetails);
  }

  register(name: string, email: string, password: string): Observable<IUserData | IError> {
    const registerDetails = {
      email,
      name,
      password,
    };
    return this.http.post<IUserData | IError>('/users', registerDetails);
  }

  logout(): Observable<null | IError> {
    return this.http.delete<null | IError>('/sessions');
  }

  updateProfile(updatedData: Partial<IUserUpdateData>): Observable<IUserData | IError> {
    return this.http.patch<IUserData | IError>('/users', updatedData);
  }

  changePassword(oldPassword: string, newPassword: string): Observable<null | IError> {
    const changePasswordDetails = {
      old_password: oldPassword,
      new_password: newPassword,
    };
    return this.http.post<null | IError>('/users/change-password', changePasswordDetails);
  }

  deleteAccount(): Observable<null | IError> {
    return this.http.delete<null | IError>('/users');
  }
}
