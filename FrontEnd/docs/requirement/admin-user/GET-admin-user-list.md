# GET `/api/v1/admin/user`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — SuperAdmin (ADMIN group)

## Response Success

**Status:** `200`

Format envelope `{ success, message, data, error }`.

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
  ]
}
```

## Frontend

- Service: `userService.admin.listSuperAdmin`
- Hook: `useSuperAdminsQuery`
- API path: `API_PATH` di `lib/api-path.ts`
