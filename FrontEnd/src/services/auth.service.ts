import { api } from "../lib/axios";
import { API_PATH } from "../lib/api-path";
import {
  clearAccessToken,
  extractAccessToken,
  setAccessToken,
} from "../lib/func/auth";
import type {
  LoginPayload,
  LoginResponse,
  MessageResponse,
  RefreshTokenResponse,
  RegisterPayload,
} from "../lib/types";

export const authService = {
  async login(payload: LoginPayload): Promise<LoginResponse> {
    const { data } = await api.post<LoginResponse>(API_PATH.auth.login, payload);
    const token = extractAccessToken(data);
    setAccessToken(token);
    return data;
  },

  async register(payload: RegisterPayload): Promise<MessageResponse> {
    const { data } = await api.post<MessageResponse>(
      API_PATH.auth.register,
      payload,
    );
    return data;
  },

  async refresh(): Promise<RefreshTokenResponse> {
    const { data } = await api.post<RefreshTokenResponse>(API_PATH.auth.refresh);
    const token = extractAccessToken(data);
    setAccessToken(token);
    return data;
  },

  async logout(): Promise<MessageResponse> {
    const { data } = await api.post<MessageResponse>(API_PATH.auth.logout);
    clearAccessToken();
    return data;
  },
};
