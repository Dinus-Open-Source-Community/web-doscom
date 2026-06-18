# GET `/api/v1/user/{id}`

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
  "message": "Get user",
  "data": {
    "id": 2,
    "username": "user",
    "email": "u@example.com",
    "role": "pemroAnggota",
    "full_name": "User Name"
  }
}
```

## Frontend

- Service: `userService.getById`
- Hook: `useUserQuery`
- API path: `API_PATH` di `lib/api-path.ts`
