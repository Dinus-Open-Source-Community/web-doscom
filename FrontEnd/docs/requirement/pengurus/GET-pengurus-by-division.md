# GET `/api/v1/pengurus/division/{division}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Public

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| division | string | True | pemro|jaringan|data|medcrev|bph |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "List of pengurus",
  "data": [
    {
      "id": 1,
      "name": "Member",
      "position": "PemroAng",
      "photo_url": "https://...",
      "divisi": "pemro",
      "sosmed": [],
      "start_periode_year": 2025,
      "end_periode_year": 2026
    }
  ]
}
```

## Frontend

- Service: `pengurusService.listByDivision`
- Hook: `usePengurusByDivisionQuery`
- API path: `API_PATH` di `lib/api-path.ts`
