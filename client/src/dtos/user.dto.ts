export interface IUserDTO {
  name: string;
  email: string;
  password: string;
  phoneNumber?: number;
  profilePhotoUri?: string;
}

export interface IWatchlistDTO {
  id: string;
  stocks?: string[];
}

export interface IUserResponseDTO {
  id: string;
  name: string;
  email: string;
  phoneNumber?: number;
  profilePhotoUri?: string;
  watchlists: string[][];
  createdAt: Date;
  updatedAt: Date;
}

export interface IUserLoginDTO {
  email: string;
  password: string;
}
