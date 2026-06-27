# Auth & Session

Modul: halaman `/admin/login`, session management, refresh token.

## Functional Requirements

| ID | Requirement |
| --- | --- |
| AUTH-01 | User dapat login dengan email & password |
| AUTH-02 | Setelah login sukses, sesi cookie tersimpan dan user diarahkan ke `/admin` |
| AUTH-03 | Request API admin memakai cookie `AccessToken` otomatis (`withCredentials: true`) |
| AUTH-04 | User dapat logout; cookie dihapus dan cache query dibersihkan |
| AUTH-05 | Session refresh via cookie `RefreshToken` (background / saat 401) |
| AUTH-06 | Halaman `/admin/*` (kecuali login) redirect ke login jika tidak autentikasi |

## UI Requirements (Login Page)

| Elemen | Requirement |
| --- | --- |
| Email input | Required, type email |
| Password input | Required, masked |
| Submit | Panggil `useLoginMutation`, disable saat `isPending` |
| Error | Tampilkan `parseApiError(error)` |
| Success | Redirect `ROUTES.admin.dashboard` |

## Non-Functional

- Password minimal 8 karakter (backend register; disarankan sama di UI)
- Jangan log token ke console production
- Cookie `RefreshToken` path `/api/v1/auth`; `AccessToken` path `/`
- Axios wajib `withCredentials: true`

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| POST | `/api/v1/auth/login` | [POST-auth-login.md](./POST-auth-login.md) |
| POST | `/api/v1/auth/register` | [POST-auth-register.md](./POST-auth-register.md) |
| POST | `/api/v1/auth/refresh` | [POST-auth-refresh.md](./POST-auth-refresh.md) |
| POST | `/api/v1/auth/logout` | [POST-auth-logout.md](./POST-auth-logout.md) |

## Response Format

Semua endpoint auth memakai **envelope** `{ success, message, data, error }`. Token **tidak** ada di body — disimpan di HttpOnly cookie.
