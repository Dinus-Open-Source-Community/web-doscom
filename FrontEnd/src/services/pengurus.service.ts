import { API_PATH } from "../lib/api-path";
import {
  deleteEnvelope,
  getEnvelopeData,
  postEnvelopeData,
  putEnvelopeData,
  toFormData,
} from "../lib/func/http";
import type { Pengurus, PengurusPublic } from "../lib/types";

export interface CreatePengurusPayload {
  id_user?: number;
  email?: string;
  divisi: string;
  name: string;
  position: string;
  sosmed?: string[];
  start_periode_year: number;
  end_periode_year: number;
}

export interface UpdatePengurusPayload extends Partial<CreatePengurusPayload> {}

export const pengurusService = {
  listByDivision(division: string): Promise<PengurusPublic[]> {
    return getEnvelopeData<PengurusPublic[]>(
      API_PATH.pengurus.byDivision(division),
    );
  },

  getProfile(): Promise<Pengurus> {
    return getEnvelopeData<Pengurus>(API_PATH.pengurus.profile);
  },

  createProfile(
    payload: CreatePengurusPayload,
    file?: File,
  ): Promise<Pengurus> {
    const formData = toFormData(payload, file ? { file } : undefined);
    return postEnvelopeData<Pengurus>(API_PATH.pengurus.createProfile, formData);
  },

  updateMe(
    payload: UpdatePengurusPayload,
    file?: File,
  ): Promise<Pengurus> {
    const formData = toFormData(payload, file ? { file } : undefined);
    return putEnvelopeData<Pengurus>(API_PATH.pengurus.me, formData);
  },

  deleteMe(): Promise<null> {
    return deleteEnvelope(API_PATH.pengurus.me).then(() => null);
  },

  admin: {
    list(divisi?: string): Promise<Pengurus[]> {
      return getEnvelopeData<Pengurus[]>(API_PATH.admin.pengurus.list, {
        params: divisi ? { divisi } : undefined,
      });
    },

    getById(id: number | string): Promise<Pengurus> {
      return getEnvelopeData<Pengurus>(API_PATH.admin.pengurus.detail(id));
    },

    getByUserId(userId: number | string): Promise<Pengurus> {
      return getEnvelopeData<Pengurus>(
        API_PATH.admin.pengurus.byUser(userId),
      );
    },

    create(
      payload: CreatePengurusPayload,
      file?: File,
    ): Promise<Pengurus> {
      const formData = toFormData(payload, file ? { file } : undefined);
      return postEnvelopeData<Pengurus>(
        API_PATH.admin.pengurus.list,
        formData,
      );
    },

    update(
      id: number | string,
      payload: UpdatePengurusPayload,
      file?: File,
    ): Promise<Pengurus> {
      const formData = toFormData(payload, file ? { file } : undefined);
      return putEnvelopeData<Pengurus>(
        API_PATH.admin.pengurus.detail(id),
        formData,
      );
    },

    remove(id: number | string): Promise<null> {
      return deleteEnvelope(API_PATH.admin.pengurus.delete(id)).then(
        () => null,
      );
    },
  },
};
