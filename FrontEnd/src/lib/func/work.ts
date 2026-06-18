import { toFormData } from "./http";

export function buildWorkFormData<T extends { technologies?: string[] }>(
  payload: T,
  files?: File[],
): FormData {
  const { technologies, ...rest } = payload;

  return toFormData(
    technologies ? { ...rest, "technologies[]": technologies } : rest,
    files ? { files } : undefined,
  );
}
