# POST `/api/v1/auth/login`

## Ringkasan

Login admin panel. Token disimpan di HttpOnly cookie, bukan response body.

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

Format envelope. **Tidak ada token di body** — token di-set via cookie.

```json
{
  "success": true,
  "message": "login success bolo, nasi padang satu bungkus",
  "data": null,
  "error": null
}
```

**Cookies yang di-set:**

| Cookie | Path | Durasi | Keterangan |
| --- | --- | --- | --- |
| `AccessToken` | `/` | ~15 menit | JWT access token (HttpOnly) |
| `RefreshToken` | `/api/v1/auth` | ~2 jam | Refresh token (HttpOnly) |

## Response Error

**401** — envelope:

```json
{
  "success": false,
  "message": "Invalid email or password",
  "data": null,
  "error": null
}
```

**401** — role tidak valid:

```json
{
  "success": false,
  "message": "acces denied",
  "data": null,
  "error": null
}
```

## Catatan

- Butuh `withCredentials: true` di axios agar cookie tersimpan.
- Backend juga menerima `Authorization: Bearer` sebagai fallback (opsional).
- Role user harus terdaftar di `authorization.GetRoleInfo`.

## Frontend

- Service: `authService.login`
- Hook: `useLoginMutation`
- API path: `API_PATH.auth.login`
