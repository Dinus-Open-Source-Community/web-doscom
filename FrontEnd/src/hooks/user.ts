import {
  useMutation,
  useQuery,
  useQueryClient,
  type UseMutationOptions,
  type UseQueryOptions,
} from "@tanstack/react-query";
import { userService } from "../services/user.service";
import type {
  AdminChangePasswordPayload,
  ChangePasswordPayload,
  PaginationQuery,
  RegisterPayload,
  User,
  UserPatch,
} from "../lib/types";
import { userKeys } from "./keys";

export function useMeQuery(
  options?: Omit<UseQueryOptions<User>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: userKeys.me(),
    queryFn: () => userService.getMe(),
    ...options,
  });
}

export function useUsersQuery(
  params?: PaginationQuery,
  options?: Omit<UseQueryOptions<User[]>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: userKeys.list(params),
    queryFn: () => userService.list(params),
    ...options,
  });
}

export function useUserQuery(
  id: number | string,
  options?: Omit<UseQueryOptions<User>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: userKeys.detail(id),
    queryFn: () => userService.getById(id),
    enabled: id !== undefined && id !== "",
    ...options,
  });
}

export function useUpdateProfileMutation(
  options?: UseMutationOptions<User, Error, UserPatch>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: UserPatch) => userService.updateProfile(payload),
    onSuccess: (data, ...rest) => {
      queryClient.setQueryData(userKeys.me(), data);
      options?.onSuccess?.(data, ...rest);
    },
    ...options,
  });
}

export function useChangePasswordMutation(
  options?: UseMutationOptions<null, Error, ChangePasswordPayload>,
) {
  return useMutation({
    mutationFn: (payload: ChangePasswordPayload) =>
      userService.changePassword(payload),
    ...options,
  });
}

export function useCreateUserMutation(
  options?: UseMutationOptions<null, Error, RegisterPayload>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: RegisterPayload) => userService.create(payload),
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}

export function useUpdateUserMutation(
  options?: UseMutationOptions<
    User,
    Error,
    { id: number | string; payload: UserPatch }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, payload }) => userService.update(id, payload),
    onSuccess: (data, variables, ...rest) => {
      queryClient.setQueryData(userKeys.detail(variables.id), data);
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
      options?.onSuccess?.(data, variables, ...rest);
    },
    ...options,
  });
}

export function useDeleteUserMutation(
  options?: UseMutationOptions<null, Error, number | string>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number | string) => userService.remove(id),
    onSuccess: (data, id, ...rest) => {
      queryClient.removeQueries({ queryKey: userKeys.detail(id) });
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
      options?.onSuccess?.(data, id, ...rest);
    },
    ...options,
  });
}

export function useSuperAdminsQuery(
  options?: Omit<UseQueryOptions<User[]>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: userKeys.superAdmins(),
    queryFn: () => userService.admin.listSuperAdmin(),
    ...options,
  });
}

export function useCreateSuperAdminMutation(
  options?: UseMutationOptions<null, Error, RegisterPayload>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: RegisterPayload) =>
      userService.admin.createSuperAdmin(payload),
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: userKeys.superAdmins() });
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}

export function useAdminChangePasswordMutation(
  options?: UseMutationOptions<
    User,
    Error,
    { id: number | string; payload: AdminChangePasswordPayload }
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, payload }) =>
      userService.admin.changePassword(id, payload),
    onSuccess: (data, variables, ...rest) => {
      queryClient.setQueryData(userKeys.detail(variables.id), data);
      options?.onSuccess?.(data, variables, ...rest);
    },
    ...options,
  });
}
