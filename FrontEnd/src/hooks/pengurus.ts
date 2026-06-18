import {
  useMutation,
  useQuery,
  useQueryClient,
  type UseMutationOptions,
  type UseQueryOptions,
} from "@tanstack/react-query";
import {
  pengurusService,
  type CreatePengurusPayload,
  type UpdatePengurusPayload,
} from "../services/pengurus.service";
import type { Pengurus, PengurusPublic } from "../lib/types";
import { pengurusKeys } from "./keys";

export function usePengurusByDivisionQuery(
  division: string,
  options?: Omit<UseQueryOptions<PengurusPublic[]>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: pengurusKeys.byDivision(division),
    queryFn: () => pengurusService.listByDivision(division),
    enabled: Boolean(division),
    ...options,
  });
}

export function usePengurusProfileQuery(
  options?: Omit<UseQueryOptions<Pengurus>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: pengurusKeys.profile(),
    queryFn: () => pengurusService.getProfile(),
    ...options,
  });
}

export function useAdminPengurusListQuery(
  divisi?: string,
  options?: Omit<UseQueryOptions<Pengurus[]>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: pengurusKeys.admin.list(divisi),
    queryFn: () => pengurusService.admin.list(divisi),
    ...options,
  });
}

export function useAdminPengurusQuery(
  id: number | string,
  options?: Omit<UseQueryOptions<Pengurus>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: pengurusKeys.admin.detail(id),
    queryFn: () => pengurusService.admin.getById(id),
    enabled: id !== undefined && id !== "",
    ...options,
  });
}

export function useAdminPengurusByUserQuery(
  userId: number | string,
  options?: Omit<UseQueryOptions<Pengurus>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: pengurusKeys.admin.byUser(userId),
    queryFn: () => pengurusService.admin.getByUserId(userId),
    enabled: userId !== undefined && userId !== "",
    ...options,
  });
}

export function useCreatePengurusProfileMutation(
  options?: UseMutationOptions<
    Pengurus,
    Error,
    { payload: CreatePengurusPayload; file?: File }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ payload, file }) =>
      pengurusService.createProfile(payload, file),
    onSuccess: (data, ...rest) => {
      queryClient.setQueryData(pengurusKeys.profile(), data);
      queryClient.invalidateQueries({ queryKey: pengurusKeys.all });
      options?.onSuccess?.(data, ...rest);
    },
    ...options,
  });
}

export function useUpdatePengurusMeMutation(
  options?: UseMutationOptions<
    Pengurus,
    Error,
    { payload: UpdatePengurusPayload; file?: File }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ payload, file }) =>
      pengurusService.updateMe(payload, file),
    onSuccess: (data, ...rest) => {
      queryClient.setQueryData(pengurusKeys.profile(), data);
      queryClient.invalidateQueries({ queryKey: pengurusKeys.all });
      options?.onSuccess?.(data, ...rest);
    },
    ...options,
  });
}

export function useDeletePengurusMeMutation(
  options?: UseMutationOptions<null, Error, void>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: () => pengurusService.deleteMe(),
    onSuccess: (...args) => {
      queryClient.removeQueries({ queryKey: pengurusKeys.profile() });
      queryClient.invalidateQueries({ queryKey: pengurusKeys.all });
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}

export function useCreateAdminPengurusMutation(
  options?: UseMutationOptions<
    Pengurus,
    Error,
    { payload: CreatePengurusPayload; file?: File }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ payload, file }) =>
      pengurusService.admin.create(payload, file),
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: pengurusKeys.admin.lists() });
      queryClient.invalidateQueries({ queryKey: pengurusKeys.all });
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}

export function useUpdateAdminPengurusMutation(
  options?: UseMutationOptions<
    Pengurus,
    Error,
    { id: number | string; payload: UpdatePengurusPayload; file?: File }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, payload, file }) =>
      pengurusService.admin.update(id, payload, file),
    onSuccess: (data, variables, ...rest) => {
      queryClient.setQueryData(
        pengurusKeys.admin.detail(variables.id),
        data,
      );
      queryClient.invalidateQueries({ queryKey: pengurusKeys.admin.lists() });
      queryClient.invalidateQueries({ queryKey: pengurusKeys.all });
      options?.onSuccess?.(data, variables, ...rest);
    },
    ...options,
  });
}

export function useDeleteAdminPengurusMutation(
  options?: UseMutationOptions<null, Error, number | string>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number | string) => pengurusService.admin.remove(id),
    onSuccess: (data, id, ...rest) => {
      queryClient.removeQueries({ queryKey: pengurusKeys.admin.detail(id) });
      queryClient.invalidateQueries({ queryKey: pengurusKeys.admin.lists() });
      queryClient.invalidateQueries({ queryKey: pengurusKeys.all });
      options?.onSuccess?.(data, id, ...rest);
    },
    ...options,
  });
}
