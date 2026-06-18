import { api } from "../lib/axios";
import { API_PATH } from "../lib/api-path";
import type {
  DeleteFilePayload,
  UploadCategory,
  UploadDeleteResponse,
  UploadListResponse,
} from "../lib/types";

export const uploadService = {
  listFiles(params: UploadCategory): Promise<UploadListResponse> {
    return api
      .get<UploadListResponse>(API_PATH.upload.files, { params })
      .then((response) => response.data);
  },

  deleteFile(payload: DeleteFilePayload): Promise<UploadDeleteResponse> {
    return api
      .delete<UploadDeleteResponse>(API_PATH.upload.file, { data: payload })
      .then((response) => response.data);
  },
};
