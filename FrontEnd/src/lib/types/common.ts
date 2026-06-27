export interface ApiEnvelope<T = unknown> {
  success: boolean;
  message?: string;
  data?: T;
  error?: unknown;
}

export interface PaginationQuery {
  page?: number;
  limit?: number;
}

export interface PaginatedMeta {
  totalPage?: number;
  totalPages?: number;
  currentPage?: number;
}

export interface MessageResponse {
  message: string;
}
