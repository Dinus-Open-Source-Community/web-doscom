# PUT `/api/v1/admin/blogs/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin, KoorMedcrev

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
| slug | string | False |  |
| content | string | False |  |
| kategori | string[] | False |  |
| status | string | False |  |
| published_at | datetime | False |  |
| existingID_image | number[] | False |  |
| files | file[] | False |  |

## Response Success

**Status:** `200`

```json
{
  "message": "successfully update blog",
  "data": {
    "id": 1,
    "title": "Updated"
  }
}
```

## Frontend

- Service: `blogService.admin.update`
- Hook: `useUpdateBlogMutation`
- API path: `API_PATH` di `lib/api-path.ts`
