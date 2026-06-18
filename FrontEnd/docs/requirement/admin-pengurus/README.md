# Admin — Pengurus / Core Team

Modul admin: halaman `/admin/core-team` — manajemen anggota core team DOSCOM.

**Hak akses:** middleware `KOOR`, `BPH`, `ADMIN`

| Group | Role JWT |
| --- | --- |
| ADMIN | `SuperAdmin` |
| KOOR | `KoorPemro`, `KoorJaringan`, `KoorData`, `KoorMedcrev`, `BPH` |
| BPH | `BPH` |

## Functional Requirements

| ID | Requirement |
| --- | --- |
| PENG-01 | Tambah member (create pengurus + link user) |
| PENG-02 | Edit data pengurus (nama, posisi, periode, sosmed, foto) |
| PENG-03 | Hapus pengurus |
| PENG-04 | Lookup pengurus by user ID |
| PENG-05 | Koordinator hanya kelola divisi sendiri (backend enforce) |
| PENG-06 | Upload foto profil (single file `file`) |

## UI Requirements (Core Team Page)

| UI Column | API Field |
| --- | --- |
| ID | `pengurus.id` |
| Member (nama + avatar) | `name`, `photo_url` |
| Email | `email` |
| Periode | `start_periode_year` – `end_periode_year` |
| Posisi | `position` |
| Actions | Edit, Delete |

## Divisi & Posisi Valid

Divisi: `bph`, `pemro`, `jaringan`, `medcrev`, `data`

Posisi: key camelCase dari `constants.ValidPosition` — contoh `ketuaUmum`, `koordinatorPemrograman`, `pemrogramanAnggota`, dll.

Periode: `start_periode_year`, `end_periode_year` (integer tahun).

## Backend Gap — List Endpoint

Handler `GetAllPengurusByDivision` **ada** di backend, tetapi route `GET /admin/pengurus` **belum terdaftar** di `pengurus_route.go`.

| Status | Detail |
| --- | --- |
| Handler | `PengurusHandler.GetAllPengurusByDivision` |
| Query param | `divisi` (opsional) |
| Route | **Tidak ada** — perlu ditambahkan tim backend |

Frontend `pengurusService.admin.list` dan `useAdminPengurusListQuery` sudah disiapkan, tetapi **akan gagal** sampai route diregister. Untuk dashboard stat Core Team, gunakan workaround client-side (agregasi dari divisi public) atau tunggu fix backend.

## Endpoints Tersedia

| Method | Path | Dokumen |
| --- | --- | --- |
| POST | `/api/v1/admin/pengurus` | [POST-admin-pengurus-create.md](./POST-admin-pengurus-create.md) |
| GET | `/api/v1/admin/pengurus/{id}` | [GET-admin-pengurus-by-id.md](./GET-admin-pengurus-by-id.md) |
| GET | `/api/v1/admin/pengurus/by-user/{user_id}` | [GET-admin-pengurus-by-user.md](./GET-admin-pengurus-by-user.md) |
| PUT | `/api/v1/admin/pengurus/{id}` | [PUT-admin-pengurus-by-id.md](./PUT-admin-pengurus-by-id.md) |
| DELETE | `/api/v1/admin/pengurus/delete/{id}` | [DELETE-admin-pengurus-by-id.md](./DELETE-admin-pengurus-by-id.md) |

## Endpoint Belum Tersedia (BE Gap)

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/admin/pengurus` | [GET-admin-pengurus-list.md](./GET-admin-pengurus-list.md) — **route not registered** |

## Types & Hooks

- Types: `lib/types/pengurus.ts`
- Hooks: `hooks/pengurus.ts` — `useAdminPengurusQuery`, `useCreateAdminPengurusMutation`, dll.
- Service: `pengurusService.admin.*`
