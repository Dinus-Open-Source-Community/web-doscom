# POST `/api/v1/admin/gallery`

## Ringkasan

Upload satu atau lebih file gallery (max 5) dengan metadata.

## Authentication

Bearer atau cookie — middleware `SuperAdmin`, `KoorMedcrev`

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |

## Request Body (Form Data)

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| gallery_name | string | True | Nama gallery |
| gallery_type | string | True | Tipe: `fun`, `proker`, `achievment`, `work`, `activity`, dll. |
| description | string | True | Deskripsi |
| event_date | string | True | Format `YYYY-MM-DD` |
| files | file[] | True | Max 5 file |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`. `data` adalah array item yang dibuat.

```json
{
  "success": true,
  "message": "Successfully insert data",
  "data": [
    {
      "id": 1,
      "id_users": 2,
      "file_upload_id": 10,
      "gallery_name": "DOSCOM Fun Day",
      "gallery_type": "fun",
      "description": "Kegiatan fun day",
      "event_date": "2024-08-15T00:00:00Z",
      "file_url": "https://..."
    }
  ],
  "error": null
}
```

## Response Error

**400** — validasi gagal / no files / max 5 exceeded.

**403** — role tidak diizinkan.

## Frontend

- Service: `galleryService.admin.create`
- Hook: `useCreateGalleryMutation`
- API path: `API_PATH.admin.gallery.list`
