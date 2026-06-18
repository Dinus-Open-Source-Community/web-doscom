# GET `/api/v1/admin/works`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, Koordinator

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False |  |
| limit | number | False |  |

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
        "title": "Project",
        "status": "published"
      }
    ],
    "totalPage": 1,
    "currentPage": 1
  }
}
```

## Catatan

- Koordinator hanya melihat works divisi sendiri.

## Frontend

- Service: `workService.admin.list`
- Hook: `useAdminWorksQuery`
- API path: `API_PATH` di `lib/api-path.ts`
