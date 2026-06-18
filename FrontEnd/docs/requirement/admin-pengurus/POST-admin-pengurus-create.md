# POST `/api/v1/admin/pengurus`

## Ringkasan

Admin/koordinator/BPH membuat profil pengurus untuk user lain.

## Authentication

Bearer atau cookie — middleware `KOOR`, `BPH`, `ADMIN`

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |

## Request Body (Form Data)

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id_user | number | False | User ID (wajib untuk admin create) |
| email | string | False | Email (auto dari user jika kosong) |
| divisi | string | True | Divisi valid |
| name | string | True | Nama (2–150 karakter) |
| position | string | True | Key `ValidPosition` |
| start_periode_year | number | True | Tahun awal periode |
| end_periode_year | number | True | Tahun akhir periode |
| sosmed | string[] | False | Max 4 URL sosmed |
| file | file | False | Foto profil |

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Pengurus created successfully",
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

**400** — validasi gagal / user tidak ditemukan.

**403** — koordinator tidak bisa create di luar divisi.

**409** — email sudah terdaftar.

## Catatan

- Admin create: status **201**. Self create (`POST /pengurus`): status **200**.
- Koordinator tidak bisa upload foto pengurus anggota (403 dari service).

## Frontend

- Service: `pengurusService.admin.create`
- Hook: `useCreateAdminPengurusMutation`
- API path: `API_PATH.admin.pengurus.list`
