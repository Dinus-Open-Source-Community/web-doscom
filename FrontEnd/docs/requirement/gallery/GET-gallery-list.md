# GET `/api/v1/gallery`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Public

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False |  |
| limit | number | False |  |
| start_year | string | False | Filter tahun |
| end_year | string | False | Filter tahun |

## Response Success

**Status:** `200`

```json
{
  "message": "Successfully get data",
  "gallery": [
    {
      "id": 1,
      "gallery_name": "Workshop",
      "gallery_type": "event",
      "description": "...",
      "event_date": "2026-03-15",
      "file_url": "https://..."
    }
  ],
  "totalPages": 2,
  "currentPage": 1
}
```

## Frontend

- Service: `galleryService.list`
- Hook: `useGalleryQuery`
- API path: `API_PATH` di `lib/api-path.ts`
