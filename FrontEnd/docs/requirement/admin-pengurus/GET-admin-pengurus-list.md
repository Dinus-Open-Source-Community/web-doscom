# GET `/api/v1/admin/pengurus`

## Ringkasan

List pengurus admin dengan filter divisi opsional.

## Authentication

Bearer atau cookie — middleware `KOOR`, `BPH`, `ADMIN` (jika route terdaftar)

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| divisi | string | False | Filter: `bph`, `pemro`, `jaringan`, `medcrev`, `data` |

## Response Success (Expected)

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "successfully get data",
  "data": [
    {
      "id": 1,
      "id_user": 5,
      "photo_url": "https://...",
      "email": "user@doscom.id",
      "divisi": "pemro",
      "name": "John Doe",
      "position": "koordinatorPemrograman",
      "sosmed": [],
      "start_periode_year": 2024,
      "end_periode_year": 2025,
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": "2025-01-01T00:00:00Z"
    }
  ],
  "error": null
}
```

## Backend Gap

| Item | Status |
| --- | --- |
| Handler | `GetAllPengurusByDivision` — **implemented** |
| Route registration | `GET /admin/pengurus` — **NOT registered** in `pengurus_route.go` |

Saat ini request ke path ini akan mengembalikan **404** dari router Gin.

## Catatan

- Koordinator hanya melihat pengurus divisi sendiri (logic di service).
- Frontend hook `useAdminPengurusListQuery` sudah ada — **blocked** sampai backend menambahkan route.
- Workaround dashboard: agregasi `GET /pengurus/division/{division}` per divisi (public, tanpa email).

## Frontend

- Service: `pengurusService.admin.list` (blocked)
- Hook: `useAdminPengurusListQuery` (blocked)
- API path: `API_PATH.admin.pengurus.list`
