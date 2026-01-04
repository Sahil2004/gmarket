export interface IUserDTO {
  name: string;
  email: string;
  password: string;
}

export interface IUserResponseDTO {
  id: string;
  name: string;
  email: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface IUserLoginDTO {
  email: string;
  password: string;
}
