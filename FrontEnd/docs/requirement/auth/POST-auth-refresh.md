# POST `/api/v1/auth/refresh`

## Ringkasan

Perpanjang sesi dengan refresh token dari cookie.

## Authentication

Cookie `RefreshToken` (path `/api/v1/auth`)

## Response Success

**Status:** `200`

Format envelope. Token baru di-set via cookie `AccessToken`.

```json
{
  "success": true,
  "message": "refresh token success",
  "data": null,
  "error": null
}
```

## Response Error

**401** — envelope:

```json
{
  "success": false,
  "message": "Refresh token not found, what are you doing heree????",
  "data": null,
  "error": null
}
```

**401** — token invalid/expired:

```json
{
  "success": false,
  "message": "refresh token invalid or expired",
  "data": null,
  "error": "..."
}
```

## Catatan

- Butuh `withCredentials: true` di axios.
- Cookie `RefreshToken` di-clear jika invalid/expired.

## Frontend

- Service: `authService.refresh`
- Hook: `useRefreshMutation`
- API path: `API_PATH.auth.refresh`
