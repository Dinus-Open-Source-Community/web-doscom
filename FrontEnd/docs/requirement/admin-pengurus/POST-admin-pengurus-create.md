# POST `/api/v1/admin/pengurus`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — KOOR, BPH, ADMIN

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |
| Authorization | `Bearer {access_token}` |

## Form Data

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id_user | number | False | User existing |
| email | string | False | Buat user baru jika kosong |
| divisi | string | True |  |
| name | string | True |  |
| position | string | True |  |
| start_periode_year | number | True |  |
| end_periode_year | number | True |  |
| sosmed | string[] | False |  |
| file | file | False |  |

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Pengurus created successfully",
  "data": {
    "id": 1
  }
}
```

## Frontend

- Service: `pengurusService.admin.create`
- Hook: `useCreateAdminPengurusMutation`
- API path: `API_PATH` di `lib/api-path.ts`
