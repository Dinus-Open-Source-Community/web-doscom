# PUT `/api/v1/pengurus/me`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — PENGURUS, KOOR, BPH, ADMIN

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |
| Authorization | `Bearer {access_token}` |

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
    "id": 1,
    "name": "Updated"
  }
}
```

## Frontend

- Service: `pengurusService.updateMe`
- Hook: `useUpdatePengurusMeMutation`
- API path: `API_PATH` di `lib/api-path.ts`
