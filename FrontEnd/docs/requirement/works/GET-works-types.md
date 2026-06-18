# GET `/api/v1/works/types`

## Ringkasan

Daftar tipe proyek yang tersedia di database.

## Authentication

Public

## Response Success

**Status:** `200`

```json
{
  "success": true,
  "message": "Successfully fetch work types",
  "data": ["web", "mobile", "design"],
  "error": null
}
```

## Catatan

- Route terdaftar setelah `GET /works/:id` di `work_route.go` — request ke `/works/types` bisa tertangkap sebagai `:id=types`. Perlu perbaikan urutan route di backend.

## Frontend

- Service: `workService.getTypes`
- Hook: `useWorkTypesQuery`
- API path: `API_PATH.works.types`
