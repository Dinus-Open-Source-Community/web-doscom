# GET `/api/v1/gallery`

## Ringkasan

List gallery publik dengan pagination dan filter tahun opsional.

## Authentication

Public

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False | Default `1` |
| limit | number | False | Default `10` |
| start_year | string | False | Filter tahun mulai (event_date) |
| end_year | string | False | Filter tahun akhir (event_date) |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Successfully get data",
  "data": {
    "totalPages": 3,
    "currentPage": 1,
    "gallery": [
      {
        "id": 1,
        "id_users": 2,
        "file_upload_id": 10,
        "gallery_name": "DOSCOM Fun Day 2024",
        "gallery_type": "fun",
        "description": "Kegiatan fun day",
        "event_date": "2024-08-15T00:00:00Z",
        "file_url": "https://..."
      }
    ]
  },
  "error": null
}
```

## Catatan

- Meta pagination: `totalPages`, `currentPage` (camelCase) di dalam `data`.
- Array items: key `gallery`.
- Frontend service unwrap `data` ke `GalleryListResponse`.

## Frontend

- Service: `galleryService.list`
- Hook: `useGalleryQuery`
- API path: `API_PATH.gallery.list`
