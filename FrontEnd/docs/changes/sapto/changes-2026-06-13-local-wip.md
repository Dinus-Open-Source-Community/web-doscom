# changes-2026-06-13-local-wip

| Field | Value |
| --- | --- |
| **Status** | Belum di-commit (working tree) |
| **Author** | zappto / Sapto Gusty |
| **Date** | 2026-06-13 |
| **Branch** | `frontend-v2` |

> Setelah di-commit, rename file ini menjadi `changes-2026-06-13-[hash].md` dan update [README.md](./README.md).

## Ringkasan

Setup infrastruktur frontend v2: React Query + Axios data layer, centralisasi paths/types/messages, dokumentasi requirements & progress. Perbaikan kecil UI navbar dan config Astro/Vite.

**Belum ada integrasi API di halaman** — layer siap pakai, pages masih static/mock.

---

## 1. Dependencies & Config

### Modified

| File | Perubahan |
| --- | --- |
| `package.json` | Tambah `axios`, `@astrojs/check`, `typescript`; script `lint`, `dev` dengan `NODE_ENV=development` |
| `package-lock.json` | Lock dependencies |
| `astro.config.mjs` | Vite `optimizeDeps` fix untuk React JSX dev |
| `tsconfig.json` | Penyesuaian TypeScript |
| `prettier.config.mjs` | **Baru** — config Prettier |
| `.env.example` | **Baru** — `PUBLIC_API_URL=http://localhost:8080/api/v1` |

---

## 2. Data Layer (Baru)

### `src/lib/`

| File/Folder | Fungsi |
| --- | --- |
| `axios.ts` | Axios instance, Bearer interceptor, `toApiError` |
| `api-path.ts` | Semua path API backend terpusat |
| `routes.ts` | Route Astro pages terpusat |
| `message.ts` | Pesan UI, error, success terpusat |
| `types/` | Types per domain: auth, user, blog, gallery, work, pengurus, upload, common |
| `func/` | Helpers: http, auth, error, message, blog, work |

### `src/services/` (7 service)

| Service | Domain |
| --- | --- |
| `auth.service.ts` | login, register, refresh, logout |
| `user.service.ts` | user + admin super admin |
| `blog.service.ts` | public + admin blog |
| `gallery.service.ts` | public + admin gallery |
| `work.service.ts` | public + admin works |
| `pengurus.service.ts` | self + admin pengurus |
| `upload.service.ts` | list/delete files MinIO |

### `src/hooks/` (43 hooks)

TanStack Query hooks untuk semua service di atas — query + mutation dengan cache invalidation.

| File | Hooks |
| --- | --- |
| `auth.ts` | login, register, refresh, logout |
| `user.ts` | me, users, CRUD, super admin |
| `blog.ts` | public + admin blog CRUD |
| `gallery.ts` | list, create, delete |
| `work.ts` | public + admin works CRUD |
| `pengurus.ts` | division, profile, self, admin CRUD |
| `upload.ts` | files list, delete |
| `keys.ts` | Query key factory |
| `query-client.ts` | Singleton QueryClient |

### `src/plugins/QueryProvider.tsx`

- Modified: merge `withQueryProvider` HOC
- **Belum** dipasang di `AdminLayout` / `AuthLayout`

---

## 3. UI Fixes (Modified)

| File | Perubahan |
| --- | --- |
| `components/layouts/HeroNavbar.tsx` | Fix `jsxDEV is not a function`; pakai `ROUTES` |
| `components/layouts/StickyNavbar.astro` | Pakai `ROUTES` |
| `components/layouts/Sidebar.astro` | Pakai `ROUTES.admin.*` |
| `components/layouts/Footer.astro` | Refactor links/styling |
| `components/ui/Card.astro` | Penyesuaian variant |
| `components/ui/admin/SearchModal.astro` | Minor fix |

---

## 4. Dokumentasi (Baru)

### `docs/requirement/` — 59 files

Spesifikasi API per endpoint + modul FR/UI:

```
requirement/
├── README.md, shared.md, dashboard.md
├── auth/, user/, blog/, gallery/, works/, pengurus/, upload/
└── admin-user/, admin-blog/, admin-gallery/, admin-works/, admin-pengurus/
```

### `docs/progress/` — 2 files

- `README.md` — ringkasan progress implementasi
- `checklist.md` — checklist granular per modul

### README layer

- `src/services/README.md`
- `src/hooks/README.md`

---

## Statistik

| Kategori | Jumlah |
| --- | --- |
| Modified | 11 files |
| Untracked (baru) | 103 files |
| Total baris diff (modified) | +1622 / -639 |

---

## Yang Belum (follow-up commit)

- [ ] Commit semua perubahan frontend
- [ ] Wire `QueryProvider` ke layout admin
- [ ] Login form → `useLoginMutation`
- [ ] Route guard `/admin/*`
- [ ] Integrasi API di pages

---

## Command Referensi

```bash
# Lihat semua perubahan
git status FrontEnd/
git diff --stat FrontEnd/

# Setelah commit
git log -1 --format='%h %s'
# → rename file: changes-2026-06-13-[hash].md
```

## Related

- Commit hari yang sama (backend): [changes-2026-06-13-1be0749.md](./changes-2026-06-13-1be0749.md)
- Progress: [../../docs/progress/README.md](../../docs/progress/README.md)
