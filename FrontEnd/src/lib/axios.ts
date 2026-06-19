import axios from "axios";
import { getAccessToken } from "./func/auth";
import { toApiError } from "./func/error";

export const api = axios.create({
  baseURL: import.meta.env.SSR_API_URL ?? import.meta.env.PUBLIC_API_URL ?? "http://localhost:8080/api/v1",
  withCredentials: true,
});

api.interceptors.request.use((config) => {
  const token = getAccessToken();
  if (token) config.headers.Authorization = `Bearer ${token}`;
  return config;
});

api.interceptors.response.use(
  (response) => response,
  (error) => Promise.reject(toApiError(error)),
);
