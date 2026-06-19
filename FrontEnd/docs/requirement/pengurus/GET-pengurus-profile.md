# GET `/api/v1/pengurus/profile`

## Ringkasan

Ambil profil pengurus user yang sedang login.

## Authentication

Bearer atau cookie — middleware `PENGURUS`, `KOOR`, `BPH`, `ADMIN`

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
    "photo_url": "https://...",
    "email": "user@doscom.id",
    "divisi": "pemro",
    "name": "John Doe",
    "position": "pemrogramanAnggota",
    "sosmed": [],
    "start_periode_year": 2024,
    "end_periode_year": 2025,
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  },
  "error": null
}
```

## Response Error

**500** — profil belum ada / gagal fetch.

## Catatan

- Dipakai bersama `useMeQuery()` di halaman Profile admin.
- Lookup by `user_id` dari JWT, bukan path param.

## Frontend

- Service: `pengurusService.getProfile`
- Hook: `usePengurusProfileQuery`
- API path: `API_PATH.pengurus.profile`
