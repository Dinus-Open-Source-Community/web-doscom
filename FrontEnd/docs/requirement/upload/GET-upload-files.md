# GET `/api/v1/upload/files`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — semua user login

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| category | string | True | gallery|blog|work|pengurus |

## Response Success

**Status:** `200`

```json
{
  "success": true,
  "message": "Files retrieved successfully",
  "files": [
    "photo-1.webp"
  ],
  "count": 1
}
```

## Frontend

- Service: `uploadService.listFiles`
- Hook: `useUploadFilesQuery`
- API path: `API_PATH` di `lib/api-path.ts`
