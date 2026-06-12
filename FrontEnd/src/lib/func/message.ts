import {
  DEFAULT_SUCCESS_MESSAGE,
  SUCCESS_MESSAGES,
  SUCCESS_PATTERNS,
  type ApiMessageEnvelope,
} from "../message";
import { normalizeMessageKey, translateErrorMessage } from "./error";

export { normalizeMessageKey };

export function translateSuccessMessage(raw: string): string {
  const normalized = normalizeMessageKey(raw);
  if (!normalized) return DEFAULT_SUCCESS_MESSAGE;

  const exact = SUCCESS_MESSAGES[normalized];
  if (exact) return exact;

  for (const { pattern, message } of SUCCESS_PATTERNS) {
    if (pattern.test(raw)) return message;
  }

  for (const [key, message] of Object.entries(SUCCESS_MESSAGES)) {
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

export function parseApiMessage(data: unknown, success = false): string {
  if (typeof data === "string" && data.trim()) {
    return success ? translateSuccessMessage(data) : translateErrorMessage(data);
  }

  if (data && typeof data === "object") {
    const raw = extractRawMessage(data as ApiMessageEnvelope);
    if (raw) {
      const isSuccess =
        success ||
        ("success" in data && (data as ApiMessageEnvelope).success === true);
      return isSuccess ? translateSuccessMessage(raw) : translateErrorMessage(raw);
    }
  }

  return success ? DEFAULT_SUCCESS_MESSAGE : translateErrorMessage("");
}

export function translateApiMessage(raw: string, success = false): string {
  return success ? translateSuccessMessage(raw) : translateErrorMessage(raw);
}
