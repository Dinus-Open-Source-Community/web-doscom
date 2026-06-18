import type { PaginationQuery } from "../lib/types";
import type { AdminBlogQuery, PublicBlogQuery } from "../services/blog.service";
import type { GalleryQuery } from "../services/gallery.service";
import type { PublicWorkQuery } from "../services/work.service";
import type { UploadCategory } from "../lib/types";

export const authKeys = {
  all: ["auth"] as const,
};

export const userKeys = {
  all: ["users"] as const,
  me: () => [...userKeys.all, "me"] as const,
  lists: () => [...userKeys.all, "list"] as const,
  list: (params?: PaginationQuery) => [...userKeys.lists(), params] as const,
  details: () => [...userKeys.all, "detail"] as const,
  detail: (id: number | string) => [...userKeys.details(), id] as const,
  superAdmins: () => [...userKeys.all, "super-admin"] as const,
};

export const blogKeys = {
  all: ["blogs"] as const,
  lists: () => [...blogKeys.all, "list"] as const,
  list: (params?: PublicBlogQuery) => [...blogKeys.lists(), params] as const,
  details: () => [...blogKeys.all, "detail"] as const,
  detail: (id: number | string) => [...blogKeys.details(), id] as const,
  admin: {
    all: ["blogs", "admin"] as const,
    lists: () => ["blogs", "admin", "list"] as const,
    list: (params?: AdminBlogQuery) =>
      ["blogs", "admin", "list", params] as const,
    detail: (id: number | string) =>
      ["blogs", "admin", "detail", id] as const,
  },
};

export const galleryKeys = {
  all: ["gallery"] as const,
  lists: () => [...galleryKeys.all, "list"] as const,
  list: (params?: GalleryQuery) => [...galleryKeys.lists(), params] as const,
};

export const workKeys = {
  all: ["works"] as const,
  lists: () => [...workKeys.all, "list"] as const,
  list: (params?: PublicWorkQuery) => [...workKeys.lists(), params] as const,
  details: () => [...workKeys.all, "detail"] as const,
  detail: (id: number | string) => [...workKeys.details(), id] as const,
  types: () => [...workKeys.all, "types"] as const,
  admin: {
    all: ["works", "admin"] as const,
    lists: () => ["works", "admin", "list"] as const,
    list: (params?: PaginationQuery) =>
      ["works", "admin", "list", params] as const,
    detail: (id: number | string) =>
      ["works", "admin", "detail", id] as const,
  },
};

export const pengurusKeys = {
  all: ["pengurus"] as const,
  byDivision: (division: string) =>
    [...pengurusKeys.all, "division", division] as const,
  profile: () => [...pengurusKeys.all, "profile"] as const,
  admin: {
    all: ["pengurus", "admin"] as const,
    lists: () => ["pengurus", "admin", "list"] as const,
    list: (divisi?: string) => ["pengurus", "admin", "list", divisi] as const,
    detail: (id: number | string) =>
      ["pengurus", "admin", "detail", id] as const,
    byUser: (userId: number | string) =>
      ["pengurus", "admin", "user", userId] as const,
  },
};

export const uploadKeys = {
  all: ["upload"] as const,
  files: (params: UploadCategory) =>
    [...uploadKeys.all, "files", params.category] as const,
};
