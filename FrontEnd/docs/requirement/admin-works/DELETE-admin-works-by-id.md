# DELETE `/api/v1/admin/works/{id}`

## Ringkasan

Hapus work beserta relasi gallery-nya.

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

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`. **`data` bernilai `null`** (bukan HTTP 204).

```json
{
  "success": true,
  "message": "Work deleted successfully",
  "data": null,
  "error": null
}
```

## Response Error

**400** — ID invalid.

**403** — role tidak diizinkan.

**500** — gagal hapus di server.

## Catatan

- Koordinator hanya bisa menghapus work divisi sendiri.
- Frontend service mengembalikan `null` setelah unwrap envelope.

## Frontend

- Service: `workService.admin.remove`
- Hook: `useDeleteWorkMutation`
- API path: `API_PATH.admin.works.detail(id)`
