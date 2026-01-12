export interface IUserData {
  id: string;
  name: string;
  email: string;
  phoneNumber?: number;
  profilePhotoUri?: string;
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
  createdAt: Date;
  updatedAt: Date;
}
