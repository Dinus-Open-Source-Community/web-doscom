import { API_PATH } from "../lib/api-path";
import {
  deleteEnvelope,
  getEnvelopeData,
  postEnvelopeData,
  toFormData,
} from "../lib/func/http";
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

interface GalleryListData {
  totalPages?: number;
  currentPage?: number;
  gallery: GalleryItem[];
}

export const galleryService = {
  list(params?: GalleryQuery): Promise<GalleryListResponse> {
    return getEnvelopeData<GalleryListData>(API_PATH.gallery.list, { params });
  },

  admin: {
    create(
      payload: CreateGalleryPayload,
      files: File[],
    ): Promise<GalleryItem[]> {
      const formData = toFormData(payload, { files });
      return postEnvelopeData<GalleryItem[]>(
        API_PATH.admin.gallery.list,
        formData,
      );
    },

    remove(id: number | string): Promise<null> {
      return deleteEnvelope(API_PATH.admin.gallery.detail(id)).then(() => null);
    },
  },
};
