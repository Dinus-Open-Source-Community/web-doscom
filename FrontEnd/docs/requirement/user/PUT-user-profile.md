# PUT `/api/v1/user/profile`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — ADMIN, KOOR, BPH, PENGURUS

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |
| Authorization | `Bearer {access_token}` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| username | string | False | Username baru |
| email | string | False | Email baru |
| fullname | string | False | Nama lengkap |

**Contoh:**

```json
{
  "username": "newname",
  "email": "new@mhs.dinus.ac.id",
  "fullname": "Nama Baru"
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
    "id": 1,
    "username": "newname",
    "email": "new@mhs.dinus.ac.id",
    "role": "KoorPemro",
    "full_name": "Nama Baru"
  }
}
```

## Frontend

- Service: `userService.updateProfile`
- Hook: `useUpdateProfileMutation`
- API path: `API_PATH` di `lib/api-path.ts`
