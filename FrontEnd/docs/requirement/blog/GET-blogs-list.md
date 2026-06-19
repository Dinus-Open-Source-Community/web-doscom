# GET `/api/v1/blogs`

## Ringkasan

List blog publik (published) dengan pagination dan filter kategori opsional.

## Authentication

Public

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False | Default `1` |
| limit | number | False | Default `10` |
| kategori[] | string[] | False | Max 3 kategori (query array) |

## Response Success

**Status:** `200`

Format **flat** (tanpa envelope):

```json
{
  "message": "successfully fetch data",
  "blogs": [
    {
      "id": 1,
      "title": "Judul Artikel",
      "slug": "judul-artikel",
      "kategori": "tech",
      "thumbnail_url": "https://..."
    }
  ],
  "totalPage": 3,
  "currentPage": 1
}
```

## Catatan

- Query param public: `kategori[]` (bukan `kategory`).
- Response key list: `blogs` (thumbnail view).
- Hanya blog published yang ditampilkan (filter di service).

## Frontend

- Service: `blogService.list`
- Hook: `useBlogsQuery`
- API path: `API_PATH.blogs.list`
