# Frontend Progress

Dokumentasi status implementasi frontend DOSCOM per **13 Juni 2026**.

Branch aktif: `frontend-v2`

## Ringkasan

| Area | Status | Progress |
| --- | --- | --- |
| Infrastruktur data layer | Selesai | 100% |
| Dokumentasi requirements | Selesai | 100% |
| Public site — UI | Sebagian | ~70% |
| Public site — integrasi API | Belum | 0% |
| Admin panel — UI shell | Sebagian | ~35% |
| Admin panel — integrasi API | Belum | 0% |
| Auth & session flow | Belum | 0% |

**Kesimpulan:** Layer `services/`, `hooks/`, dan `lib/` sudah siap dipakai. Semua halaman (public & admin) masih **static/mock** — belum ada satu pun page yang memanggil API backend.

---

## Yang Sudah Selesai

### Stack & tooling

- [x] Astro 6 + React 19 + Tailwind CSS 4
- [x] `@tanstack/react-query` + `axios` terpasang
- [x] `@astrojs/react` integration
- [x] `npm run lint` — 0 errors
- [x] `npm run build` — berhasil (per sesi sebelumnya)
- [x] `.env.example` — `PUBLIC_API_URL`

### Data layer (`src/lib/`, `src/services/`, `src/hooks/`)

| Komponen | Status | Path |
| --- | --- | --- |
| Axios instance + interceptor | Selesai | `lib/axios.ts` |
| API paths terpusat | Selesai | `lib/api-path.ts` |
| Routes terpusat | Selesai | `lib/routes.ts` |
| Types per domain | Selesai | `lib/types/` |
| Helper functions | Selesai | `lib/func/` |
| Pesan UI & error | Selesai | `lib/message.ts` |
| Services (7 domain) | Selesai | `services/` |
| TanStack Query hooks | Selesai | `hooks/` |
| Query client singleton | Selesai | `hooks/query-client.ts` |
| Query keys | Selesai | `hooks/keys.ts` |
| QueryProvider + HOC | Selesai | `plugins/QueryProvider.tsx` |

**Services:** `auth`, `user`, `blog`, `gallery`, `work`, `pengurus`, `upload`

**Hooks tersedia (43):**

| Domain | Query | Mutation |
| --- | --- | --- |
| Auth | — | login, register, refresh, logout |
| User | me, list, detail, superAdmins | updateProfile, changePassword, CRUD user, createSuperAdmin, adminChangePassword |
| Blog | public list/detail, admin list/detail | create, update, delete |
| Gallery | list | create, delete |
| Work | byProjectType, admin list/detail | create, update, delete |
| Pengurus | byDivision, profile, admin list/detail/byUser | create/update/delete self, admin CRUD |
| Upload | files by category | delete file |

### Dokumentasi

- [x] `docs/requirement/` — spesifikasi API per endpoint (44 endpoint)
- [x] `docs/requirement/` — modul FR/UI per kategori + folder admin terpisah
- [x] `src/services/README.md`
- [x] `src/hooks/README.md`

### Public site — UI (static)

| Halaman | Route | Status UI |
| --- | --- | --- |
| Home | `/` | Selesai (static) |
| About | `/about` | Selesai (static) |
| Blog | `/blog` | Selesai (mock cards) |
| Works | `/works` | Selesai (static sections) |
| Story | `/story` | Selesai (static) |
| Gallery | `/gallery` | Placeholder saja |
| Contact | `/contact` | Selesai (form static) |
| Division list | `/division` | Selesai (static) |
| Division detail | `/division/[slug]` | Selesai (mock data) |
| Navbar | — | `HeroNavbar.tsx` + `StickyNavbar.astro`, pakai `ROUTES` |

### Admin panel — UI shell (static)

| Halaman | Route | Status UI |
| --- | --- | --- |
| Login | `/admin/login` | Form UI selesai, **tanpa handler submit** |
| Dashboard | `/admin` | Stat cards hardcode (61/21/56/238) |
| Core Team | `/admin/core-team` | Tabel + pagination mock |
| Layout | — | `AdminLayout`, `Sidebar`, `SearchModal`, `StatCard` |
| Sidebar menu | — | Dashboard & Core Team aktif; sisanya `href: '#'` |

---

## Yang Belum / Dalam Progress

### Integrasi API — belum dimulai di pages

Tidak ada import `hooks/` atau `services/` di folder `src/pages/` maupun komponen page.

| Task | Prioritas | Keterangan |
| --- | --- | --- |
| Wire `QueryProvider` ke layout admin | Tinggi | Belum dipasang di `AdminLayout` / `AuthLayout` |
| Login → API → redirect | Tinggi | `useLoginMutation` belum dipakai |
| Route guard `/admin/*` | Tinggi | Cek token, redirect ke login |
| Logout sidebar | Tinggi | `useLogoutMutation` + clear cache |
| Dashboard stats dinamis | Sedang | Agregasi client-side atau tunggu endpoint backend |
| Core Team CRUD | Sedang | Ganti mock → `useAdminPengurusListQuery` dll. |
| Sidebar user info dinamis | Sedang | `useMeQuery` — masih mock "Husnul Fikri" |
| Toast / notifikasi | Sedang | Pakai `lib/message.ts`, komponen toast TBD |

### Admin panel — halaman belum ada

| Modul | Route (rencana) | Status |
| --- | --- | --- |
| Super Admin | `/admin/super-admin` | Belum ada page |
| BPH / User management | `/admin/users` | Belum ada page |
| Blog admin | `/admin/blog` | Belum ada page |
| Gallery admin | `/admin/gallery` | Belum ada page |
| Works admin | `/admin/works` | Belum ada page |
| Profile | `/admin/profile` | Belum ada page |
| Search global | modal | UI ada, fungsi TBD |

Tambahkan route di `lib/routes.ts` saat page dibuat.

### Public site — integrasi API

| Halaman | Hook siap | Integrasi page |
| --- | --- | --- |
| Blog list/detail | `useBlogsQuery`, `useBlogQuery` | Belum |
| Gallery | `useGalleryQuery` | Belum (page masih placeholder) |
| Works | `useWorksByProjectTypeQuery` | Belum |
| Division pengurus | `usePengurusByDivisionQuery` | Belum (masih mock slug) |
| Contact form | — | Belum ada endpoint backend |

### Infrastruktur frontend — sisa

- [ ] Komponen React reusable untuk form admin (Input controlled, FileUpload)
- [ ] Rich text editor blog (Tiptap — TBD)
- [ ] Media library modal (`useUploadFilesQuery`)
- [ ] Error boundary / fallback UI global
- [ ] Refresh token otomatis saat 401 (fase 2)
- [ ] Role-based UI guard di komponen

---

## Checklist Integrasi Admin (dari requirements)

- [ ] Login → simpan token → redirect `/admin`
- [ ] Route guard: halaman admin redirect ke login jika tidak ada token
- [ ] Dashboard stats dari API (atau agregasi client-side)
- [ ] Core Team CRUD via pengurus admin API
- [ ] Blog, Gallery admin pages + hooks
- [ ] Profile: `useMeQuery` + `usePengurusProfileQuery`
- [ ] Logout: `useLogoutMutation` + redirect login
- [ ] Toast/notifikasi pakai `lib/message.ts`

---

## Ketergantungan Backend

Hal yang perlu koordinasi sebelum/saat integrasi:

| Item | Dampak |
| --- | --- |
| `GET /admin/pengurus` belum terdaftar di route | Core Team list & dashboard stat gagal |
| `POST /upload/image` disabled | Upload standalone tidak tersedia |
| Tidak ada `GET /admin/dashboard/stats` | Dashboard harus agregasi client-side |
| Login response typo `acces_token`, `message:` | Sudah di-handle di `extractAccessToken()` |
| MinIO down → upload routes tidak register | Upload/gallery/blog create bisa error |

---

## Rekomendasi Urutan Pengerjaan

```
Fase 1 — Auth & shell admin
  1. QueryProvider di AdminLayout + AuthLayout
  2. LoginForm (React) + useLoginMutation
  3. Route guard client-side
  4. Logout + sidebar useMeQuery

Fase 2 — Admin CRUD prioritas
  5. Core Team (admin-pengurus)
  6. Dashboard agregasi stats
  7. Profile page

Fase 3 — Admin content
  8. Blog admin
  9. Gallery admin
  10. Super Admin / User management

Fase 4 — Public site API
  11. Blog public
  12. Works + division pengurus
  13. Gallery public
```

---

## Referensi

| Dokumen | Path |
| --- | --- |
| Requirements & API spec | [../requirement/README.md](../requirement/README.md) |
| Services layer | [../../src/services/README.md](../../src/services/README.md) |
| Hooks layer | [../../src/hooks/README.md](../../src/hooks/README.md) |
| Shared requirements | [../requirement/shared.md](../requirement/shared.md) |
