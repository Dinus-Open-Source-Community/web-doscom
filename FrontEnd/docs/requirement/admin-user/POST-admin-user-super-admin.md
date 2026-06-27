# POST `/api/v1/admin/user/super-admin`

## Ringkasan

Buat akun super admin baru.

## Authentication

Bearer atau cookie — middleware `ADMIN` (`SuperAdmin` only)

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| email | string | True | Email unik |
| fullname | string | True | Nama lengkap |
| password | string | True | Password (min 8, aturan kompleksitas) |

**Contoh:**

```json
{
  "email": "newadmin@doscom.id",
  "fullname": "New Super Admin",
  "password": "SecurePass123!"
}
```

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "Superadmin created successfully",
  "data": null,
  "error": null
}
```

## Response Error

**403** — bukan SuperAdmin.

**409** — email sudah terdaftar.

**400** — field wajib kosong.

## Frontend

- Service: `userService.admin.createSuperAdmin`
- Hook: `useCreateSuperAdminMutation`
- API path: `API_PATH.admin.user.superAdmin`
