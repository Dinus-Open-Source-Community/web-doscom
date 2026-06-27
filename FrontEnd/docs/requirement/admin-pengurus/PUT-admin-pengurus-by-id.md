# PUT `/api/v1/admin/pengurus/{id}`

## Ringkasan

Update data pengurus by ID (admin/koordinator). Partial update via form-data.

## Authentication

Bearer atau cookie — middleware `KOOR`, `BPH`, `ADMIN`

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Pengurus ID |

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |

## Request Body (Form Data)

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| name | string | False | Nama |
| email | string | False | Email |
| divisi | string | False | Divisi valid |
| position | string | False | Key `ValidPosition` |
| start_periode_year | number | False | Tahun awal periode |
| end_periode_year | number | False | Tahun akhir periode |
| sosmed | string[] | False | Max 4 URL |
| file | file | False | Foto profil (SuperAdmin only untuk foto anggota) |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Successfully update pengurus data",
  "data": {
    "id": 1,
    "id_user": 5,
    "photo_url": "https://...",
    "email": "user@doscom.id",
    "divisi": "pemro",
    "name": "John Updated",
    "position": "koordinatorPemrograman",
    "sosmed": [],
    "start_periode_year": 2024,
    "end_periode_year": 2026,
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-06-01T00:00:00Z"
  },
  "error": null
}
```

## Response Error

**403** — koordinator update di luar scope / foto anggota.

**404** — pengurus tidak ditemukan.

## Catatan

- Permission field per role: `RoleFieldPermission` di `constants.go`.
- Koordinator tidak bisa update `position` pengurus di divisi lain.

## Frontend

- Service: `pengurusService.admin.update`
- Hook: `useUpdateAdminPengurusMutation`
- API path: `API_PATH.admin.pengurus.detail(id)`
