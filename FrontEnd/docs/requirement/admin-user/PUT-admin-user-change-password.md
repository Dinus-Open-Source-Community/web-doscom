# PUT `/api/v1/admin/user/{id}/change-password`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |
| Authorization | `Bearer {access_token}` |

## Path Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| id | integer | True | Target user ID |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| new_password | string | True | Password baru |

**Contoh:**

```json
{
  "new_password": "newSecurePass123"
}
```

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "password changed successfully",
  "data": {
    "id": 2,
    "username": "user",
    "email": "u@example.com",
    "role": "KoorPemro",
    "full_name": "User"
  }
}
```

## Frontend

- Service: `userService.admin.changePassword`
- Hook: `useAdminChangePasswordMutation`
- API path: `API_PATH` di `lib/api-path.ts`
