# DELETE `/api/v1/admin/gallery/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, KoorMedcrev

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True |  |

## Response Success

**Status:** `200`

```json
{
  "message": "Successfully delete Gallery, bang nasi padang satu bungkus bang"
}
```

## Frontend

- Service: `galleryService.admin.remove`
- Hook: `useDeleteGalleryMutation`
- API path: `API_PATH` di `lib/api-path.ts`
