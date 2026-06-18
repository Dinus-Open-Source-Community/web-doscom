# POST `/api/v1/auth/logout`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Cookie RefreshToken

## Response Success

**Status:** `200`

```json
{
  "message": "Logout Success, nasi padang satu bungkus"
}
```

## Frontend

- Service: `authService.logout`
- Hook: `useLogoutMutation`
- API path: `API_PATH` di `lib/api-path.ts`
