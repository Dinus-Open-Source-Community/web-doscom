# GET `/api/v1/pengurus/profile`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — PENGURUS, KOOR, BPH, ADMIN

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "success",
  "data": {
    "id": 1,
    "id_user": 5,
    "name": "Member",
    "email": "m@example.com",
    "divisi": "pemro",
    "position": "PemroAng",
    "photo_url": "https://..."
  }
}
```

## Frontend

- Service: `pengurusService.getProfile`
- Hook: `usePengurusProfileQuery`
- API path: `API_PATH` di `lib/api-path.ts`
