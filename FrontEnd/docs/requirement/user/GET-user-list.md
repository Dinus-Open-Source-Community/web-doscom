# GET `/api/v1/user`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — ADMIN, KOOR, BPH

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False | Default 1 |
| limit | number | False | Default 10 |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "List of users data",
  "data": [
    {
      "id": 2,
      "username": "koor",
      "email": "koor@example.com",
      "role": "KoorPemro",
      "full_name": "Koor PEMRO"
    }
  ]
}
```

## Frontend

- Service: `userService.list`
- Hook: `useUsersQuery`
- API path: `API_PATH` di `lib/api-path.ts`
