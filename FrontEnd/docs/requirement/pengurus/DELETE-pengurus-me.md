# DELETE `/api/v1/pengurus/me`

## Ringkasan

Hapus profil pengurus user yang sedang login.

## Authentication

Bearer atau cookie — middleware `PENGURUS`, `KOOR`, `BPH`, `ADMIN`

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "pengurus deleted",
  "data": null,
  "error": null
}
```

## Response Error

**404** — profil tidak ditemukan.

## Frontend

- Service: `pengurusService.deleteMe`
- Hook: `useDeletePengurusMeMutation`
- API path: `API_PATH.pengurus.me`
