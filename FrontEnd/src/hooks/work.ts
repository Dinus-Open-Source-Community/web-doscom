import {
  useMutation,
  useQuery,
  useQueryClient,
  type UseMutationOptions,
  type UseQueryOptions,
} from "@tanstack/react-query";
import {
  workService,
  type CreateWorkPayload,
  type PublicWorkQuery,
  type UpdateWorkPayload,
} from "../services/work.service";
import type {
  PaginationQuery,
  WorkInternal,
  WorkListResponse,
  WorkPublic,
} from "../lib/types";
import { workKeys } from "./keys";

export function useWorksByProjectTypeQuery(
  projectType: string,
  params?: PaginationQuery,
  options?: Omit<UseQueryOptions<WorkListResponse>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: workKeys.byProjectType(projectType, params),
    queryFn: () => workService.listByProjectType(projectType, params),
    enabled: Boolean(projectType),
    ...options,
  });
}

export function useAdminWorksQuery(
  params?: PublicWorkQuery,
  options?: Omit<UseQueryOptions<WorkListResponse>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: workKeys.admin.list(params),
    queryFn: () => workService.admin.list(params),
    ...options,
  });
}

export function useAdminWorkQuery(
  id: number | string,
  options?: Omit<UseQueryOptions<WorkInternal>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: workKeys.admin.detail(id),
    queryFn: () => workService.admin.getById(id),
    enabled: id !== undefined && id !== "",
    ...options,
  });
}

export function useCreateWorkMutation(
  options?: UseMutationOptions<
    WorkPublic,
    Error,
    { payload: CreateWorkPayload; files?: File[] }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ payload, files }) =>
      workService.admin.create(payload, files),
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: workKeys.admin.lists() });
      queryClient.invalidateQueries({ queryKey: workKeys.all });
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}

export function useUpdateWorkMutation(
  options?: UseMutationOptions<
    WorkInternal,
    Error,
    { id: number | string; payload: UpdateWorkPayload; files?: File[] }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, payload, files }) =>
      workService.admin.update(id, payload, files),
    onSuccess: (data, variables, ...rest) => {
      queryClient.setQueryData(workKeys.admin.detail(variables.id), data);
      queryClient.invalidateQueries({ queryKey: workKeys.admin.lists() });
      queryClient.invalidateQueries({ queryKey: workKeys.all });
      options?.onSuccess?.(data, variables, ...rest);
    },
    ...options,
  });
}

export function useDeleteWorkMutation(
  options?: UseMutationOptions<null, Error, number | string>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number | string) => workService.admin.remove(id),
    onSuccess: (data, id, ...rest) => {
      queryClient.removeQueries({ queryKey: workKeys.admin.detail(id) });
      queryClient.invalidateQueries({ queryKey: workKeys.admin.lists() });
      queryClient.invalidateQueries({ queryKey: workKeys.all });
      options?.onSuccess?.(data, id, ...rest);
    },
    ...options,
  });
}
