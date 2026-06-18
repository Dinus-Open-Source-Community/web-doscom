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

Backend belum menyediakan `GET /admin/dashboard/stats`. Gunakan agregasi client-side (fase 1).

### Opsi A — Agregasi client-side (fase 1)

| Stat | Endpoint | Hook | Catatan |
| --- | --- | --- | --- |
| Core Team | `GET /admin/pengurus?divisi=` | `useAdminPengurusListQuery` | **Blocked** — route belum terdaftar di backend |
| Project | `GET /admin/works?page=1&limit=1` | `useAdminWorksQuery` | Ambil `totalPage` dari meta |
| Article | `GET /admin/blogs?page=1&limit=1` | `useAdminBlogsQuery` | Ambil `totalPage` dari meta |
| Images | `GET /gallery?page=1&limit=1` | `useGalleryQuery` | Ambil `totalPages` dari envelope `data` |

**Workaround Core Team count:** agregasi `GET /pengurus/division/{division}` untuk setiap divisi (`bph`, `pemro`, `jaringan`, `medcrev`, `data`) dan jumlahkan `data.length`. Data public tidak expose email, cukup untuk count.

### Opsi B — Endpoint dedicated (fase 2, backend)

```
GET /admin/dashboard/stats
```

Response contoh (envelope):

```json
{
  "success": true,
  "message": "Dashboard stats",
  "data": {
    "core_team_count": 61,
    "project_count": 21,
    "article_count": 56,
    "image_count": 238
  },
  "error": null
}
```

## Backend Gap — Admin Pengurus List

`GET /admin/pengurus` **tidak terdaftar** di router meski handler sudah ada. Dashboard stat Core Team tidak bisa rely on `useAdminPengurusListQuery` sampai backend fix. Lihat [admin-pengurus/GET-admin-pengurus-list.md](./admin-pengurus/GET-admin-pengurus-list.md).

## Implementasi Frontend

- Buat komponen React `AdminDashboard.tsx` dengan `withQueryProvider`
- Gunakan multiple `useQuery` paralel (TanStack handle otomatis)
- Handle partial failure per stat card (satu API gagal tidak block yang lain)

## Endpoint Pendukung

- [admin-works/GET-admin-works-list.md](./admin-works/GET-admin-works-list.md)
- [admin-blog/GET-admin-blogs-list.md](./admin-blog/GET-admin-blogs-list.md)
- [gallery/GET-gallery-list.md](./gallery/GET-gallery-list.md)
- [pengurus/GET-pengurus-by-division.md](./pengurus/GET-pengurus-by-division.md) (workaround Core Team)
