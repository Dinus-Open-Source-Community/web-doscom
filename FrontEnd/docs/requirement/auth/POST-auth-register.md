# POST `/api/v1/auth/register`

## Ringkasan

Registrasi user baru (endpoint publik).

## Authentication

Public

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| username | string | False | Username |
| email | string | True | Email valid |
| password | string | True | Min 8 karakter |
| role | string | True | Role key |
| fullname | string | True | Nama lengkap |

**Contoh:**

```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "password123",
  "role": "pemroAnggota",
  "fullname": "John Doe"
}
```

## Response Success

**Status:** `200`

```json
{
  "success": true,
  "message": "user created successfully",
  "data": null,
  "error": null
}
```

## Response Error

**400** — field kosong / password < 8:

```json
{
  "success": false,
  "message": "Invalid request body or missing fields",
  "data": null,
  "error": null
}
```

## Frontend

- Service: `authService.register`
- Hook: `useRegisterMutation`
- API path: `API_PATH.auth.register`
