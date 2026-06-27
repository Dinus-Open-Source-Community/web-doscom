# GET `/api/v1/user/{id}`

## Ringkasan

Detail user by ID (scoped by role caller).

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`, `BPH`

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
  "message": "Get user",
  "data": {
    "id": 2,
    "username": "koor_pemro",
    "email": "koor@doscom.id",
    "role": "KoorPemro",
    "full_name": "Koordinator Pemro"
  },
  "error": null
}
```

## Response Error

**400** — ID invalid.

**403** — role tidak valid.

**404** — user tidak ditemukan atau di luar scope.

## Frontend

- Service: `userService.getById`
- Hook: `useUserQuery`
- API path: `API_PATH.user.detail(id)`
