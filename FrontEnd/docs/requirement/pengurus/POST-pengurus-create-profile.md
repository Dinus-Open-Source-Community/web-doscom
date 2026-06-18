# POST `/api/v1/pengurus`

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
| divisi | string | True |  |
| name | string | True |  |
| position | string | True |  |
| start_periode_year | number | True |  |
| end_periode_year | number | True |  |
| email | string | False |  |
| sosmed | string[] | False | Max 4 URL |
| file | file | False | Foto profil |

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Pengurus created successfully",
  "data": {
    "id": 1,
    "name": "Member"
  }
}
```

## Frontend

- Service: `pengurusService.createProfile`
- Hook: `useCreatePengurusProfileMutation`
- API path: `API_PATH` di `lib/api-path.ts`
