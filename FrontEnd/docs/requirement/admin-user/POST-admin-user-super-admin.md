# POST `/api/v1/admin/user/super-admin`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |
| Authorization | `Bearer {access_token}` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| email | string | True |  |
| password | string | True |  |
| fullname | string | True |  |
| username | string | False |  |

**Contoh:**

```json
{
  "email": "super@doscom.id",
  "password": "securePass123",
  "fullname": "Super Admin Baru"
}
```

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Superadmin created successfully",
  "data": null
}
```

## Frontend

- Service: `userService.admin.createSuperAdmin`
- Hook: `useCreateSuperAdminMutation`
- API path: `API_PATH` di `lib/api-path.ts`
