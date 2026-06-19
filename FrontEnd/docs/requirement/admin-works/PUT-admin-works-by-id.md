# PUT `/api/v1/admin/works/{id}`

## Ringkasan

Update konten work (metadata, teknologi, gambar). Status moderasi di endpoint terpisah (`PUT …/status`).

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`

| Group | Role JWT |
| --- | --- |
| ADMIN | `SuperAdmin` |
| KOOR | `KoorPemro`, `KoorJaringan`, `KoorData`, `KoorMedcrev`, `BPH` |

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Work ID |

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |

## Request Body (Form Data)

Semua field opsional (partial update):

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| pengurus_id | number | False | Owner pengurus |
| title | string | False | Judul |
| tagline | string | False | Tagline |
| description | string | False | Deskripsi |
| slug | string | False | Slug |
| project_type | string | False | Tipe proyek |
| technologies[] | string[] | False | Teknologi |
| project_date | string | False | `YYYY-MM-DD` |
| status | string | False | Status work |
| existingID_image[] | number[] | False | Gallery ID yang dipertahankan/ditambah |
| files | file[] | False | Upload gambar baru |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Work updated successfully",
  "data": {
    "id": 1,
    "pengurus_id": 2,
    "title": "Website DOSCOM v2",
    "status": "pending_review",
    "image_url": "https://...",
    "work_gallery": []
  },
  "error": null
}
```

## Response Error

**400** — ID/body invalid.

**403** — role atau divisi tidak diizinkan.

## Catatan

- Gunakan `existingID_image[]` untuk mempertahankan gambar gallery lama saat menambah file baru.
- Moderasi status (BPH) via `PUT /admin/works/{id}/status`.

## Frontend

- Service: `workService.admin.update`
- Hook: `useUpdateWorkMutation`
- API path: `API_PATH.admin.works.detail(id)`
