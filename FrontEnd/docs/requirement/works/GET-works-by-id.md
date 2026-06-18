# GET `/api/v1/works/{id}`

## Ringkasan

Detail work public by ID.

## Authentication

Public

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Work ID |

## Response Success

**Status:** `200`

```json
{
  "success": true,
  "message": "Success fetching work detail",
  "data": {
    "id": 1,
    "title": "Website DOSCOM",
    "tagline": "Official website",
    "description": "...",
    "slug": "website-doscom",
    "project_type": "web",
    "technologies": ["Astro"],
    "project_date": "2025-01-15T00:00:00Z",
    "image_url": "https://...",
    "gallery": []
  },
  "error": null
}
```

Response type: `WorkResponseClient` — tidak menyertakan `pengurus_id`, `status`.

## Response Error

**404** — work tidak ditemukan atau belum published.

## Frontend

- Service: `workService.getById`
- Hook: `useWorkQuery`
- API path: `API_PATH.works.detail(id)`
