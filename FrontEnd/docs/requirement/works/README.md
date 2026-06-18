# Works (Public)

Modul public: list proyek DOSCOM per tipe.

**Hak akses:** Public (tanpa autentikasi)

Endpoint admin ada di [admin-works/](../admin-works/README.md).

## Functional Requirements

| ID | Requirement |
| --- | --- |
| WORK-05 | Filter by project type |

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/works/{projecttype}` | [GET-works-by-project-type.md](./GET-works-by-project-type.md) |

## Types & Hooks

- Types: `lib/types/work.ts`
- Hooks: `hooks/work.ts` — `useWorksByProjectTypeQuery`
- Service: `workService`
