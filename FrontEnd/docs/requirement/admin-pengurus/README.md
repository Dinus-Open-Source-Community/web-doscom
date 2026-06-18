# Admin — Pengurus / Core Team

Modul admin: halaman `/admin/core-team` — manajemen anggota core team DOSCOM.

**Hak akses:** `KOOR`, `BPH`, `ADMIN`

## Functional Requirements

| ID | Requirement |
| --- | --- |
| PENG-01 | List pengurus dengan filter divisi |
| PENG-02 | Tambah member (create pengurus + link user) |
| PENG-03 | Edit data pengurus (nama, posisi, periode, sosmed, foto) |
| PENG-04 | Hapus pengurus |
| PENG-05 | Search by nama/email |
| PENG-06 | Pagination datatable |
| PENG-07 | Upload foto profil (single file) |
| PENG-08 | Koordinator hanya kelola divisi sendiri (backend enforce) |

## UI Requirements (Core Team Page)

| UI Column | API Field |
| --- | --- |
| ID | `CT-{id}` atau `pengurus.id` |
| Member (nama + avatar) | `name`, `photo_url` |
| Email | `email` |
| Join Date | `created_at` (format display) |
| Status | TBD — backend belum punya field status |
| Actions | Edit, Delete |

Fitur tombol: **Add Member** (modal create), **Download** (export CSV, frontend only), **Search**, **Filter divisi**.

## Divisi & Posisi Valid

Divisi: `pemro`, `jaringan`, `data`, `medcrev`, `bph`

Posisi contoh: `KoorPemro`, `PemroAng`, `ketum`, `KoorJaringan`, dll. (lihat `BackEnd/internal/constants/constants.go`).

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/admin/pengurus` | [GET-admin-pengurus-list.md](./GET-admin-pengurus-list.md) |
| POST | `/api/v1/admin/pengurus` | [POST-admin-pengurus-create.md](./POST-admin-pengurus-create.md) |
| GET | `/api/v1/admin/pengurus/{id}` | [GET-admin-pengurus-by-id.md](./GET-admin-pengurus-by-id.md) |
| GET | `/api/v1/admin/pengurus/by-user/{user_id}` | [GET-admin-pengurus-by-user.md](./GET-admin-pengurus-by-user.md) |
| PUT | `/api/v1/admin/pengurus/{id}` | [PUT-admin-pengurus-by-id.md](./PUT-admin-pengurus-by-id.md) |
| DELETE | `/api/v1/admin/pengurus/delete/{id}` | [DELETE-admin-pengurus-by-id.md](./DELETE-admin-pengurus-by-id.md) |

## Types & Hooks

- Types: `lib/types/pengurus.ts`
- Hooks: `hooks/pengurus.ts` — `useAdminPengurusListQuery`, `useCreateAdminPengurusMutation`, dll.
- Service: `pengurusService.admin.*`
