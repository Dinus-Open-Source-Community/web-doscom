import {
  useMutation,
  useQuery,
  useQueryClient,
  type UseMutationOptions,
  type UseQueryOptions,
} from "@tanstack/react-query";
import {
  blogService,
  type AdminBlogQuery,
  type CreateBlogPayload,
  type PublicBlogQuery,
  type UpdateBlogPayload,
} from "../services/blog.service";
import type { Blog, BlogListResponse } from "../lib/types";
import { blogKeys } from "./keys";

export function useBlogsQuery(
  params?: PublicBlogQuery,
  options?: Omit<UseQueryOptions<BlogListResponse>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: blogKeys.list(params),
    queryFn: () => blogService.list(params),
    ...options,
  });
}

export function useBlogQuery(
  id: number | string,
  options?: Omit<UseQueryOptions<Blog>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: blogKeys.detail(id),
    queryFn: () => blogService.getById(id),
    enabled: id !== undefined && id !== "",
    ...options,
  });
}

export function useAdminBlogsQuery(
  params?: AdminBlogQuery,
  options?: Omit<UseQueryOptions<BlogListResponse>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: blogKeys.admin.list(params),
    queryFn: () => blogService.admin.list(params),
    ...options,
  });
}

export function useAdminBlogQuery(
  id: number | string,
  options?: Omit<UseQueryOptions<Blog>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: blogKeys.admin.detail(id),
    queryFn: () => blogService.admin.getById(id),
    enabled: id !== undefined && id !== "",
    ...options,
  });
}

export function useCreateBlogMutation(
  options?: UseMutationOptions<
    { message: string; data?: Blog },
    Error,
    { payload: CreateBlogPayload; files?: File[] }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ payload, files }) =>
      blogService.admin.create(payload, files),
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: blogKeys.admin.lists() });
      queryClient.invalidateQueries({ queryKey: blogKeys.lists() });
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}

export function useUpdateBlogMutation(
  options?: UseMutationOptions<
    { message: string },
    Error,
    { id: number | string; payload: UpdateBlogPayload; files?: File[] }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, payload, files }) =>
      blogService.admin.update(id, payload, files),
    onSuccess: (data, variables, ...rest) => {
      queryClient.invalidateQueries({ queryKey: blogKeys.admin.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: blogKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: blogKeys.admin.lists() });
      queryClient.invalidateQueries({ queryKey: blogKeys.lists() });
      options?.onSuccess?.(data, variables, ...rest);
    },
    ...options,
  });
}

export function useDeleteBlogMutation(
  options?: UseMutationOptions<
    { message: string },
    Error,
    number | string
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number | string) => blogService.admin.remove(id),
    onSuccess: (data, id, ...rest) => {
      queryClient.removeQueries({ queryKey: blogKeys.admin.detail(id) });
      queryClient.removeQueries({ queryKey: blogKeys.detail(id) });
      queryClient.invalidateQueries({ queryKey: blogKeys.admin.lists() });
      queryClient.invalidateQueries({ queryKey: blogKeys.lists() });
      options?.onSuccess?.(data, id, ...rest);
    },
    ...options,
  });
}
