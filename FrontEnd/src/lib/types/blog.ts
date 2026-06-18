import type { PaginatedMeta } from "./common";

export interface BlogGalleryImage {
  id: number;
  file_url?: string;
  thumbnail_url?: string;
}

export interface Blog {
  id: number;
  author_id: number;
  title: string;
  slug: string;
  content: string;
  kategori: string[];
  thumbnail_url: string;
  published_at?: string | null;
  blog_image?: BlogGalleryImage[];
}

export interface BlogThumbnail {
  id: number;
  title: string;
  slug: string;
  kategori: string;
  thumbnail_url: string;
}

export interface BlogListResponse extends PaginatedMeta {
  message?: string;
  blogs: BlogThumbnail[] | Blog[];
}
