# GET `/api/v1/user/me`

## Ringkasan

Ambil data user yang sedang login.

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`, `BPH`, `PENGURUS`

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
  },
  "error": null
}
```

## Response Error

**403** — role JWT tidak valid.

## Catatan

- Dipakai di seluruh admin panel via `useMeQuery()` untuk role-based UI.
- Token dibaca dari cookie `AccessToken` (utama) atau header Bearer.

## Frontend

- Service: `userService.getMe`
- Hook: `useMeQuery`
- API path: `API_PATH.user.me`
