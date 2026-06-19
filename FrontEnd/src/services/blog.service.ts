import { api } from "../lib/axios";
import { API_PATH } from "../lib/api-path";
import { unwrapBlogDetail } from "../lib/func/blog";
import { getEnvelopeData, toFormData } from "../lib/func/http";
import type {
  Blog,
  BlogListResponse,
  PaginationQuery,
} from "../lib/types";

export interface PublicBlogQuery extends PaginationQuery {
  kategori?: string[];
}

export interface AdminBlogQuery extends PaginationQuery {
  kategory?: string[];
}

export interface CreateBlogPayload {
  title: string;
  slug: string;
  content: string;
  kategori: string[];
  status: string;
  published_at?: string;
  existingID_image?: number[];
}

export interface UpdateBlogPayload extends Partial<CreateBlogPayload> {}

export const blogService = {
  list(params?: PublicBlogQuery): Promise<BlogListResponse> {
    return getEnvelopeData<BlogListResponse>(API_PATH.blogs.list, { params });
  },

  getById(id: number | string): Promise<Blog> {
    return getEnvelopeData<Blog>(API_PATH.blogs.detail(id)).then((data) => {
      return unwrapBlogDetail(data as { blog?: Blog } & Blog);
    });
  },

  admin: {
    list(params?: AdminBlogQuery): Promise<BlogListResponse> {
      return api
        .get<BlogListResponse>(API_PATH.admin.blogs.list, { params })
        .then((response) => response.data);
    },

    getById(id: number | string): Promise<Blog> {
      return blogService.getById(id);
    },

    create(
      payload: CreateBlogPayload,
      files?: File[],
    ): Promise<{ message: string; data?: Blog }> {
      const formData = toFormData(payload, { files });
      return api
        .post(API_PATH.admin.blogs.list, formData)
        .then((response) => response.data);
    },

    update(
      id: number | string,
      payload: UpdateBlogPayload,
      files?: File[],
    ): Promise<{ message: string }> {
      const formData = toFormData(payload, { files });
      return api
        .put(API_PATH.admin.blogs.detail(id), formData)
        .then((response) => response.data);
    },

    remove(id: number | string): Promise<{ message: string }> {
      return api
        .delete(API_PATH.admin.blogs.detail(id))
        .then((response) => response.data);
    },
  },
};
