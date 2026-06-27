# PUT `/api/v1/admin/user/{id}/change-password`

## Ringkasan

Super Admin mengganti password user lain (tanpa old password).

## Authentication

Bearer atau cookie — middleware `ADMIN` (`SuperAdmin` only)

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | number | True | Target user ID |

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| new_password | string | True | Password baru |

**Contoh:**

```json
{
  "new_password": "NewSecurePass123!"
}
```

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`. **`data` selalu `null`**.

```json
{
  "success": true,
  "message": "password changed successfully",
  "data": null,
  "error": null
}
```

## Response Error

**403** — bukan SuperAdmin.

**400** — user ID invalid / body invalid.

## Catatan

- Berbeda dari self change password (`PUT /user/change-password`) yang butuh `old_password` dan return 201.

## Frontend

- Service: `userService.admin.changePassword`
- Hook: `useAdminChangePasswordMutation`
- API path: `API_PATH.admin.user.changePassword(id)`
