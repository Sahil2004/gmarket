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
      watchlists: Array.from({ length: 7 }, () => []),
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

  updatePassword(id: string, oldPassword: string, newPassword: string): IUserResponseDTO | null {
    const _user = this._userData.findIndex((u) => u.id === id);
    if (_user === -1) return null;
    if (this._userData[_user].password !== oldPassword) return null;
    this._userData[_user].password = newPassword;
    this._userData[_user].updatedAt = new Date();
    this.saveToSession();
    return this._userData[_user];
  }

  deleteUser(id: string): IUserResponseDTO | null {
    const _userIndex = this._userData.findIndex((u) => u.id === id);
    if (_userIndex === -1) return null;
    const deletedUser = this._userData.splice(_userIndex, 1)[0];
    this.saveToSession();
    return deletedUser;
  }

  addStockToWatchlist(userId: string, watchlistIndex: number, stockSymbol: string): boolean {
    const _userIdx = this._userData.findIndex((u) => u.id === userId);
    if (_userIdx === -1) return false;
    const watchlist = this._userData[_userIdx].watchlists[watchlistIndex];
    if (watchlist.includes(stockSymbol)) return false;
    watchlist.push(stockSymbol);
    this._userData[_userIdx].updatedAt = new Date();
    this.saveToSession();
    return true;
  }
}
