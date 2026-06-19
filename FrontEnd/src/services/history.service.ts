import { API_PATH } from "../lib/api-path";
import { getEnvelopeData } from "../lib/func/http";
import type { HistoryItem, HistoryListResponse, PaginationQuery } from "../lib/types";

export interface HistoryQuery extends PaginationQuery {}

export const historyService = {
  list(params?: HistoryQuery): Promise<HistoryListResponse> {
    return getEnvelopeData<HistoryListResponse>(API_PATH.history.list, { params });
  },

  getById(id: number | string): Promise<HistoryItem> {
    return getEnvelopeData<HistoryItem>(API_PATH.history.detail(id));
  },
};
