# PUT `/api/v1/admin/pengurus/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — KOOR, BPH, ADMIN

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |
| Authorization | `Bearer {access_token}` |

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True |  |

## Form Data

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| name | string | False |  |
| email | string | False |  |
| divisi | string | False |  |
| position | string | False |  |
| start_periode_year | number | False |  |
| end_periode_year | number | False |  |
| sosmed | string[] | False |  |
| file | file | False |  |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Successfully update pengurus data",
  "data": {
    "id": 1
  }
}
```

## Frontend

- Service: `pengurusService.admin.update`
- Hook: `useUpdateAdminPengurusMutation`
- API path: `API_PATH` di `lib/api-path.ts`
