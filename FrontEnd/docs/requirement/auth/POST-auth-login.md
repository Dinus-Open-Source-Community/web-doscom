# POST `/api/v1/auth/login`

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
| email | string | True | Email login |
| password | string | True | Password |

**Contoh:**

```json
{
  "email": "111202412345@mhs.dinus.ac.id",
  "password": "password123"
}
```

## Response Success

**Status:** `200`

```json
{
  "acces_token": "eyJhbGciOiJIUzI1NiIs...",
  "message:": "login success bolo, nasi padang satu bungkus"
}
```

## Response Error

**401**

```json
{"error": "Invalid email or password"}
```

**401**

```json
{"error": "acces denied"}
```

## Catatan

- Field token backend: `acces_token` (typo, bukan `access_token`).
- Set cookie HttpOnly `RefreshToken`, path `/api/v1/auth`, max-age 5 hari.

## Frontend

- Service: `authService.login`
- Hook: `useLoginMutation`
- API path: `API_PATH` di `lib/api-path.ts`
