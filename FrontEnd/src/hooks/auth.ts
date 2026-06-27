import {
  useMutation,
  useQueryClient,
  type UseMutationOptions,
} from "@tanstack/react-query";
import { authService } from "../services/auth.service";
import type {
  AuthResponse,
  LoginPayload,
  RegisterPayload,
} from "../lib/types";
import { userKeys } from "./keys";

export function useLoginMutation(
  options?: UseMutationOptions<AuthResponse, Error, LoginPayload>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: LoginPayload) => authService.login(payload),
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: userKeys.me() });
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}

export function useRegisterMutation(
  options?: UseMutationOptions<AuthResponse, Error, RegisterPayload>,
) {
  return useMutation({
    mutationFn: (payload: RegisterPayload) => authService.register(payload),
    ...options,
  });
}

export function useRefreshMutation(
  options?: UseMutationOptions<AuthResponse, Error, void>,
) {
  return useMutation({
    mutationFn: () => authService.refresh(),
    ...options,
  });
}

export function useLogoutMutation(
  options?: UseMutationOptions<AuthResponse, Error, void>,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: () => authService.logout(),
    onSuccess: (...args) => {
      queryClient.clear();
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}
