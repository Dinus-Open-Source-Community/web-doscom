# GET `/api/v1/admin/pengurus/by-user/{user_id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — KOOR, BPH, ADMIN

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| user_id | integer | True |  |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Successfully get data",
  "data": {
    "id": 1,
    "id_user": 5
  }
}
```

## Frontend

- Service: `pengurusService.admin.getByUserId`
- Hook: `useAdminPengurusByUserQuery`
- API path: `API_PATH` di `lib/api-path.ts`
