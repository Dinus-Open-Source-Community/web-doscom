# POST `/api/v1/admin/gallery`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, KoorMedcrev

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |
| Authorization | `Bearer {access_token}` |

## Form Data

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| gallery_name | string | True |  |
| gallery_type | string | True | fun|proker|event|... |
| description | string | True |  |
| event_date | string | True | YYYY-MM-DD |
| files | file[] | True | Max 5 file |

## Response Success

**Status:** `200`

```json
{
  "message": "Successfully insert data",
  "data": [
    {
      "id": 1,
      "file_url": "https://..."
    }
  ]
}
```

## Response Error

**400**

```json
{"error": "Max upload 5 file"}
```

**403**

```json
{"error": "You're not allowed broo"}
```

## Frontend

- Service: `galleryService.admin.create`
- Hook: `useCreateGalleryMutation`
- API path: `API_PATH` di `lib/api-path.ts`
