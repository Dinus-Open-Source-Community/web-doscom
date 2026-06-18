# PUT `/api/v1/admin/works/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, Koordinator

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
| title | string | False |  |
| tagline | string | False |  |
| description | string | False |  |
| slug | string | False |  |
| project_type | string | False |  |
| technologies[] | string[] | False |  |
| project_date | string | False |  |
| status | string | False |  |
| existingID_image | number[] | False |  |
| files | file[] | False |  |

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Work updated successfully",
  "data": {
    "id": 1,
    "title": "Updated"
  }
}
```

## Frontend

- Service: `workService.admin.update`
- Hook: `useUpdateWorkMutation`
- API path: `API_PATH` di `lib/api-path.ts`
