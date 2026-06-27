# GET `/api/v1/blogs/{id}`

## Ringkasan

Detail blog by ID (public view).

## Authentication

Public

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Blog ID |

## Response Success

**Status:** `200`

Format **flat** (tanpa envelope):

```json
{
  "message": "blog found successfully",
  "blog": {
    "id": 1,
    "author_id": 2,
    "title": "Judul Artikel",
    "slug": "judul-artikel",
    "content": "<p>Konten HTML...</p>",
    "kategori": ["tech", "doscom"],
    "thumbnail_url": "https://...",
    "published_at": "2025-01-15T00:00:00Z",
    "blog_image": [
      {
        "id": 10,
        "file_url": "https://..."
      }
    ]
  }
}
```

## Response Error

**400** — ID invalid.

**404** — blog tidak ditemukan.

## Catatan

- Frontend helper `unwrapBlogDetail()` menormalisasi key `blog` vs root object.
- Endpoint yang sama dipakai admin detail (`blogService.admin.getById` delegasi ke public).

## Frontend

- Service: `blogService.getById`
- Hook: `useBlogQuery`, `useAdminBlogQuery`
- API path: `API_PATH.blogs.detail(id)`
