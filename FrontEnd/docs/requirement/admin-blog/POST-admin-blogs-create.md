# POST `/api/v1/admin/blogs`

## Ringkasan

Buat blog baru dengan konten, kategori, status, dan gambar.

## Authentication

Bearer atau cookie — middleware `SuperAdmin`, `KoorMedcrev`

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |

## Request Body (Form Data)

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| title | string | True | Judul |
| slug | string | True | Slug unik |
| content | string | True | Konten (HTML/markdown) |
| kategori | string[] | True | Array kategori |
| status | string | True | `draft`, `published`, `unpublished`, `rejected`, `pending_review` |
| published_at | string | False | RFC3339 datetime |
| existingID_image | number[] | False | ID gallery existing |
| files | file[] | False | Upload gambar baru |

## Response Success

**Status:** `200`

Format **flat** (tanpa envelope):

```json
{
  "message": "successfully create blog",
  "data": {
    "id": 1,
    "author_id": 2,
    "title": "Judul Artikel",
    "slug": "judul-artikel",
    "content": "...",
    "kategori": ["tech"],
    "thumbnail_url": "https://...",
    "published_at": null,
    "blog_image": []
  }
}
```

## Response Error

**403** — role tidak diizinkan.

**400** — validasi gagal.

## Catatan

- Tidak ada status `scheduled`.
- Status default di service: `draft` jika kosong.

## Frontend

- Service: `blogService.admin.create`
- Hook: `useCreateBlogMutation`
- API path: `API_PATH.admin.blogs.list`
