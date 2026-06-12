# GET `/api/v1/admin/blogs/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, KoorMedcrev

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True |  |

## Response Success

**Status:** `200`

```json
{
  "message": "blog found successfully",
  "blog": {
    "id": 1,
    "title": "Judul",
    "content": "...",
    "kategori": [
      "tech"
    ],
    "thumbnail_url": "https://..."
  }
}
```

## Frontend

- Service: `blogService.admin.getById`
- Hook: `useAdminBlogQuery`
- API path: `API_PATH` di `lib/api-path.ts`
