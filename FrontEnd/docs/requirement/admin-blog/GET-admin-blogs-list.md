# GET `/api/v1/admin/blogs`

## Ringkasan

List semua blog untuk admin (semua status) dengan pagination dan filter kategori.

## Authentication

Bearer atau cookie — middleware `SuperAdmin`, `KoorMedcrev`

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False | Default `1` |
| limit | number | False | Default `10` (BE bug: handler baca `page` dua kali untuk limit) |
| kategory[] | string[] | False | Max 3 kategori — **typo backend**, bukan `kategori[]` |

## Response Success

**Status:** `200`

Format **flat** (tanpa envelope):

```json
{
  "message": "successfully fetch data",
  "blogs": [
    {
      "id": 1,
      "title": "Draft Artikel",
      "slug": "draft-artikel",
      "kategori": "tech",
      "thumbnail_url": "https://..."
    }
  ],
  "totalPage": 5,
  "currentPage": 1
}
```

## Catatan

- Query filter admin: **`kategory[]`** (typo di `BlogHandler.ListBlogs`).
- Public list memakai `kategori[]` — jangan tukar param antara public vs admin.
- Frontend `AdminBlogQuery` sudah mirror typo backend.

## Frontend

- Service: `blogService.admin.list`
- Hook: `useAdminBlogsQuery`
- API path: `API_PATH.admin.blogs.list`
