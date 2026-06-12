# GET `/api/v1/user/me`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — ADMIN, KOOR, BPH, PENGURUS

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "current user data",
  "data": {
    "id": 1,
    "username": "admin",
    "email": "admin@doscom.id",
    "role": "SuperAdmin",
    "full_name": "Super Admin"
  }
}
```

## Frontend

- Service: `userService.getMe`
- Hook: `useMeQuery`
- API path: `API_PATH` di `lib/api-path.ts`
