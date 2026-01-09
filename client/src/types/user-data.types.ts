export interface IUserData {
  id: string;
  name: string;
  email: string;
  phoneNumber?: number;
  password: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface IUserDataClient {
  id: string;
  name: string;
  email: string;
  phoneNumber?: number;
  createdAt: Date;
  updatedAt: Date;
}
