# DOSCOM Frontend — Requirements

Dokumentasi kebutuhan admin panel dan spesifikasi request/response API backend.

Base URL: `{PUBLIC_API_URL}/api/v1` (default dev: `http://localhost:8080/api/v1`)

## Scope Admin Panel

| Route | Halaman | Status UI |
| --- | --- | --- |
| `/admin/login` | Login | Static form (belum integrasi API) |
| `/admin` | Dashboard | Static stat cards |
| `/admin/core-team` | Core Team / Pengurus | Static table (mock data) |

Menu sidebar direncanakan: Dashboard, Super Admin, BPH, Core Team, Blog, Gallery, Profile, Logout.

## Dokumen Umum

| File | Isi |
| --- | --- |
| [shared.md](./shared.md) | Auth header, pagination, form-data, error handling, role-based UI |
| [dashboard.md](./dashboard.md) | Statistik dashboard (belum ada endpoint agregat) |

## Modul — Public & User

| Kategori | Modul | Endpoints |
| --- | --- | --- |
| [auth/](./auth/) | [Auth & Session](./auth/README.md) | 4 |
| [user/](./user/) | [Users & Profile](./user/README.md) | 8 |
| [blog/](./blog/) | [Blog (Public)](./blog/README.md) | 2 |
| [gallery/](./gallery/) | [Gallery (Public)](./gallery/README.md) | 1 |
| [works/](./works/) | [Works (Public)](./works/README.md) | 1 |
| [pengurus/](./pengurus/) | [Pengurus (Public & Self)](./pengurus/README.md) | 5 |
| [upload/](./upload/) | [Upload & Media](./upload/README.md) | 2 |

## Modul — Admin (Hak Akses Terbatas)

| Kategori | Modul | Hak Akses | Endpoints |
| --- | --- | --- | --- |
| [admin-user/](./admin-user/) | [Super Admin](./admin-user/README.md) | `SuperAdmin` | 3 |
| [admin-blog/](./admin-blog/) | [Blog Management](./admin-blog/README.md) | `SuperAdmin`, `KoorMedcrev` | 5 |
| [admin-gallery/](./admin-gallery/) | [Gallery Management](./admin-gallery/README.md) | `SuperAdmin`, `KoorMedcrev` | 2 |
| [admin-works/](./admin-works/) | [Works Management](./admin-works/README.md) | `SuperAdmin`, koordinator | 5 |
| [admin-pengurus/](./admin-pengurus/) | [Core Team](./admin-pengurus/README.md) | `KOOR`, `BPH`, `ADMIN` | 6 |

## Index Endpoint

| Kategori | Method | Path | Dokumen |
| --- | --- | --- | --- |
| `auth` | POST | `/api/v1/auth/login` | [auth/POST-auth-login.md](./auth/POST-auth-login.md) |
| `auth` | POST | `/api/v1/auth/register` | [auth/POST-auth-register.md](./auth/POST-auth-register.md) |
| `auth` | POST | `/api/v1/auth/refresh` | [auth/POST-auth-refresh.md](./auth/POST-auth-refresh.md) |
| `auth` | POST | `/api/v1/auth/logout` | [auth/POST-auth-logout.md](./auth/POST-auth-logout.md) |
| `user` | GET | `/api/v1/user/me` | [user/GET-user-me.md](./user/GET-user-me.md) |
| `user` | PUT | `/api/v1/user/profile` | [user/PUT-user-profile.md](./user/PUT-user-profile.md) |
| `user` | PUT | `/api/v1/user/change-password` | [user/PUT-user-change-password.md](./user/PUT-user-change-password.md) |
| `user` | GET | `/api/v1/user` | [user/GET-user-list.md](./user/GET-user-list.md) |
| `user` | GET | `/api/v1/user/{id}` | [user/GET-user-by-id.md](./user/GET-user-by-id.md) |
| `user` | POST | `/api/v1/user` | [user/POST-user-create.md](./user/POST-user-create.md) |
| `user` | PUT | `/api/v1/user/{id}` | [user/PUT-user-by-id.md](./user/PUT-user-by-id.md) |
| `user` | DELETE | `/api/v1/user/{id}` | [user/DELETE-user-by-id.md](./user/DELETE-user-by-id.md) |
| `admin-user` | GET | `/api/v1/admin/user` | [admin-user/GET-admin-user-list.md](./admin-user/GET-admin-user-list.md) |
| `admin-user` | POST | `/api/v1/admin/user/super-admin` | [admin-user/POST-admin-user-super-admin.md](./admin-user/POST-admin-user-super-admin.md) |
| `admin-user` | PUT | `/api/v1/admin/user/{id}/change-password` | [admin-user/PUT-admin-user-change-password.md](./admin-user/PUT-admin-user-change-password.md) |
| `blog` | GET | `/api/v1/blogs` | [blog/GET-blogs-list.md](./blog/GET-blogs-list.md) |
| `blog` | GET | `/api/v1/blogs/{id}` | [blog/GET-blogs-by-id.md](./blog/GET-blogs-by-id.md) |
| `admin-blog` | GET | `/api/v1/admin/blogs` | [admin-blog/GET-admin-blogs-list.md](./admin-blog/GET-admin-blogs-list.md) |
| `admin-blog` | GET | `/api/v1/admin/blogs/{id}` | [admin-blog/GET-admin-blogs-by-id.md](./admin-blog/GET-admin-blogs-by-id.md) |
| `admin-blog` | POST | `/api/v1/admin/blogs` | [admin-blog/POST-admin-blogs-create.md](./admin-blog/POST-admin-blogs-create.md) |
| `admin-blog` | PUT | `/api/v1/admin/blogs/{id}` | [admin-blog/PUT-admin-blogs-by-id.md](./admin-blog/PUT-admin-blogs-by-id.md) |
| `admin-blog` | DELETE | `/api/v1/admin/blogs/{id}` | [admin-blog/DELETE-admin-blogs-by-id.md](./admin-blog/DELETE-admin-blogs-by-id.md) |
| `gallery` | GET | `/api/v1/gallery` | [gallery/GET-gallery-list.md](./gallery/GET-gallery-list.md) |
| `admin-gallery` | POST | `/api/v1/admin/gallery` | [admin-gallery/POST-admin-gallery-create.md](./admin-gallery/POST-admin-gallery-create.md) |
| `admin-gallery` | DELETE | `/api/v1/admin/gallery/{id}` | [admin-gallery/DELETE-admin-gallery-by-id.md](./admin-gallery/DELETE-admin-gallery-by-id.md) |
| `works` | GET | `/api/v1/works/{projecttype}` | [works/GET-works-by-project-type.md](./works/GET-works-by-project-type.md) |
| `admin-works` | GET | `/api/v1/admin/works` | [admin-works/GET-admin-works-list.md](./admin-works/GET-admin-works-list.md) |
| `admin-works` | GET | `/api/v1/admin/works/{id}` | [admin-works/GET-admin-works-by-id.md](./admin-works/GET-admin-works-by-id.md) |
| `admin-works` | POST | `/api/v1/admin/works` | [admin-works/POST-admin-works-create.md](./admin-works/POST-admin-works-create.md) |
| `admin-works` | PUT | `/api/v1/admin/works/{id}` | [admin-works/PUT-admin-works-by-id.md](./admin-works/PUT-admin-works-by-id.md) |
| `admin-works` | DELETE | `/api/v1/admin/works/{id}` | [admin-works/DELETE-admin-works-by-id.md](./admin-works/DELETE-admin-works-by-id.md) |
| `pengurus` | GET | `/api/v1/pengurus/division/{division}` | [pengurus/GET-pengurus-by-division.md](./pengurus/GET-pengurus-by-division.md) |
| `pengurus` | GET | `/api/v1/pengurus/profile` | [pengurus/GET-pengurus-profile.md](./pengurus/GET-pengurus-profile.md) |
| `pengurus` | POST | `/api/v1/pengurus` | [pengurus/POST-pengurus-create-profile.md](./pengurus/POST-pengurus-create-profile.md) |
| `pengurus` | PUT | `/api/v1/pengurus/me` | [pengurus/PUT-pengurus-me.md](./pengurus/PUT-pengurus-me.md) |
| `pengurus` | DELETE | `/api/v1/pengurus/me` | [pengurus/DELETE-pengurus-me.md](./pengurus/DELETE-pengurus-me.md) |
| `admin-pengurus` | GET | `/api/v1/admin/pengurus` | [admin-pengurus/GET-admin-pengurus-list.md](./admin-pengurus/GET-admin-pengurus-list.md) |
| `admin-pengurus` | POST | `/api/v1/admin/pengurus` | [admin-pengurus/POST-admin-pengurus-create.md](./admin-pengurus/POST-admin-pengurus-create.md) |
| `admin-pengurus` | GET | `/api/v1/admin/pengurus/{id}` | [admin-pengurus/GET-admin-pengurus-by-id.md](./admin-pengurus/GET-admin-pengurus-by-id.md) |
| `admin-pengurus` | GET | `/api/v1/admin/pengurus/by-user/{user_id}` | [admin-pengurus/GET-admin-pengurus-by-user.md](./admin-pengurus/GET-admin-pengurus-by-user.md) |
| `admin-pengurus` | PUT | `/api/v1/admin/pengurus/{id}` | [admin-pengurus/PUT-admin-pengurus-by-id.md](./admin-pengurus/PUT-admin-pengurus-by-id.md) |
| `admin-pengurus` | DELETE | `/api/v1/admin/pengurus/delete/{id}` | [admin-pengurus/DELETE-admin-pengurus-by-id.md](./admin-pengurus/DELETE-admin-pengurus-by-id.md) |
| `upload` | GET | `/api/v1/upload/files` | [upload/GET-upload-files.md](./upload/GET-upload-files.md) |
| `upload` | DELETE | `/api/v1/upload/file` | [upload/DELETE-upload-file.md](./upload/DELETE-upload-file.md) |

## Arsitektur Frontend

```
Pages (Astro/React)
    ↓
hooks/          ← TanStack Query (useXxxQuery, useXxxMutation)
    ↓
services/       ← HTTP calls
    ↓
lib/axios.ts    → Backend API
```

Provider wajib: `plugins/QueryProvider.tsx` (atau `withQueryProvider`).

## Format Respons

### Envelope (mayoritas endpoint domain)

```json
{ "success": true, "message": "...", "data": {}, "error": null }
```

### Flat (auth & sebagian blog/gallery)

JSON langsung tanpa field `success`.

## Checklist Integrasi Admin Panel

- [ ] Login → simpan token → redirect `/admin`
- [ ] Route guard: halaman admin redirect ke login jika tidak ada token
- [ ] Dashboard stats dari API (atau agregasi client-side)
- [ ] Core Team CRUD via pengurus admin API
- [ ] Blog, Gallery admin pages + hooks
- [ ] Profile: `useMeQuery` + `usePengurusProfileQuery`
- [ ] Logout: `useLogoutMutation` + redirect login
- [ ] Toast/notifikasi pakai `lib/message.ts`

## Referensi Kode

| Area | Path |
| --- | --- |
| API paths | `FrontEnd/src/lib/api-path.ts` |
| Types | `FrontEnd/src/lib/types/` |
| Services | `FrontEnd/src/services/` |
| Hooks | `FrontEnd/src/hooks/` |
| Backend routes | `BackEnd/internal/routes/route/` |
| Swagger | `BackEnd/docs/swagger.yaml` |
