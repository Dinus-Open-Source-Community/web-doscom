export interface UploadCategory {
  category: "gallery" | "blog" | "work" | "pengurus";
}

export interface DeleteFilePayload {
  file_name: string;
}

export interface UploadListResponse {
  success: boolean;
  message: string;
  files?: string[];
  count?: number;
}

export interface UploadDeleteResponse {
  success: boolean;
  message: string;
}
