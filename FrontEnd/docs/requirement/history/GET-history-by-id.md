# GET `/api/v1/history/{id}`

## Ringkasan
Detail history/timeline by ID.

## Authentication
Public

## Path Parameters
| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True | History timeline ID |

## Response Success
**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 1,
    "id_author": 1,
    "title": "Sejarah DOSCOM",
    "year": 2020,
    "description": "...",
    "photos": [
      {
        "id": 1,
        "id_history": 1,
        "image_url": "https://..."
      }
    ]
  },
  "error": null
}
```

## Response Error
**Status:** `400` / `404`
