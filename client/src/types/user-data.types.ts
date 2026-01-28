export interface IUserData {
  id: string;
  name: string;
  email: string;
  phone_number?: number;
  profile_picture_url?: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface IUserUpdateData {
  email?: string;
  name?: string;
  phone_number?: number;
  profile_picture_url?: string;
}
