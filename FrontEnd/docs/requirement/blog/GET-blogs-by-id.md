# GET `/api/v1/blogs/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Public

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True | Blog ID |

## Response Success

**Status:** `200`

```json
{
  "message": "blog found successfully",
  "blog": {
    "id": 1,
    "author_id": 2,
    "title": "Judul",
    "slug": "judul",
    "content": "<p>...</p>",
    "kategori": [
      "tech"
    ],
    "thumbnail_url": "https://...",
    "blog_image": []
  }
}
```

## Frontend

- Service: `blogService.getById`
- Hook: `useBlogQuery`
- API path: `API_PATH` di `lib/api-path.ts`
