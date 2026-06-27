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
  WorkListData,
  WorkListResponse,
  WorkPublic,
  WorkUpdateStatusPayload,
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
  list(params?: PublicWorkQuery): Promise<WorkListResponse> {
    return getEnvelopeData<WorkListData>(API_PATH.works.list, { params });
  },

  getById(id: number | string): Promise<WorkPublic> {
    return getEnvelopeData<WorkPublic>(API_PATH.works.detail(id));
  },

  getTypes(): Promise<string[]> {
    return getEnvelopeData<string[]>(API_PATH.works.types);
  },

  admin: {
    list(params?: PaginationQuery): Promise<WorkListResponse> {
      return getEnvelopeData<WorkListData>(API_PATH.admin.works.list, { params });
    },

    getById(id: number | string): Promise<WorkInternal> {
      return getEnvelopeData<WorkInternal>(API_PATH.admin.works.detail(id));
    },

    create(payload: CreateWorkPayload, files?: File[]): Promise<WorkInternal> {
      const formData = buildWorkFormData(payload, files);
      return postEnvelopeData<WorkInternal>(API_PATH.admin.works.list, formData);
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

    updateStatus(
      id: number | string,
      payload: WorkUpdateStatusPayload,
    ): Promise<WorkInternal> {
      return putEnvelopeData<WorkInternal, WorkUpdateStatusPayload>(
        API_PATH.admin.works.status(id),
        payload,
      );
    },

    remove(id: number | string): Promise<null> {
      return deleteEnvelope(API_PATH.admin.works.detail(id)).then(() => null);
    },
  },
};
