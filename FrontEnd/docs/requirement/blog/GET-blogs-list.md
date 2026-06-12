# GET `/api/v1/blogs`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Public

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False | Default 1 |
| limit | number | False | Default 10 |
| kategori[] | string[] | False | Max 3 kategori |

## Response Success

**Status:** `200`

```json
{
  "message": "successfully fetch data",
  "blogs": [
    {
      "id": 1,
      "title": "Judul",
      "slug": "judul",
      "kategori": "tech",
      "thumbnail_url": "https://..."
    }
  ],
  "totalPage": 3,
  "currentPage": 1
}
```

## Frontend

- Service: `blogService.list`
- Hook: `useBlogsQuery`
- API path: `API_PATH` di `lib/api-path.ts`
