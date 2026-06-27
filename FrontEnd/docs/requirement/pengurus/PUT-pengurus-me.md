# PUT `/api/v1/pengurus/me`

## Ringkasan

Update profil pengurus user yang sedang login (partial update + optional foto).

## Authentication

Bearer atau cookie — middleware `PENGURUS`, `KOOR`, `BPH`, `ADMIN`

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
| file | file | False | Foto profil baru |

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
    "position": "pemrogramanAnggota",
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

**404** — profil pengurus belum ada.

## Catatan

- Wajib kirim `multipart/form-data` meski tanpa file.
- Role pengurus anggota tidak bisa ubah `position` (enforced di service).

## Frontend

- Service: `pengurusService.updateMe`
- Hook: `useUpdatePengurusMeMutation`
- API path: `API_PATH.pengurus.me`
