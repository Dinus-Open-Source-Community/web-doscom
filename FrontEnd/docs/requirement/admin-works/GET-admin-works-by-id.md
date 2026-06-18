# GET `/api/v1/admin/works/{id}`

## Ringkasan

Detail work internal (termasuk `status`, `pengurus_id`) untuk halaman edit admin.

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`

| Group | Role JWT |
| --- | --- |
| ADMIN | `SuperAdmin` |
| KOOR | `KoorPemro`, `KoorJaringan`, `KoorData`, `KoorMedcrev`, `BPH` |

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Work ID |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Success fetching work detail",
  "data": {
    "id": 1,
    "pengurus_id": 2,
    "title": "Website DOSCOM",
    "tagline": "Official website",
    "description": "...",
    "slug": "website-doscom",
    "project_type": "web",
    "status": "published",
    "technologies": ["Astro", "Go"],
    "project_date": "2025-01-15T00:00:00Z",
    "image_url": "https://...",
    "gallery": []
  },
  "error": null
}
```

## Response Error

**400** — ID invalid.

**404** — work tidak ditemukan atau di luar scope divisi koordinator.

## Catatan

- Response internal berbeda dari public detail (public tidak expose `status` / `pengurus_id`).
- Koordinator hanya bisa akses work divisi sendiri.

## Frontend

- Service: `workService.admin.getById`
- Hook: `useAdminWorkQuery`
- API path: `API_PATH.admin.works.detail(id)`
