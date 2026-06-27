# Super Admin — User Management

Modul admin: sidebar **Super Admin** — manajemen akun super admin.

**Hak akses:** middleware `ADMIN` (`SuperAdmin` only)

## Functional Requirements

| ID | Requirement |
| --- | --- |
| SA-01 | Super Admin dapat melihat daftar super admin |
| SA-02 | Super Admin dapat membuat akun super admin baru |
| SA-03 | Super Admin dapat mengganti password user lain |

## UI Requirements

| Kolom tabel | Sumber |
| --- | --- |
| ID | `user.id` |
| Username | `user.username` |
| Email | `user.email` |
| Role | `user.role` |
| Full name | `user.full_name` |
| Actions | Change password |

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/admin/user` | [GET-admin-user-list.md](./GET-admin-user-list.md) |
| POST | `/api/v1/admin/user/super-admin` | [POST-admin-user-super-admin.md](./POST-admin-user-super-admin.md) |
| PUT | `/api/v1/admin/user/{id}/change-password` | [PUT-admin-user-change-password.md](./PUT-admin-user-change-password.md) |

## Types & Hooks

- Types: `lib/types/user.ts`
- Hooks: `hooks/user.ts` — `useSuperAdminsQuery`, `useCreateSuperAdminMutation`, `useAdminChangePasswordMutation`
- Service: `userService.admin.*`
