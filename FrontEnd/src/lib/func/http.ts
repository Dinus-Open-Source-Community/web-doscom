import type { AxiosRequestConfig, AxiosResponse } from "axios";
import { api } from "../axios";
import type { ApiEnvelope } from "../types";

export function unwrapEnvelopeData<T>(
  response: AxiosResponse<ApiEnvelope<T>>,
): T {
  return response.data.data as T;
}

export function unwrapEnvelope<T>(
  response: AxiosResponse<ApiEnvelope<T>>,
): ApiEnvelope<T> {
  return response.data;
}

export async function getEnvelopeData<T>(
  url: string,
  config?: AxiosRequestConfig,
): Promise<T> {
  const response = await api.get<ApiEnvelope<T>>(url, config);
  return unwrapEnvelopeData(response);
}

export async function getEnvelope<T>(
  url: string,
  config?: AxiosRequestConfig,
): Promise<ApiEnvelope<T>> {
  const response = await api.get<ApiEnvelope<T>>(url, config);
  return unwrapEnvelope(response);
}

export async function postEnvelopeData<T, B = unknown>(
  url: string,
  body?: B,
  config?: AxiosRequestConfig,
): Promise<T> {
  const response = await api.post<ApiEnvelope<T>>(url, body, config);
  return unwrapEnvelopeData(response);
}

export async function putEnvelopeData<T, B = unknown>(
  url: string,
  body?: B,
  config?: AxiosRequestConfig,
): Promise<T> {
  const response = await api.put<ApiEnvelope<T>>(url, body, config);
  return unwrapEnvelopeData(response);
}

export async function deleteEnvelope(
  url: string,
  config?: AxiosRequestConfig,
): Promise<ApiEnvelope<null>> {
  const response = await api.delete<ApiEnvelope<null>>(url, config);
  return unwrapEnvelope(response);
}

export function toFormData(
  payload: object,
  files?: Record<string, File | File[] | undefined>,
): FormData {
  const formData = new FormData();

  Object.entries(payload as Record<string, unknown>).forEach(([key, value]) => {
    if (value === undefined || value === null) return;

    if (Array.isArray(value)) {
      value.forEach((item) => formData.append(key, String(item)));
      return;
    }

    formData.append(key, String(value));
  });

  if (files) {
    Object.entries(files).forEach(([key, value]) => {
      if (!value) return;

      if (Array.isArray(value)) {
        value.forEach((file) => formData.append(key, file));
        return;
      }

      formData.append(key, value);
    });
  }

  return formData;
}
