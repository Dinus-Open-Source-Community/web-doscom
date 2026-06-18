# POST `/api/v1/auth/logout`

## Ringkasan

Logout user dan hapus refresh token dari database serta cookie.

## Authentication

Cookie `RefreshToken`

## Response Success

**Status:** `200`

```json
{
  "success": true,
  "message": "Logout Success, nasi padang satu bungkus",
  "data": null,
  "error": null
}
```

Kedua cookie (`AccessToken`, `RefreshToken`) di-clear.

## Response Error

**401** — cookie tidak ditemukan:

```json
{
  "success": false,
  "message": "Cookie not found, what are you doing here????",
  "data": null,
  "error": null
}
```

## Catatan

- Butuh `withCredentials: true`.
- Frontend memanggil `clearAccessToken()` untuk legacy localStorage (opsional).

## Frontend

- Service: `authService.logout`
- Hook: `useLogoutMutation`
- API path: `API_PATH.auth.logout`
