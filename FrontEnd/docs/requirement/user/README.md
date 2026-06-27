# Users & Profile

Modul: manajemen user dan profil akun login.

**Hak akses umum:** middleware `ADMIN`, `KOOR`, `BPH`, `PENGURUS` (endpoint `/me`, `/profile`, `/change-password`)

**Hak akses CRUD user:** middleware `ADMIN`, `KOOR`, `BPH` (role `PENGURUS` ditolak)

## Functional Requirements

| ID | Requirement |
| --- | --- |
| USER-01 | User login dapat melihat profil sendiri (`GET /user/me`) |
| USER-02 | User dapat update profil sendiri |
| USER-03 | User dapat ganti password sendiri |
| USER-04 | Admin/koordinator/BPH dapat list user (tanpa pagination) |
| USER-05 | Admin/koordinator/BPH dapat CRUD user sesuai scope role |
| USER-06 | List user difilter by role creator (backend enforce) |

## UI Requirements

| Halaman | Endpoint |
| --- | --- |
| Profile admin | `useMeQuery`, `usePengurusProfileQuery` |
| User management | `useUsersQuery`, `useCreateUserMutation` |
| Change password modal | `useChangePasswordMutation` |

## Catatan Penting

- `GET /user` **tidak** mendukung pagination — mengembalikan seluruh array user sesuai role.
- Auth via cookie `AccessToken` atau header `Authorization: Bearer`.
- Response format envelope untuk semua endpoint modul ini.

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/user/me` | [GET-user-me.md](./GET-user-me.md) |
| PUT | `/api/v1/user/profile` | [PUT-user-profile.md](./PUT-user-profile.md) |
| PUT | `/api/v1/user/change-password` | [PUT-user-change-password.md](./PUT-user-change-password.md) |
| GET | `/api/v1/user` | [GET-user-list.md](./GET-user-list.md) |
| GET | `/api/v1/user/{id}` | [GET-user-by-id.md](./GET-user-by-id.md) |
| POST | `/api/v1/user` | [POST-user-create.md](./POST-user-create.md) |
| PUT | `/api/v1/user/{id}` | [PUT-user-by-id.md](./PUT-user-by-id.md) |
| DELETE | `/api/v1/user/{id}` | [DELETE-user-by-id.md](./DELETE-user-by-id.md) |

## Types & Hooks

- Types: `lib/types/user.ts`
- Service: `userService.*`
- Hooks: `hooks/user.ts` — `useMeQuery`, `useUsersQuery`, `useUpdateProfileMutation`, dll.
