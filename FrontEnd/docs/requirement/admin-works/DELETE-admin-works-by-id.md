# DELETE `/api/v1/admin/works/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, Koordinator

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True |  |

## Response Success

**Status:** `204`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Work deleted successfully",
  "data": null
}
```

## Frontend

- Service: `workService.admin.remove`
- Hook: `useDeleteWorkMutation`
- API path: `API_PATH` di `lib/api-path.ts`
