# GET `/api/v1/works/{projecttype}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Public

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| projecttype | string | True | Tipe proyek, contoh: web |

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False |  |
| limit | number | False |  |
| projecttype | string | False | Query alternatif |

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
        "project_type": "web",
        "technologies": [
          "Astro"
        ],
        "image_url": "https://..."
      }
    ],
    "totalPage": 1,
    "currentPage": 1
  }
}
```

## Catatan

- Route public: `GET /works/:projecttype`.

## Frontend

- Service: `workService.listByProjectType`
- Hook: `useWorksByProjectTypeQuery`
- API path: `API_PATH` di `lib/api-path.ts`
