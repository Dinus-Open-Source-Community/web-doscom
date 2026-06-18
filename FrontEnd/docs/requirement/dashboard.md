# Dashboard

Modul: halaman `/admin` — ringkasan statistik organisasi.

## Functional Requirements

| ID | Requirement |
| --- | --- |
| DASH-01 | Tampilkan jumlah Core Team (pengurus) |
| DASH-02 | Tampilkan jumlah Project (works) |
| DASH-03 | Tampilkan jumlah Article (blog) |
| DASH-04 | Tampilkan jumlah Images (gallery) |
| DASH-05 | Data statistik real-time dari API (bukan hardcode) |
| DASH-06 | Loading skeleton saat fetch |
| DASH-07 | Tampilkan error jika gagal load stat |

## UI Saat Ini

Stat cards static di `pages/admin/index.astro` (Core Team: 61, Project: 21, Article: 56, Images: 238).

## API — Belum Ada Endpoint Agregat

### Opsi A — Agregasi client-side (fase 1)

| Stat | Endpoint | Hook |
| --- | --- | --- |
| Core Team | `GET /admin/pengurus?divisi=` | `useAdminPengurusListQuery` |
| Project | `GET /admin/works?page=1&limit=1` | `useAdminWorksQuery` |
| Article | `GET /admin/blogs?page=1&limit=1` | `useAdminBlogsQuery` |
| Images | `GET /gallery?page=1&limit=1` | `useGalleryQuery` |

Hitung total dari pagination meta atau `data.length`.

### Opsi B — Endpoint dedicated (fase 2, backend)

```
GET /admin/dashboard/stats
```

Response contoh:

```json
{
  "success": true,
  "message": "Dashboard stats",
  "data": {
    "core_team_count": 61,
    "project_count": 21,
    "article_count": 56,
    "image_count": 238
  }
}
```

## Implementasi Frontend

- Buat komponen React `AdminDashboard.tsx` dengan `withQueryProvider`
- Gunakan multiple `useQuery` paralel (TanStack handle otomatis)
- Search trigger (`SearchModal`) — requirement terpisah (global search TBD)

## Endpoint Pendukung

Lihat dokumentasi per endpoint:

- [admin-pengurus/GET-admin-pengurus-list.md](./admin-pengurus/GET-admin-pengurus-list.md)
- [admin-works/GET-admin-works-list.md](./admin-works/GET-admin-works-list.md)
- [admin-blog/GET-admin-blogs-list.md](./admin-blog/GET-admin-blogs-list.md)
- [gallery/GET-gallery-list.md](./gallery/GET-gallery-list.md)
