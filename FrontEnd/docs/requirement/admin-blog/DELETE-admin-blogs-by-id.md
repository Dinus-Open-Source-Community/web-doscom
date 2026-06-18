# DELETE `/api/v1/admin/blogs/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, KoorMedcrev

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True |  |

## Response Success

**Status:** `204`

```json
{
  "message": "successfully delete data"
}
```

## Frontend

- Service: `blogService.admin.remove`
- Hook: `useDeleteBlogMutation`
- API path: `API_PATH` di `lib/api-path.ts`
