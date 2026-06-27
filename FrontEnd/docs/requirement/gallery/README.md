# Gallery (Public)

Modul: galeri foto/video publik website DOSCOM.

## Functional Requirements

| ID | Requirement |
| --- | --- |
| GAL-01 | List gallery dengan pagination |
| GAL-02 | Filter by tahun (`start_year`, `end_year`) |
| GAL-03 | Tampilkan metadata: nama, tipe, deskripsi, tanggal event |

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/gallery` | [GET-gallery-list.md](./GET-gallery-list.md) |

Admin create/delete: lihat [admin-gallery/](../admin-gallery/README.md).

## Types & Hooks

- Types: `lib/types/gallery.ts`
- Service: `galleryService.list`
- Hook: `useGalleryQuery`
