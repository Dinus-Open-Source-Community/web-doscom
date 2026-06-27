# DELETE `/api/v1/admin/gallery/{id}`

## Ringkasan

Hapus item gallery beserta file MinIO-nya.

## Authentication

Bearer atau cookie — middleware `SuperAdmin`, `KoorMedcrev`

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Gallery ID |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Successfully delete Gallery, bang nasi padang satu bungkus bang",
  "data": null,
  "error": null
}
```

## Response Error

**403** — role tidak diizinkan.

**404** — gallery tidak ditemukan.

## Frontend

- Service: `galleryService.admin.remove`
- Hook: `useDeleteGalleryMutation`
- API path: `API_PATH.admin.gallery.detail(id)`
