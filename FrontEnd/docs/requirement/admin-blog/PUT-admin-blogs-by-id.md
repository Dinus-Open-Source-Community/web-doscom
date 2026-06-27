# PUT `/api/v1/admin/blogs/{id}`

## Ringkasan

Update blog existing (partial update via form-data).

## Authentication

Bearer atau cookie — middleware `SuperAdmin`, `KoorMedcrev`

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Blog ID |

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |

## Request Body (Form Data)

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| title | string | False | Judul |
| slug | string | False | Slug |
| content | string | False | Konten |
| kategori | string[] | False | Kategori |
| status | string | False | Status valid |
| published_at | string | False | RFC3339 |
| existingID_image | number[] | False | Gallery ID retained |
| files | file[] | False | Upload gambar baru |

## Response Success

**Status:** `200`

Format **flat** (tanpa envelope):

```json
{
  "message": "successfully update blog",
  "data": {
    "id": 1,
    "author_id": 2,
    "title": "Judul Updated",
    "slug": "judul-updated",
    "content": "...",
    "kategori": ["tech"],
    "thumbnail_url": "https://...",
    "published_at": "2025-01-15T00:00:00Z",
    "blog_image": []
  }
}
```

## Response Error

**403** — role tidak diizinkan.

**400/500** — validasi atau server error.

## Frontend

- Service: `blogService.admin.update`
- Hook: `useUpdateBlogMutation`
- API path: `API_PATH.admin.blogs.detail(id)`
