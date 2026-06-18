# PUT `/api/v1/admin/works/{id}/status`

## Ringkasan

Update status work untuk moderasi (BPH / SuperAdmin).

## Authentication

Bearer atau cookie — middleware `ADMIN`, `BPH`

| Group | Role JWT |
| --- | --- |
| ADMIN | `SuperAdmin` |
| BPH | `BPH` |

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Work ID |

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| status | string | True | `published`, `rejected`, `unpublished`, dll. |

**Contoh:**

```json
{
  "status": "published"
}
```

## Response Success

**Status:** `200`

```json
{
  "success": true,
  "message": "Work updated successfully",
  "data": {
    "id": 1,
    "pengurus_id": 2,
    "title": "Project",
    "status": "published",
    "work_gallery": []
  },
  "error": null
}
```

## Response Error

**403** — role tidak diizinkan.

**400** — body invalid.

## Catatan

- Endpoint terpisah dari `PUT /admin/works/{id}` (update konten).
- Gunakan untuk workflow approval BPH.

## Frontend

- Service: `workService.admin.updateStatus`
- Hook: `useUpdateWorkStatusMutation`
- API path: `API_PATH.admin.works.status(id)`
