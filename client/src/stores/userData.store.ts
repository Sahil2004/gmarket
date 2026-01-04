import { Injectable } from '@angular/core';
import type { IUserData } from '../types/user-data.types';
import type { IUserDTO, IUserResponseDTO } from '../dtos/user.dto';

@Injectable({ providedIn: 'root' })
export class UserDataStore {
  private _userData: IUserData[];

  constructor() {
    this._userData = [];
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
      createdAt: new Date(),
      updatedAt: new Date(),
    };
    this._userData.push(newUser);
    return newUser;
  }
}
