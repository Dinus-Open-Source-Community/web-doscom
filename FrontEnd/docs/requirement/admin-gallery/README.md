# Admin — Gallery Management

Modul admin: upload dan hapus item gallery.

**Hak akses:** `SuperAdmin`, `KoorMedcrev`

## Functional Requirements

| ID | Requirement |
| --- | --- |
| GAL-ADM-01 | Upload multiple file (max 5) sekaligus |
| GAL-ADM-02 | Set metadata: nama, tipe, deskripsi, tanggal event |
| GAL-ADM-03 | Hapus gallery by ID |

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| POST | `/api/v1/admin/gallery` | [POST-admin-gallery-create.md](./POST-admin-gallery-create.md) |
| DELETE | `/api/v1/admin/gallery/{id}` | [DELETE-admin-gallery-by-id.md](./DELETE-admin-gallery-by-id.md) |

## Types & Hooks

- Types: `lib/types/gallery.ts`
- Service: `galleryService.admin.*`
- Hooks: `useCreateGalleryMutation`, `useDeleteGalleryMutation`
