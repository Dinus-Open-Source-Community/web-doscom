# POST `/api/v1/auth/register`

## Ringkasan

Dokumentasi request/response endpoint backend.

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
| email | string | True | Email |
| password | string | True | Min 8 karakter |
| role | string | False | Role key, contoh: pemroAnggota |
| fullname | string | True | Nama lengkap |

**Contoh:**

```json
{
  "username": "johndoe",
  "email": "user@mhs.dinus.ac.id",
  "password": "password123",
  "role": "pemroAnggota",
  "fullname": "John Doe"
}
```

## Response Success

**Status:** `200`

```json
{
  "message": "user created successfully"
}
```

## Response Error

**400**

```json
{"error": "Invalid request body or missing fields"}
```

**500**

```json
{"error": "failed to register user"}
```

## Frontend

- Service: `authService.register`
- Hook: `useRegisterMutation`
- API path: `API_PATH` di `lib/api-path.ts`
