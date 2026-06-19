# Works (Public)

Modul public: list dan detail proyek DOSCOM.

**Hak akses:** Public (tanpa autentikasi)

Endpoint admin ada di [admin-works/](../admin-works/README.md).

## Functional Requirements

| ID | Requirement |
| --- | --- |
| WORK-05 | Filter by project type via query `projecttype` |
| WORK-08 | Detail work public (tanpa field internal) |
| WORK-09 | Daftar tipe proyek yang tersedia |

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/works` | [GET-works-list.md](./GET-works-list.md) |
| GET | `/api/v1/works/{id}` | [GET-works-by-id.md](./GET-works-by-id.md) |
| GET | `/api/v1/works/types` | [GET-works-types.md](./GET-works-types.md) |

## Types & Hooks

- Types: `lib/types/work.ts`
- Hooks: `hooks/work.ts` — `useWorksQuery`, `useWorkQuery`, `useWorkTypesQuery`
- Service: `workService.list`, `workService.getById`, `workService.getTypes`

## Catatan BE

Route `GET /works/types` terdaftar setelah `GET /works/:id` — bisa tertangkap sebagai `:id=types`. Perlu perbaikan urutan route di backend.
