export interface LoginPayload {
  email: string;
  password: string;
}

export interface RegisterPayload {
  username: string;
  email: string;
  password: string;
  role: string;
  fullname: string;
}

export interface LoginResponse {
  acces_token?: string;
  access_token?: string;
  message?: string;
  "message:"?: string;
}

export interface RefreshTokenResponse {
  access_token: string;
  message?: string;
}
