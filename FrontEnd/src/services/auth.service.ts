import { api } from "../lib/axios";
import { API_PATH } from "../lib/api-path";
import { clearAccessToken } from "../lib/func/auth";
import { unwrapEnvelope } from "../lib/func/http";
import type {
  AuthResponse,
  LoginPayload,
  RegisterPayload,
} from "../lib/types";

export const authService = {
  async login(payload: LoginPayload): Promise<AuthResponse> {
    const response = await api.post(API_PATH.auth.login, payload);
    return unwrapEnvelope(response);
  },

  async register(payload: RegisterPayload): Promise<AuthResponse> {
    const response = await api.post(API_PATH.auth.register, payload);
    return unwrapEnvelope(response);
  },

  async refresh(): Promise<AuthResponse> {
    const response = await api.post(API_PATH.auth.refresh);
    return unwrapEnvelope(response);
  },

  async logout(): Promise<AuthResponse> {
    const response = await api.post(API_PATH.auth.logout);
    clearAccessToken();
    return unwrapEnvelope(response);
  },
};
