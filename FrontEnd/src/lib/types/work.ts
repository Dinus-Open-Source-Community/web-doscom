import type { PaginatedMeta } from "./common";

export interface WorkPublic {
  id: number;
  title: string;
  tagline: string;
  description: string;
  slug: string;
  project_type: string;
  technologies: string[];
  project_date: string;
  image_url: string;
  gallery?: unknown;
}

export interface WorkInternal extends WorkPublic {
  status: string;
}

export interface WorkListResponse extends PaginatedMeta {
  message?: string;
  "work data"?: WorkPublic[];
  worksData?: WorkInternal[];
}
