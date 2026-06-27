# POST `/api/v1/admin/works`

## Ringkasan

Buat work baru dengan metadata proyek, gambar upload, dan/atau referensi gallery existing.

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`

| Group | Role JWT |
| --- | --- |
| ADMIN | `SuperAdmin` |
| KOOR | `KoorPemro`, `KoorJaringan`, `KoorData`, `KoorMedcrev`, `BPH` |

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |

## Request Body (Form Data)

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| pengurus_id | number | True | ID pengurus owner |
| title | string | True | Judul proyek |
| tagline | string | True | Tagline singkat |
| description | string | True | Deskripsi lengkap |
| slug | string | True | URL slug unik |
| project_type | string | True | Tipe proyek, contoh: `web` |
| technologies[] | string[] | True | Array teknologi |
| project_date | string | True | Format `YYYY-MM-DD` |
| status | string | True | `draft`, `pending_review`, `published`, dll. |
| division | string | False | Divisi (auto dari role jika kosong) |
| existingID_image[] | number[] | False | ID gallery existing (max 5 total gambar) |
| files | file[] | False | Upload file baru (max 5 total dengan existing) |

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Work created successfully",
  "data": {
    "id": 1,
    "pengurus_id": 2,
    "title": "Website DOSCOM",
    "status": "draft",
    "image_url": "https://...",
    "work_gallery": []
  },
  "error": null
}
```

## Response Error

**400** — validasi form gagal / file error.

**403** — role tidak diizinkan.

## Catatan

- Backend bind via `ShouldBind` multipart — **wajib** kirim `multipart/form-data`, bukan JSON.
- Field gallery existing: `existingID_image[]` (dengan bracket).
- Helper frontend: `buildWorkFormData()` di `lib/func/work.ts`.

## Frontend

- Service: `workService.admin.create`
- Hook: `useCreateWorkMutation`
- API path: `API_PATH.admin.works.list`
