# GET `/api/v1/admin/pengurus/{id}`

## Ringkasan

Detail pengurus by pengurus ID (admin view, termasuk email dan id_user).

## Authentication

Bearer atau cookie — middleware `KOOR`, `BPH`, `ADMIN`

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Pengurus ID |

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
    "position": "koordinatorPemrograman",
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

**404** — tidak ditemukan atau di luar scope divisi.

## Frontend

- Service: `pengurusService.admin.getById`
- Hook: `useAdminPengurusQuery`
- API path: `API_PATH.admin.pengurus.detail(id)`
