# Admin — Gallery Management

Modul admin: sidebar **Gallery** — upload & kelola foto kegiatan.

**Hak akses:** `SuperAdmin`, `KoorMedcrev`

## Functional Requirements

| ID | Requirement |
| --- | --- |
| GAL-02 | Upload batch foto kegiatan (max 5 file) |
| GAL-03 | Hapus item galeri |
| GAL-04 | Form metadata: nama, tipe, deskripsi, tanggal event |
| GAL-05 | Preview grid gambar |

## UI Requirements

| Elemen | Requirement |
| --- | --- |
| Upload zone | Drag & drop multi file, max 5 |
| Grid/List | Thumbnail + metadata |
| Delete | Konfirmasi modal (`UI_MESSAGES.common.confirmDelete`) |

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| POST | `/api/v1/admin/gallery` | [POST-admin-gallery-create.md](./POST-admin-gallery-create.md) |
| DELETE | `/api/v1/admin/gallery/{id}` | [DELETE-admin-gallery-by-id.md](./DELETE-admin-gallery-by-id.md) |

## Types & Hooks

- Types: `lib/types/gallery.ts`
- Hooks: `hooks/gallery.ts` — `useCreateGalleryMutation`, `useDeleteGalleryMutation`
- Service: `galleryService.admin.*`
