import {
  useMutation,
  useQuery,
  useQueryClient,
  type UseMutationOptions,
  type UseQueryOptions,
} from "@tanstack/react-query";
import {
  galleryService,
  type CreateGalleryPayload,
  type GalleryQuery,
} from "../services/gallery.service";
import type { GalleryItem, GalleryListResponse } from "../lib/types";
import { galleryKeys } from "./keys";

export function useGalleryQuery(
  params?: GalleryQuery,
  options?: Omit<
    UseQueryOptions<GalleryListResponse>,
    "queryKey" | "queryFn"
  >,
) {
  return useQuery({
    queryKey: galleryKeys.list(params),
    queryFn: () => galleryService.list(params),
    ...options,
  });
}

export function useCreateGalleryMutation(
  options?: UseMutationOptions<
    { message: string; data: GalleryItem[] },
    Error,
    { payload: CreateGalleryPayload; files: File[] }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ payload, files }) =>
      galleryService.admin.create(payload, files),
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: galleryKeys.lists() });
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}

export function useDeleteGalleryMutation(
  options?: UseMutationOptions<{ message: string }, Error, number | string>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number | string) => galleryService.admin.remove(id),
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: galleryKeys.lists() });
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}
