# POST `/api/v1/pengurus`

## Ringkasan

Buat profil pengurus untuk user yang sedang login (self-service onboarding).

## Authentication

Bearer atau cookie — middleware `PENGURUS`, `KOOR`, `BPH`, `ADMIN`

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |

## Request Body (Form Data)

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| divisi | string | True | `bph`, `pemro`, `jaringan`, `medcrev`, `data` |
| name | string | True | Nama (2–150 karakter) |
| position | string | True | Key dari `ValidPosition`, contoh: `pemrogramanAnggota` |
| start_periode_year | number | True | Tahun awal periode |
| end_periode_year | number | True | Tahun akhir periode |
| sosmed | string[] | False | Max 4 URL sosmed |
| file | file | False | Foto profil |

Field `id_user` di-override otomatis ke user login.

## Response Success

**Status:** `200`

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

**400** — validasi divisi/posisi gagal.

**409** — email sudah terdaftar di pengurus.

## Catatan

- Self-create mengembalikan **200** (bukan 201). Admin create via `POST /admin/pengurus` mengembalikan 201.
- Posisi harus exact key camelCase dari `constants.ValidPosition`.

## Frontend

- Service: `pengurusService.createProfile`
- Hook: `useCreatePengurusProfileMutation`
- API path: `API_PATH.pengurus.createProfile`
