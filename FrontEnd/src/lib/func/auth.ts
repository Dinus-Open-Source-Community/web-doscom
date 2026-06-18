import type { LoginResponse, RefreshTokenResponse } from "../types";

const ACCESS_TOKEN_KEY = "access_token";

export function getAccessToken(): string | null {
  if (typeof window === "undefined") return null;
  return localStorage.getItem(ACCESS_TOKEN_KEY);
}

export function setAccessToken(token: string): void {
  localStorage.setItem(ACCESS_TOKEN_KEY, token);
}

export function clearAccessToken(): void {
  localStorage.removeItem(ACCESS_TOKEN_KEY);
}

export function extractAccessToken(
  data: LoginResponse | RefreshTokenResponse,
): string {
  const token =
    "acces_token" in data ? data.acces_token : data.access_token;

  if (!token) {
    throw new Error("Token autentikasi tidak ditemukan pada respons server.");
  }

  return token;
}
