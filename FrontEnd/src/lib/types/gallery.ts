import type { PaginatedMeta } from "./common";

export interface GalleryItem {
  id: number;
  id_users: number;
  file_upload_id: number;
  gallery_name: string;
  gallery_type: string;
  description: string;
  event_date: string;
  file_url: string;
}

export interface GalleryListResponse extends PaginatedMeta {
  message?: string;
  gallery: GalleryItem[];
}
