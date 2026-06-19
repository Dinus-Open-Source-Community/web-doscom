# DELETE `/api/v1/admin/blogs/{id}`

## Ringkasan

Hapus blog by ID.

## Authentication

Bearer atau cookie — middleware `SuperAdmin`, `KoorMedcrev`

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Blog ID |

## Response Success

**Status:** `204`

Format **flat** — backend mengirim JSON body meski status 204:

```json
{
  "message": "successfully delete data"
}
```

## Response Error

**403** — role tidak diizinkan.

**400** — ID invalid.

**500** — gagal hapus.

## Catatan

- Frontend `blogService.admin.remove` unwrap `response.data`.
- Invalidate cache list dan detail setelah delete.

## Frontend

- Service: `blogService.admin.remove`
- Hook: `useDeleteBlogMutation`
- API path: `API_PATH.admin.blogs.detail(id)`
