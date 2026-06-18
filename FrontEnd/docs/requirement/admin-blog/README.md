# Admin — Blog Management

Modul admin: sidebar **Blog** — CRUD artikel.

**Hak akses:** `SuperAdmin`, `KoorMedcrev`

## Functional Requirements

| ID | Requirement |
| --- | --- |
| BLOG-01 | List blog dengan pagination & filter kategori |
| BLOG-02 | Create blog dengan rich content + gambar |
| BLOG-03 | Edit blog (termasuk ganti/reuse gambar existing) |
| BLOG-04 | Delete blog |
| BLOG-05 | Preview thumbnail & kategori di tabel |
| BLOG-06 | Status: draft, published, scheduled, dll. |
| BLOG-07 | Max 3 kategori per blog |

## UI Requirements

| Elemen | Requirement |
| --- | --- |
| Tabel | Kolom: thumbnail, title, kategori, status, published_at, actions |
| Editor | Rich text untuk `content` (Tiptap/similar — TBD) |
| Image picker | Upload baru + pilih dari gallery existing (`existingID_image`) |
| Slug | Auto-generate dari title, editable |
| Validasi | Max 3 kategori, slug required |

## Status Values (Backend)

`draft`, `published`, `scheduled`, `unpublished`, `rejected`, `pending_review`

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
- Hooks: `hooks/blog.ts` — `useAdminBlogsQuery`, `useCreateBlogMutation`, dll.
- Service: `blogService.admin.*`
