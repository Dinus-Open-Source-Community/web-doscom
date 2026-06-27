import type { ApiEnvelope } from "./common";

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

/** Auth endpoints return envelope; tokens are set via HttpOnly cookies. */
export type AuthResponse = ApiEnvelope<null>;

export type LoginResponse = AuthResponse;
export type RefreshTokenResponse = AuthResponse;
