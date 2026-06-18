# Users & Profile

Modul: manajemen user (`/user`) dan halaman **Profile** akun login.

**Hak akses:** `ADMIN`, `KOOR`, `BPH` (CRUD user) · `ADMIN`, `KOOR`, `BPH`, `PENGURUS` (profil sendiri)

Endpoint super admin ada di [admin-user/](../admin-user/README.md).

## Functional Requirements

| ID | Requirement |
| --- | --- |
| USER-04 | Admin/Koor/BPH dapat list user sesuai role hierarchy |
| USER-05 | Admin/Koor/BPH dapat create/update/delete user |
| USER-06 | Tabel user: search, filter role, pagination |
| USER-07 | Form create user: username, email, fullname, role, password |
| PROF-01 | Tampilkan data user login (nama, email, role) |
| PROF-02 | Edit profil user (username, email, fullname) |
| PROF-03 | Ganti password sendiri |
| PROF-06 | Sidebar header menampilkan nama & email real (bukan mock) |

## UI Requirements

### Profile Page

1. **Account** — form user profile + change password
2. Sidebar: ganti mock di `Sidebar.astro` dengan `useMeQuery().data.full_name` dan `.email`

### Validasi Form

- Email format valid
- Password min 8 karakter (jika diinput manual)
- Role harus dari daftar role key valid (`constants.RoleGroup` backend)

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
- Hooks: `hooks/user.ts`
- Service: `userService`
