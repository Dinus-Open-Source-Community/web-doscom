# GET `/api/v1/works`

## Ringkasan

List works public dengan pagination dan filter opsional by project type.

## Authentication

Public

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False | Default `1` |
| limit | number | False | Default `10` |
| projecttype | string | False | Filter tipe proyek, contoh: `web` |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Success fetching all works",
  "data": {
    "work data": [
      {
        "id": 1,
        "title": "Website DOSCOM",
        "tagline": "Official website",
        "description": "...",
        "slug": "website-doscom",
        "project_type": "web",
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

## Catatan

- Field key `"work data"` (dengan spasi) sesuai response backend.
- `currentPage` di BE dihitung `(offset/limit)*1` — page 1 bisa bernilai `0`.
- Hanya works dengan status published yang ditampilkan (filter di service).

## Frontend

- Service: `workService.list`
- Hook: `useWorksQuery` (ganti `useWorksByProjectTypeQuery` — deprecated)
- API path: `API_PATH.works.list`
