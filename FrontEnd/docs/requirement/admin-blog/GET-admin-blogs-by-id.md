# GET `/api/v1/admin/blogs/{id}`

## Ringkasan

Detail blog untuk admin edit. Route terdaftar di admin group, tetapi handler delegasi ke `GetBlogByID` (sama dengan public).

## Authentication

Bearer atau cookie — middleware `SuperAdmin`, `KoorMedcrev`

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
    "content": "<p>Konten...</p>",
    "kategori": ["tech"],
    "thumbnail_url": "https://...",
    "published_at": "2025-01-15T00:00:00Z",
    "blog_image": []
  }
}
```

## Frontend

- Service: `blogService.admin.getById` (delegasi ke `blogService.getById`)
- Hook: `useAdminBlogQuery`
- API path: `API_PATH.admin.blogs.detail(id)`
