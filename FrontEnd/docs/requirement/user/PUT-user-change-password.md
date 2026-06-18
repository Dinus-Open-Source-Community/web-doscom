# PUT `/api/v1/user/change-password`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — ADMIN, KOOR, BPH, PENGURUS

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |
| Authorization | `Bearer {access_token}` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| old_password | string | True | Password lama |
| new_password | string | True | Password baru |

**Contoh:**

```json
{
  "old_password": "lama123",
  "new_password": "baru123456"
}
```

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "password changed successfully, anjayy aku tau pw mu :)",
  "data": null
}
```

## Frontend

- Service: `userService.changePassword`
- Hook: `useChangePasswordMutation`
- API path: `API_PATH` di `lib/api-path.ts`
