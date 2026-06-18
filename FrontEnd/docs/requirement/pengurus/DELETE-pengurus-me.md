# DELETE `/api/v1/pengurus/me`

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
  "message": "pengurus deleted",
  "data": null
}
```

## Frontend

- Service: `pengurusService.deleteMe`
- Hook: `useDeletePengurusMeMutation`
- API path: `API_PATH` di `lib/api-path.ts`
