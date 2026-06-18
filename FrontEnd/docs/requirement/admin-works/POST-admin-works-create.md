# POST `/api/v1/admin/works`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, Koordinator

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `multipart/form-data` |
| Authorization | `Bearer {access_token}` |

## Form Data

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| pengurus_id | number | True |  |
| title | string | True |  |
| tagline | string | True |  |
| description | string | True |  |
| slug | string | True |  |
| project_type | string | True |  |
| technologies[] | string[] | True |  |
| project_date | string | True | YYYY-MM-DD |
| status | string | True |  |
| division | string | False |  |
| existingID_image | number[] | False | Max 5 gallery |
| files | file[] | False |  |

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Work created successfully",
  "data": {
    "id": 1,
    "title": "Project",
    "image_url": "https://..."
  }
}
```

## Catatan

- Handler memanggil ShouldBindJSON + multipart — kirim sebagai form-data dari frontend.

## Frontend

- Service: `workService.admin.create`
- Hook: `useCreateWorkMutation`
- API path: `API_PATH` di `lib/api-path.ts`
