# GET `/api/v1/history`

## Ringkasan
List history/timeline publik dengan paginasi.

## Authentication
Public

## Query Parameters
| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| page | number | False | Default `1` |
| limit | number | False | Default `10` |

## Response Success
**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "message": "Success",
    "history": [
      {
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
      }
    ],
    "totalPage": 1,
    "currentPage": 1
  },
  "error": null
}
```
