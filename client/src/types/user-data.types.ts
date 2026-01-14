export interface IUserData {
  id: string;
  name: string;
  email: string;
  phoneNumber?: number;
  profilePhotoUri?: string;
  watchlists: string[][];
  password: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface IUserDataClient {
  id: string;
  name: string;
  email: string;
  phoneNumber?: number;
  profilePhotoUri?: string;
  watchlists: string[][];
  createdAt: Date;
  updatedAt: Date;
}
