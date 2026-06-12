import { api } from "../lib/axios";
import { API_PATH } from "../lib/api-path";
import {
  deleteEnvelope,
  getEnvelopeData,
  postEnvelopeData,
  putEnvelopeData,
} from "../lib/func/http";
import { buildWorkFormData } from "../lib/func/work";
import type {
  PaginationQuery,
  WorkInternal,
  WorkListResponse,
  WorkPublic,
} from "../lib/types";

export interface PublicWorkQuery extends PaginationQuery {
  projecttype?: string;
}

export interface CreateWorkPayload {
  pengurus_id: number;
  title: string;
  tagline: string;
  description: string;
  slug: string;
  project_type: string;
  technologies: string[];
  project_date: string;
  status: string;
  division?: string;
  existingID_image?: number[];
}

export interface UpdateWorkPayload extends Partial<CreateWorkPayload> {}

export const workService = {
  listByProjectType(
    projectType: string,
    params?: PaginationQuery,
  ): Promise<WorkListResponse> {
    return api
      .get<WorkListResponse>(API_PATH.works.byProjectType(projectType), {
        params,
      })
      .then((response) => response.data);
  },

  admin: {
    list(params?: PaginationQuery): Promise<WorkListResponse> {
      return api
        .get<WorkListResponse>(API_PATH.admin.works.list, { params })
        .then((response) => response.data);
    },

    getById(id: number | string): Promise<WorkInternal> {
      return getEnvelopeData<WorkInternal>(API_PATH.admin.works.detail(id));
    },

    create(payload: CreateWorkPayload, files?: File[]): Promise<WorkPublic> {
      const formData = buildWorkFormData(payload, files);
      return postEnvelopeData<WorkPublic>(API_PATH.admin.works.list, formData);
    },

    update(
      id: number | string,
      payload: UpdateWorkPayload,
      files?: File[],
    ): Promise<WorkInternal> {
      const formData = buildWorkFormData(payload, files);
      return putEnvelopeData<WorkInternal>(
        API_PATH.admin.works.detail(id),
        formData,
      );
    },

    remove(id: number | string): Promise<null> {
      return deleteEnvelope(API_PATH.admin.works.detail(id)).then(() => null);
    },
  },
};
