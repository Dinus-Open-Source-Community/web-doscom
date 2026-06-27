import { isAxiosError } from "axios";
import {
  DEFAULT_ERROR_MESSAGE,
  ERROR_MESSAGES,
  ERROR_PATTERNS,
  HTTP_STATUS_MESSAGES,
  NETWORK_ERROR_MESSAGE,
  type ApiMessageEnvelope,
} from "../message";

export class ApiError extends Error {
  readonly status?: number;
  readonly rawMessage?: string;

  constructor(message: string, status?: number, rawMessage?: string) {
    super(message);
    this.name = "ApiError";
    this.status = status;
    this.rawMessage = rawMessage;
  }
}

export function normalizeMessageKey(value: string): string {
  return value.trim().toLowerCase().replace(/\s+/g, " ");
}

export function normalizeErrorKey(value: string): string {
  return normalizeMessageKey(value);
}

export function translateErrorMessage(raw: string): string {
  const normalized = normalizeMessageKey(raw);
  if (!normalized) return DEFAULT_ERROR_MESSAGE;

  const exact = ERROR_MESSAGES[normalized];
  if (exact) return exact;

  for (const { pattern, message } of ERROR_PATTERNS) {
    if (pattern.test(raw)) return message;
  }

  for (const [key, message] of Object.entries(ERROR_MESSAGES)) {
    if (normalized.includes(key) || key.includes(normalized)) {
      return message;
    }
  }

  return raw.trim();
}

function extractRawMessage(data: ApiMessageEnvelope): string | undefined {
  if (typeof data.message === "string" && data.message.trim()) {
    return data.message;
  }

  if (typeof data.error === "string" && data.error.trim()) {
    return data.error;
  }

  if (typeof data.errors === "string" && data.errors.trim()) {
    return data.errors;
  }

  if (data.error && typeof data.error === "object") {
    const nested = data.error as Record<string, unknown>;
    if (typeof nested.message === "string" && nested.message.trim()) {
      return nested.message;
    }
    if (typeof nested.error === "string" && nested.error.trim()) {
      return nested.error;
    }
  }

  return undefined;
}

export function parseApiError(error: unknown): string {
  if (isAxiosError(error)) {
    if (!error.response) {
      return NETWORK_ERROR_MESSAGE;
    }

    const data = error.response.data as ApiMessageEnvelope | string | undefined;
    let raw: string | undefined;

    if (typeof data === "string" && data.trim()) {
      raw = data;
    } else if (data && typeof data === "object") {
      raw = extractRawMessage(data);
    }

    if (raw) {
      return translateErrorMessage(raw);
    }

    const statusMessage = HTTP_STATUS_MESSAGES[error.response.status];
    if (statusMessage) return statusMessage;
  }

  if (error instanceof ApiError) {
    return error.message;
  }

  if (error instanceof Error && error.message.trim()) {
    return translateErrorMessage(error.message);
  }

  return DEFAULT_ERROR_MESSAGE;
}

export function toApiError(error: unknown): ApiError {
  if (error instanceof ApiError) return error;

  const message = parseApiError(error);
  const status = isAxiosError(error) ? error.response?.status : undefined;
  const raw =
    isAxiosError(error) && error.response?.data
      ? (extractRawMessage(error.response.data as ApiMessageEnvelope) ??
        JSON.stringify(error.response.data))
      : error instanceof Error
        ? error.message
        : undefined;

  return new ApiError(message, status, raw);
}
