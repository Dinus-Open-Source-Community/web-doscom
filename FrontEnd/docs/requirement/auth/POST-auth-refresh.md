# POST `/api/v1/auth/refresh`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Cookie RefreshToken

## Response Success

**Status:** `200`

```json
{
  "access_token": "eyJ...",
  "message": "refresh token success"
}
```

## Response Error

**401**

```json
{"message": "Refresh token not found, what are you doing heree????"}
```

## Catatan

- Butuh `withCredentials: true` di axios.

## Frontend

- Service: `authService.refresh`
- Hook: `useRefreshMutation`
- API path: `API_PATH` di `lib/api-path.ts`
