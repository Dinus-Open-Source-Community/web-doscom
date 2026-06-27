# PUT `/api/v1/user/profile`

## Ringkasan

Update profil user yang sedang login (partial update).

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`, `BPH`, `PENGURUS`

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| username | string | False | Username |
| email | string | False | Email |
| fullname | string | False | Nama lengkap (map ke `full_name` di DB) |

**Contoh:**

```json
{
  "username": "john_doe",
  "email": "john@doscom.id",
  "fullname": "John Doe"
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
    "username": "john_doe",
    "email": "john@doscom.id",
    "role": "KoorPemro",
    "full_name": "John Doe"
  },
  "error": null
}
```

## Response Error

**400** — body invalid.

## Catatan

- Field request `fullname` → disimpan sebagai `full_name` di response.
- Hanya field non-kosong yang di-update (`UserPatch.ToMap()`).

## Frontend

- Service: `userService.updateProfile`
- Hook: `useUpdateProfileMutation`
- API path: `API_PATH.user.profile`
