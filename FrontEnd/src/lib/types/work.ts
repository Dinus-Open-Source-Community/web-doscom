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
  gallery?: string[];
}

export interface WorkInternal extends WorkPublic {
  pengurus_id?: number;
  status: string;
}

export interface WorkListData extends PaginatedMeta {
  "work data"?: WorkPublic[];
  worksData?: WorkInternal[];
}

export type WorkListResponse = WorkListData & { message?: string };

export interface WorkUpdateStatusPayload {
  status: string;
}
