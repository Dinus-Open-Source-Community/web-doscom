# GET `/api/v1/pengurus/division/{division}`

## Ringkasan

List pengurus publik per divisi (untuk halaman Core Team website).

## Authentication

Public

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| division | string | True | Divisi: `bph`, `pemro`, `jaringan`, `medcrev`, `data` |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`. `data` adalah array.

```json
{
  "success": true,
  "message": "List of pengurus",
  "data": [
    {
      "id": 1,
      "photo_url": "https://...",
      "divisi": "pemro",
      "name": "John Doe",
      "position": "koordinatorPemrograman",
      "sosmed": [
        {
          "platform": "instagram",
          "username": "@johndoe",
          "url": "https://instagram.com/johndoe",
          "is_primary": true
        }
      ],
      "start_periode_year": 2024,
      "end_periode_year": 2025
    }
  ],
  "error": null
}
```

## Catatan

- Response public tidak expose `email`, `id_user`.
- Field periode: `start_periode_year`, `end_periode_year` (int).

## Frontend

- Service: `pengurusService.listByDivision`
- Hook: `usePengurusByDivisionQuery`
- API path: `API_PATH.pengurus.byDivision(division)`
