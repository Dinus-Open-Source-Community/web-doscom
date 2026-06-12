# POST `/api/v1/admin/blogs`

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
| title | string | True |  |
| slug | string | True | Unique |
| content | string | True | HTML/content |
| kategori | string[] | True | Max 3 |
| status | string | True | draft|published|scheduled|... |
| published_at | datetime | False | RFC3339 |
| existingID_image | number[] | False | Reuse gallery IDs |
| files | file[] | False | Upload gambar baru |

## Response Success

**Status:** `200`

```json
{
  "message": "successfully create blog",
  "data": {
    "id": 1,
    "title": "Judul"
  }
}
```

## Frontend

- Service: `blogService.admin.create`
- Hook: `useCreateBlogMutation`
- API path: `API_PATH` di `lib/api-path.ts`
