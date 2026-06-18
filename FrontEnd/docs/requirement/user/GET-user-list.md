# GET `/api/v1/user`

## Ringkasan

List semua user yang dapat diakses oleh role caller. **Tidak ada pagination.**

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`, `BPH`

| Group | Role JWT |
| --- | --- |
| ADMIN | `SuperAdmin` |
| KOOR | `KoorPemro`, `KoorJaringan`, `KoorData`, `KoorMedcrev`, `BPH` |
| BPH | `BPH` |

Role `PENGURUS` (**403** forbidden).

## Query Parameters

Tidak ada. Parameter `page`/`limit` **tidak** diproses backend.

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`. `data` adalah **array langsung**, bukan objek paginated.

```json
{
  "success": true,
  "message": "List of users data",
  "data": [
    {
      "id": 2,
      "username": "koor_pemro",
      "email": "koor@doscom.id",
      "role": "KoorPemro",
      "full_name": "Koordinator Pemro"
    }
  ],
  "error": null
}
```

## Response Error

**403** — role `PENGURUS` atau role tidak valid.

## Catatan

- Scope user difilter di service berdasarkan role creator (SuperAdmin lihat lebih luas).
- Frontend `useUsersQuery` masih menerima `PaginationQuery` opsional, tetapi backend mengabaikannya — implementasi client-side filter jika perlu.

## Frontend

- Service: `userService.list`
- Hook: `useUsersQuery`
- API path: `API_PATH.user.list`
