# POST `/api/v1/user`

## Ringkasan

Dokumentasi request/response endpoint backend.

## Authentication

Bearer — ADMIN, KOOR, BPH

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |
| Authorization | `Bearer {access_token}` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| username | string | False | Opsional, backend bisa assign default |
| email | string | True | Email |
| password | string | False | Opsional |
| role | string | False | Role key |
| fullname | string | True | Nama lengkap |

**Contoh:**

```json
{
  "email": "new@mhs.dinus.ac.id",
  "fullname": "Member Baru",
  "role": "pemroAnggota"
}
```

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "User created successfully",
  "data": null
}
```

## Frontend

- Service: `userService.create`
- Hook: `useCreateUserMutation`
- API path: `API_PATH` di `lib/api-path.ts`
