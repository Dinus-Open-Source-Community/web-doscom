# PUT `/api/v1/user/{id}`

## Ringkasan

Update data user by ID (partial update).

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`, `BPH`

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | User ID |

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| username | string | False | Username |
| email | string | False | Email |
| fullname | string | False | Nama lengkap |

**Contoh:**

```json
{
  "email": "updated@doscom.id",
  "fullname": "Updated Name"
}
```

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Successfully update user data",
  "data": {
    "id": 2,
    "username": "koor_pemro",
    "email": "updated@doscom.id",
    "role": "KoorPemro",
    "full_name": "Updated Name"
  },
  "error": null
}
```

## Response Error

**400** — ID/body invalid.

**403** — role tidak valid.

## Frontend

- Service: `userService.update`
- Hook: `useUpdateUserMutation`
- API path: `API_PATH.user.detail(id)`
