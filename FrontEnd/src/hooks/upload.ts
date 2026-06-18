import {
  useMutation,
  useQuery,
  useQueryClient,
  type UseMutationOptions,
  type UseQueryOptions,
} from "@tanstack/react-query";
import { uploadService } from "../services/upload.service";
import type {
  DeleteFilePayload,
  UploadCategory,
  UploadDeleteResponse,
  UploadListResponse,
} from "../lib/types";
import { uploadKeys } from "./keys";

export function useUploadFilesQuery(
  category: UploadCategory["category"],
  options?: Omit<
    UseQueryOptions<UploadListResponse>,
    "queryKey" | "queryFn"
  >,
) {
  return useQuery({
    queryKey: uploadKeys.files({ category }),
    queryFn: () => uploadService.listFiles({ category }),
    enabled: Boolean(category),
    ...options,
  });
}

export function useDeleteUploadFileMutation(
  options?: UseMutationOptions<
    UploadDeleteResponse,
    Error,
    DeleteFilePayload
  >,
) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: DeleteFilePayload) =>
      uploadService.deleteFile(payload),
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: uploadKeys.all });
      options?.onSuccess?.(...args);
    },
    ...options,
  });
}
