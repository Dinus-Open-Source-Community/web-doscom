# PUT `/api/v1/user/{id}`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — ADMIN, KOOR, BPH

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |
| Authorization | `Bearer {access_token}` |

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True | User ID |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| username | string | False |  |
| email | string | False |  |
| fullname | string | False |  |

**Contoh:**

```json
{
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
    "username": "user",
    "email": "u@example.com",
    "role": "pemroAnggota",
    "full_name": "Updated Name"
  }
}
```

## Frontend

- Service: `userService.update`
- Hook: `useUpdateUserMutation`
- API path: `API_PATH` di `lib/api-path.ts`
