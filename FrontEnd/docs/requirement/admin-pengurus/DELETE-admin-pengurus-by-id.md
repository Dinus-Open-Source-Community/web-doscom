# DELETE `/api/v1/admin/pengurus/delete/{id}`

## Ringkasan

Hapus pengurus by ID. Path menggunakan prefix `/delete/` (bukan REST standar).

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
  "message": "pengurus deleted",
  "data": null,
  "error": null
}
```

## Response Error

**404** — pengurus tidak ditemukan.

## Catatan

- Path delete: `/admin/pengurus/delete/{id}` (bukan `DELETE /admin/pengurus/{id}`).
- Koordinator hanya bisa hapus pengurus divisi sendiri.

## Frontend

- Service: `pengurusService.admin.remove`
- Hook: `useDeleteAdminPengurusMutation`
- API path: `API_PATH.admin.pengurus.delete(id)`
