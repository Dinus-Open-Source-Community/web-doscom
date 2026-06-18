import { api } from "../lib/axios";
import { API_PATH } from "../lib/api-path";
import { toFormData } from "../lib/func/http";
import type {
  GalleryItem,
  GalleryListResponse,
  PaginationQuery,
} from "../lib/types";

export interface GalleryQuery extends PaginationQuery {
  start_year?: string;
  end_year?: string;
}

export interface CreateGalleryPayload {
  gallery_name: string;
  gallery_type: string;
  description: string;
  event_date: string;
}

export const galleryService = {
  list(params?: GalleryQuery): Promise<GalleryListResponse> {
    return api
      .get<GalleryListResponse>(API_PATH.gallery.list, { params })
      .then((response) => response.data);
  },

  admin: {
    create(
      payload: CreateGalleryPayload,
      files: File[],
    ): Promise<{ message: string; data: GalleryItem[] }> {
      const formData = toFormData(payload, { files });
      return api
        .post(API_PATH.admin.gallery.list, formData)
        .then((response) => response.data);
    },

    remove(id: number | string): Promise<{ message: string }> {
      return api
        .delete<{ message: string }>(API_PATH.admin.gallery.detail(id))
        .then((response) => response.data);
    },
  },
};
