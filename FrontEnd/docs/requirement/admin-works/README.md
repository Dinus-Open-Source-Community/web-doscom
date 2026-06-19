# Admin — Works / Project Management

Modul admin: manajemen proyek DOSCOM (halaman admin dedicated direncanakan).

**Hak akses CRUD:** middleware `ADMIN` + `KOOR` (`SuperAdmin`, semua koordinator termasuk `BPH`)

**Hak akses status moderation:** middleware `ADMIN` + `BPH` (`SuperAdmin`, `BPH`)

Koordinator hanya melihat works divisi sendiri (backend filter by role).

## Functional Requirements

| ID | Requirement |
| --- | --- |
| WORK-01 | List works per divisi (admin) |
| WORK-02 | Create work dengan teknologi, gambar, metadata |
| WORK-03 | Edit work |
| WORK-04 | Delete work |
| WORK-06 | Link pengurus sebagai owner (`pengurus_id`) |
| WORK-07 | Attach max 5 gallery images ke work |
| WORK-10 | Moderasi status work (BPH/SuperAdmin) |

## UI Requirements (Planned)

| Elemen | Requirement |
| --- | --- |
| Tabel | title, project_type, status, project_date, actions |
| Form technologies | Tag input multi value (max 3) |
| Image | Upload + picker dari gallery (`existingID_image[]`) |
| Pengurus select | Dropdown dari data pengurus (by-user atau manual input) |
| Status moderation | Tombol approve/reject untuk BPH via `PUT …/status` |

## Status Values

`draft`, `pending_review`, `published`, `unpublished`, `rejected`

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/admin/works` | [GET-admin-works-list.md](./GET-admin-works-list.md) |
| GET | `/api/v1/admin/works/{id}` | [GET-admin-works-by-id.md](./GET-admin-works-by-id.md) |
| POST | `/api/v1/admin/works` | [POST-admin-works-create.md](./POST-admin-works-create.md) |
| PUT | `/api/v1/admin/works/{id}` | [PUT-admin-works-by-id.md](./PUT-admin-works-by-id.md) |
| PUT | `/api/v1/admin/works/{id}/status` | [PUT-admin-works-status.md](./PUT-admin-works-status.md) |
| DELETE | `/api/v1/admin/works/{id}` | [DELETE-admin-works-by-id.md](./DELETE-admin-works-by-id.md) |

## Types & Hooks

- Types: `lib/types/work.ts`
- Form helper: `lib/func/work.ts` → `buildWorkFormData`
- Hooks: `hooks/work.ts` — `useAdminWorksQuery`, `useCreateWorkMutation`, `useUpdateWorkStatusMutation`, dll.
- Service: `workService.admin.*`
