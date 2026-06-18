# Gallery (Public)

Modul public: list galeri foto kegiatan.

**Hak akses:** Public (tanpa autentikasi)

Endpoint admin ada di [admin-gallery/](../admin-gallery/README.md).

## Functional Requirements

| ID | Requirement |
| --- | --- |
| GAL-01 | List galeri dengan filter tahun |

## UI Requirements

| Elemen | Requirement |
| --- | --- |
| Filter tahun | Dropdown/range → query `start_year`, `end_year` |
| Grid/List | Thumbnail + metadata |

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/gallery` | [GET-gallery-list.md](./GET-gallery-list.md) |

## Types & Hooks

- Types: `lib/types/gallery.ts`
- Hooks: `hooks/gallery.ts` — `useGalleryQuery`
- Service: `galleryService`
