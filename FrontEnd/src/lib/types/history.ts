export interface HistoryPhoto {
  id: number
  id_history: number
  imager_url: string
}

export interface HistoryItem {
  id: number
  id_author: number
  title: string
  year: string
  description: string
  display_order: number
  photos: HistoryPhoto[]
}

export interface HistoryListResponse {
  message?: string
  history: HistoryItem[]
  totalPage: number
  currentPage: number
}
