# GET `/api/v1/admin/pengurus/by-user/{user_id}`

## Ringkasan

Lookup profil pengurus by user ID (berguna saat link work/blog ke owner).

## Authentication

Bearer atau cookie — middleware `KOOR`, `BPH`, `ADMIN`

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| user_id | number | True | User ID |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Successfully get data",
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

**400** — user_id invalid.

**404** — pengurus tidak ditemukan untuk user tersebut.

## Frontend

- Service: `pengurusService.admin.getByUserId`
- Hook: `useAdminPengurusByUserQuery`
- API path: `API_PATH.admin.pengurus.byUser(userId)`
