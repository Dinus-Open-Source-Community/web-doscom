# DELETE `/api/v1/user/{id}`

## Ringkasan

Hapus user by ID (scoped by role caller).

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`, `BPH`

Role `PENGURUS` ditolak (**403**).

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | User ID |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "user deleted info kopi dan gorengan bolo",
  "data": null,
  "error": null
}
```

## Response Error

**400** — ID invalid.

**403** — role tidak diizinkan.

## Frontend

- Service: `userService.remove`
- Hook: `useDeleteUserMutation`
- API path: `API_PATH.user.detail(id)`
