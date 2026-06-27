# GET `/api/v1/admin/works`

## Ringkasan

List works untuk admin panel dengan pagination. Koordinator hanya melihat works divisi sendiri (filter di service backend).

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`

| Group | Role JWT |
| --- | --- |
| ADMIN | `SuperAdmin` |
| KOOR | `KoorPemro`, `KoorJaringan`, `KoorData`, `KoorMedcrev`, `BPH` |

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False | Default `1` |
| limit | number | False | Default `10` |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Success fetching all works",
  "data": {
    "worksData": [
      {
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
      }
    ],
    "totalPage": 1,
    "currentPage": 0
  },
  "error": null
}
```

## Response Error

**403** — role tidak diizinkan.

## Catatan

- Field key `"worksData"` (camelCase) berbeda dari public list (`"work data"`).
- `currentPage` di BE dihitung `(offset/limit)*1` — page 1 bisa bernilai `0`.
- SuperAdmin melihat semua divisi; koordinator terfilter otomatis by role JWT.

## Frontend

- Service: `workService.admin.list`
- Hook: `useAdminWorksQuery`
- API path: `API_PATH.admin.works.list`
