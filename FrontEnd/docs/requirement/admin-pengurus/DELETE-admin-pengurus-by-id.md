# DELETE `/api/v1/admin/pengurus/delete/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — KOOR, BPH, ADMIN

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
  "message": "pengurus deleted",
  "data": null
}
```

## Frontend

- Service: `pengurusService.admin.remove`
- Hook: `useDeleteAdminPengurusMutation`
- API path: `API_PATH` di `lib/api-path.ts`
