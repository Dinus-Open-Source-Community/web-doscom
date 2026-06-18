# GET `/api/v1/admin/blogs`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, KoorMedcrev

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False | Default 1 |
| limit | number | False | Default 10 |
| kategory[] | string[] | False | Typo backend: kategory, max 3 |

## Response Success

**Status:** `200`

```json
{
  "message": "successfully fetch data",
  "blogs": [],
  "totalPage": 1,
  "currentPage": 1
}
```

## Frontend

- Service: `blogService.admin.list`
- Hook: `useAdminBlogsQuery`
- API path: `API_PATH` di `lib/api-path.ts`
