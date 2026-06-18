import { API_PATH } from "../lib/api-path";
import {
  deleteEnvelope,
  getEnvelopeData,
  postEnvelopeData,
  putEnvelopeData,
} from "../lib/func/http";
import type {
  AdminChangePasswordPayload,
  ChangePasswordPayload,
  PaginationQuery,
  RegisterPayload,
  User,
  UserPatch,
} from "../lib/types";

export const userService = {
  getMe(): Promise<User> {
    return getEnvelopeData<User>(API_PATH.user.me);
  },

  updateProfile(payload: UserPatch): Promise<User> {
    return putEnvelopeData<User, UserPatch>(API_PATH.user.profile, payload);
  },

  changePassword(payload: ChangePasswordPayload): Promise<null> {
    return putEnvelopeData<null, ChangePasswordPayload>(
      API_PATH.user.changePassword,
      payload,
    );
  },

  list(params?: PaginationQuery): Promise<User[]> {
    return getEnvelopeData<User[]>(API_PATH.user.list, { params });
  },

  getById(id: number | string): Promise<User> {
    return getEnvelopeData<User>(API_PATH.user.detail(id));
  },

  create(payload: RegisterPayload): Promise<null> {
    return postEnvelopeData<null, RegisterPayload>(API_PATH.user.list, payload);
  },

  update(id: number | string, payload: UserPatch): Promise<User> {
    return putEnvelopeData<User, UserPatch>(API_PATH.user.detail(id), payload);
  },

  remove(id: number | string): Promise<null> {
    return deleteEnvelope(API_PATH.user.detail(id)).then(() => null);
  },

  admin: {
    listSuperAdmin(): Promise<User[]> {
      return getEnvelopeData<User[]>(API_PATH.admin.user.list);
    },

    createSuperAdmin(payload: RegisterPayload): Promise<null> {
      return postEnvelopeData<null, RegisterPayload>(
        API_PATH.admin.user.superAdmin,
        payload,
      );
    },

    changePassword(
      id: number | string,
      payload: AdminChangePasswordPayload,
    ): Promise<null> {
      return putEnvelopeData<null, AdminChangePasswordPayload>(
        API_PATH.admin.user.changePassword(id),
        payload,
      );
    },
  },
};
