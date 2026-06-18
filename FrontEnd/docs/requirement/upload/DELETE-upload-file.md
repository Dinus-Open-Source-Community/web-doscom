# DELETE `/api/v1/upload/file`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — semua user login

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |
| Authorization | `Bearer {access_token}` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| file_name | string | True | Nama file di MinIO |

**Contoh:**

```json
{
  "file_name": "photo-1.webp"
}
```

## Response Success

**Status:** `200`

```json
{
  "success": true,
  "message": "File deleted successfully"
}
```

## Catatan

- `POST /upload/image` belum aktif (commented di backend).

## Frontend

- Service: `uploadService.deleteFile`
- Hook: `useDeleteUploadFileMutation`
- API path: `API_PATH` di `lib/api-path.ts`
