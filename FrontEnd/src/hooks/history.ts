import {
  useQuery,
  type UseQueryOptions,
} from "@tanstack/react-query";
import {
  historyService,
  type HistoryQuery,
} from "../services/history.service";
import type { HistoryItem, HistoryListResponse } from "../lib/types";
import { historyKeys } from "./keys";

export function useHistoryQuery(
  params?: HistoryQuery,
  options?: Omit<
    UseQueryOptions<HistoryListResponse>,
    "queryKey" | "queryFn"
  >,
) {
  return useQuery({
    queryKey: historyKeys.list(params),
    queryFn: () => historyService.list(params),
    ...options,
  });
}

export function useHistoryDetailQuery(
  id: number | string,
  options?: Omit<UseQueryOptions<HistoryItem>, "queryKey" | "queryFn">,
) {
  return useQuery({
    queryKey: historyKeys.detail(id),
    queryFn: () => historyService.getById(id),
    enabled: Boolean(id),
    ...options,
  });
}
