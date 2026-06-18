# GET `/api/v1/admin/user`

## Ringkasan

List akun super admin (hanya user dengan role `SuperAdmin`).

## Authentication

Bearer atau cookie — middleware `ADMIN` (`SuperAdmin` only)

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`. `data` adalah array user.

```json
{
  "success": true,
  "message": "Get super admin data",
  "data": [
    {
      "id": 1,
      "username": "admin",
      "email": "admin@doscom.id",
      "role": "SuperAdmin",
      "full_name": "Super Admin"
    }
  ],
  "error": null
}
```

## Response Error

**403** — bukan SuperAdmin.

## Catatan

- Tidak ada pagination — array langsung di `data`.
- Service exclude caller sendiri dari list (logic di `GetSuperAdmin` service).

## Frontend

- Service: `userService.admin.listSuperAdmin`
- Hook: `useSuperAdminsQuery`
- API path: `API_PATH.admin.user.list`
