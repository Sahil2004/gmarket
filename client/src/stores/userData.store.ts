import { Inject, Injectable, PLATFORM_ID } from '@angular/core';
import type { IUserData } from '../types/user-data.types';
import type { IUserDTO, IUserResponseDTO } from '../dtos/user.dto';
import { isPlatformBrowser } from '@angular/common';

@Injectable({ providedIn: 'root' })
export class UserDataStore {
  private _isBrowser: boolean;
  private _userData: IUserData[];

  constructor(@Inject(PLATFORM_ID) private platformId: Object) {
    this._isBrowser = isPlatformBrowser(this.platformId);
    if (this._isBrowser)
      this._userData = JSON.parse(sessionStorage.getItem('userDataStore') || '[]');
    else this._userData = [];
  }

  saveToSession() {
    if (!this._isBrowser) return;
    sessionStorage.setItem('userDataStore', JSON.stringify(this._userData));
  }

  findUserById(id: string): IUserData | null {
    return this._userData.find((user) => user.id === id) ?? null;
  }

  findUserByEmail(email: string): IUserData | null {
    console.log(this._userData);
    return this._userData.find((user) => user.email === email) ?? null;
  }

  setUser(user: IUserDTO): IUserResponseDTO {
    const existingUserIndex = this._userData.findIndex((u) => u.email === user.email);
    if (existingUserIndex !== -1) return this._userData[existingUserIndex];
    const newUser: IUserData = {
      id: (this._userData.length + 1).toString(),
      name: user.name,
      email: user.email,
      password: user.password,
      createdAt: new Date(),
      updatedAt: new Date(),
    };
    this._userData.push(newUser);
    this.saveToSession();
    return newUser;
  }

  updateUser(id: string, updatedData: Partial<IUserDTO>): IUserResponseDTO | null {
    const _user = this._userData.findIndex((u) => u.id === id);
    if (_user === -1) return null;
    this._userData[_user] = {
      ...this._userData[_user],
      ...updatedData,
      updatedAt: new Date(),
    };
    this.saveToSession();
    return this._userData[_user];
  }
}
