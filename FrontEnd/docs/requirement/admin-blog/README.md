# Admin — Blog Management

Modul admin: CRUD artikel blog.

**Hak akses:** `SuperAdmin`, `KoorMedcrev`

## Functional Requirements

| ID | Requirement |
| --- | --- |
| BLOG-ADM-01 | List semua blog (termasuk draft) dengan filter kategori |
| BLOG-ADM-02 | Create blog dengan upload gambar / existing gallery ID |
| BLOG-ADM-03 | Update blog |
| BLOG-ADM-04 | Delete blog |
| BLOG-ADM-05 | Kelola status publish |

## Format Respons

Blog admin menggunakan **flat JSON** (tanpa envelope). Lihat [shared.md](../shared.md).

## Status Valid

`draft`, `published`, `unpublished`, `rejected`, `pending_review`

**Tidak ada** status `scheduled`.

## Backend Quirk — Admin List Query

Admin list memakai query param **`kategory[]`** (typo di backend handler), bukan `kategori[]`.

Frontend type: `AdminBlogQuery.kategory?: string[]`

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/admin/blogs` | [GET-admin-blogs-list.md](./GET-admin-blogs-list.md) |
| GET | `/api/v1/admin/blogs/{id}` | [GET-admin-blogs-by-id.md](./GET-admin-blogs-by-id.md) |
| POST | `/api/v1/admin/blogs` | [POST-admin-blogs-create.md](./POST-admin-blogs-create.md) |
| PUT | `/api/v1/admin/blogs/{id}` | [PUT-admin-blogs-by-id.md](./PUT-admin-blogs-by-id.md) |
| DELETE | `/api/v1/admin/blogs/{id}` | [DELETE-admin-blogs-by-id.md](./DELETE-admin-blogs-by-id.md) |

## Types & Hooks

- Types: `lib/types/blog.ts`
- Service: `blogService.admin.*`
- Hooks: `useAdminBlogsQuery`, `useCreateBlogMutation`, `useUpdateBlogMutation`, `useDeleteBlogMutation`
