# GET `/api/v1/admin/pengurus`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — KOOR, BPH, ADMIN

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| divisi | string | False | Filter divisi |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "successfully get data",
  "data": []
}
```

## Catatan

- Handler `GetAllPengurusByDivision` ada di backend tetapi route `GET /admin/pengurus` belum terdaftar di `pengurus_route.go` — perlu ditambahkan tim backend.

## Frontend

- Service: `pengurusService.admin.list`
- Hook: `useAdminPengurusListQuery`
- API path: `API_PATH` di `lib/api-path.ts`
