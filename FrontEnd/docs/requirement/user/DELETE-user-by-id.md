# DELETE `/api/v1/user/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — ADMIN, KOOR, BPH

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True | User ID |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "user deleted info kopi dan gorengan bolo",
  "data": null
}
```

## Frontend

- Service: `userService.remove`
- Hook: `useDeleteUserMutation`
- API path: `API_PATH` di `lib/api-path.ts`
