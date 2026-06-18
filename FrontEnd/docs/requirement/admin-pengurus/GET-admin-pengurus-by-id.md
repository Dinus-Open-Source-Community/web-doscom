# GET `/api/v1/admin/pengurus/{id}`

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
  "message": "Successfully get data",
  "data": {
    "id": 1,
    "name": "Member",
    "email": "m@example.com"
  }
}
```

## Frontend

- Service: `pengurusService.admin.getById`
- Hook: `useAdminPengurusQuery`
- API path: `API_PATH` di `lib/api-path.ts`
