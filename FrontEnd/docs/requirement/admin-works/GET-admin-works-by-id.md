# GET `/api/v1/admin/works/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, Koordinator

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True |  |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Success fetching work detail",
  "data": {
    "id": 1,
    "title": "Project",
    "status": "published",
    "technologies": [
      "Go"
    ]
  }
}
```

## Frontend

- Service: `workService.admin.getById`
- Hook: `useAdminWorkQuery`
- API path: `API_PATH` di `lib/api-path.ts`
