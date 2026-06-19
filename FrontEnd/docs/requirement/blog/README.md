# Blog (Public)

Modul: artikel blog publik website DOSCOM.

## Functional Requirements

| ID | Requirement |
| --- | --- |
| BLOG-01 | List blog published dengan pagination |
| BLOG-02 | Filter by kategori (max 3) |
| BLOG-03 | Detail blog by ID |

## Format Respons

Blog menggunakan **flat JSON** (tanpa envelope `success`). Lihat [shared.md](../shared.md).

## Status Blog

Valid: `draft`, `published`, `unpublished`, `rejected`, `pending_review`

**Tidak ada** status `scheduled` di backend saat ini.

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/blogs` | [GET-blogs-list.md](./GET-blogs-list.md) |
| GET | `/api/v1/blogs/{id}` | [GET-blogs-by-id.md](./GET-blogs-by-id.md) |

Admin CRUD: lihat [admin-blog/](../admin-blog/README.md).

## Types & Hooks

- Types: `lib/types/blog.ts`
- Service: `blogService.list`, `blogService.getById`
- Hooks: `useBlogsQuery`, `useBlogQuery`
