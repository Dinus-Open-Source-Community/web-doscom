# RBAC dan Authorization Refactor

Dokumen ini merangkum authorization yang saat ini ada di backend, lalu memetakan aturan tersebut menjadi rancangan RBAC yang lebih fleksibel dan database-driven.

## Kesimpulan

Project ini sudah menerapkan RBAC, tetapi masih dalam bentuk hard-coded RBAC ditambah resource policy.

RBAC-nya terlihat dari:

- User memiliki `role`.
- JWT membawa role user.
- Middleware mengecek role sebelum endpoint diakses.
- Service dan authorization package mengecek role sebelum menjalankan aksi.

Namun sistem ini belum fleksibel karena role, role group, division, role level, field permission, dan status permission masih ditulis langsung di kode.

Target refactor yang disarankan:

- Role dan permission disimpan di database.
- Endpoint access memakai permission, bukan role string langsung.
- Aturan yang bergantung pada data resource tetap berada di policy layer.

## Role Saat Ini

Role didefinisikan di `internal/constants/constants.go`.

| Role Key | Group | Divisi | Level |
|---|---|---|---|
| `SuperAdmin` | `admin` | `bph` | 1 |
| `BPH` | `koor` | `bph` | 2 |
| `KoorPemro` | `koor` | `pemro` | 2 |
| `KoorJaringan` | `koor` | `jaringan` | 2 |
| `KoorData` | `koor` | `data` | 2 |
| `KoorMedcrev` | `koor` | `medcrev` | 2 |
| `pemroAnggota` | `pengurus` | `pemro` | 3 |
| `jaringanAnggota` | `pengurus` | `jaringan` | 3 |
| `medcrevAnggota` | `pengurus` | `medcrev` | 3 |
| `dataAnggota` | `pengurus` | `data` | 3 |
| `BPHAnggota` | `pengurus` | `bph` | 3 |

Rekomendasi kolom tabel `roles`:

```txt
id
name
group_name
division
level
is_active
created_at
updated_at
```

## Role Alias Saat Ini

Alias ini sekarang ada di `internal/auth/middleware.go`.

| Alias | Role yang termasuk |
|---|---|
| `ADMIN` | `SuperAdmin` |
| `KOOR` | `KoorPemro`, `KoorJaringan`, `KoorData`, `KoorMedcrev`, `BPH` |
| `BPH` | `BPH` |
| `PENGURUS` | `pemroAnggota`, `jaringanAnggota`, `medcrevAnggota`, `dataAnggota`, `BPHAnggota` |

Saat sudah database-driven, alias seperti `ADMIN`, `KOOR`, dan `PENGURUS` sebaiknya tidak lagi dipakai sebagai role. Nilai tersebut lebih cocok menjadi `group_name` di tabel `roles`.

## Matrix Rules Saat Ini

Matrix ini adalah gabungan dari route middleware, handler, service, dan package authorization.

| Resource | Action | Role yang boleh | Scope atau policy tambahan |
|---|---|---|---|
| `blog` | public read list/detail | public | Endpoint public |
| `blog` | admin list | `SuperAdmin`, `KoorMedcrev` | Tidak ada filter role tambahan |
| `blog` | create | `SuperAdmin`, `KoorMedcrev` | Wajib ada gallery image |
| `blog` | update | `SuperAdmin`, `KoorMedcrev` | Status dibatasi oleh `CanSetStatusBlog` |
| `blog` | delete | `SuperAdmin`, `KoorMedcrev` | Tidak ada ownership check |
| `gallery` | public read | public | Bisa filter by year |
| `gallery` | create | `SuperAdmin`, `KoorMedcrev` | Maksimal 5 files |
| `gallery` | delete | `SuperAdmin`, `KoorMedcrev` | Delete DB dan file storage |
| `work` | public read list/detail/types | public | Endpoint public |
| `work` | create | `SuperAdmin`, semua `koor` termasuk `BPH` | Division work otomatis dari role actor |
| `work` | admin list | `SuperAdmin`, semua `koor` termasuk `BPH` | `SuperAdmin` dan `BPH` lihat semua divisi, koor lain by own division |
| `work` | admin detail | `SuperAdmin`, semua `koor` | Saat ini tidak filter divisi by detail |
| `work` | update content | `SuperAdmin`, semua `koor` | Admin bebas, koor hanya kalau status `draft` atau `rejected` |
| `work` | update status/moderate | `SuperAdmin`, `BPH` | Transition status dibatasi |
| `work` | delete | `SuperAdmin`, semua `koor` | Status delete dibatasi |
| `pengurus` | public read by division | public | Response public tanpa email dan user id |
| `pengurus` | create own profile | semua authenticated role | `user_id` dipaksa current user |
| `pengurus` | update own profile | semua authenticated role | Field dibatasi role, self update boleh foto |
| `pengurus` | delete own profile | route mengizinkan semua role | Service menolak role `pengurus`, jadi behavior sekarang konflik |
| `pengurus` | managed create | `SuperAdmin`, semua `koor` termasuk `BPH` | Divisi/position mengikuti role, kecuali admin bebas |
| `pengurus` | managed read by id/user_id | `SuperAdmin`, semua `koor` | Koor hanya divisi sendiri, admin semua |
| `pengurus` | managed list by division | `SuperAdmin`, semua `koor` | Admin pakai query divisi, koor dipaksa divisi sendiri |
| `pengurus` | managed update | `SuperAdmin`, semua `koor` | Admin bebas, koor tidak boleh update level sama/lebih tinggi dan harus sesuai divisi |
| `pengurus` | managed delete | `SuperAdmin`, semua `koor` | Admin bebas, koor hanya divisi sendiri, pengurus ditolak |
| `user` | read own profile | semua authenticated role | Data current user |
| `user` | update own profile | semua authenticated role | Field: `username`, `email`, `full_name` |
| `user` | change own password | semua authenticated role | Wajib old password benar |
| `user` | create user | `SuperAdmin`, semua `koor` termasuk `BPH` | Admin bisa pilih role, koor auto assign anggota divisinya |
| `user` | list users | `SuperAdmin`, semua `koor` termasuk `BPH` | Admin lihat semua non-superadmin, koor lihat anggota divisinya |
| `user` | read user by id | `SuperAdmin`, semua `koor` termasuk `BPH` | Admin semua, koor hanya same division |
| `user` | update user by id | `SuperAdmin`, semua `koor` termasuk `BPH` | Saat ini tidak ada division/level check di service |
| `user` | delete user | `SuperAdmin`, semua `koor` termasuk `BPH` | Actor harus level lebih tinggi, koor hanya same division |
| `user` | create super admin | `SuperAdmin` | Endpoint khusus |
| `user` | list super admins | `SuperAdmin` | Exclude current user |
| `user` | admin change password | `SuperAdmin` | Hanya target level lebih rendah |
| `upload` | list/delete files | semua authenticated role | Tidak ada role check tambahan |

## Blog Status Rules

Aturan ini berasal dari `internal/authorization/blog/blog.go`.

| Role | Boleh set status |
|---|---|
| `KoorMedcrev` | `draft`, `pending_review`, `unpublished` |
| `BPH` | `published`, `unpublished`, `rejected` |
| `SuperAdmin` | `draft`, `pending_review`, `published`, `unpublished`, `rejected` |

Catatan:

- Route blog admin saat ini hanya mengizinkan `SuperAdmin` dan `KoorMedcrev`.
- Rule `BPH` di blog policy belum efektif karena route tidak memberi akses blog admin ke `BPH`.
- Ada bug di `CreateBlog`: kondisi status memakai `||`, sehingga logic status `KoorMedcrev` menjadi salah. Seharusnya memakai `&&`.

Kode bermasalah:

```go
if userRole == constants.RoleKeyKoorMedcrev &&
	(blogDetail.Status != constants.StatusDraft || blogDetail.Status != constants.StatusPending) {
	return nil, fmt.Errorf("you are can only set status to draft or pending")
}
```

Seharusnya:

```go
if userRole == constants.RoleKeyKoorMedcrev &&
	blogDetail.Status != constants.StatusDraft &&
	blogDetail.Status != constants.StatusPending {
	return nil, fmt.Errorf("you are can only set status to draft or pending")
}
```

## Work Status Rules

Aturan ini berasal dari `internal/authorization/work/work.go`.

### Set Status

| Role | Boleh set status |
|---|---|
| `KoorPemro` | `draft`, `pending_review` |
| `KoorJaringan` | `draft`, `pending_review` |
| `KoorData` | `draft`, `pending_review` |
| `KoorMedcrev` | `draft`, `pending_review` |
| `BPH` | `published`, `unpublished`, `rejected` |
| `SuperAdmin` | `draft`, `pending_review`, `published`, `unpublished`, `rejected` |

### View Status

| Role | Bisa lihat status |
|---|---|
| `SuperAdmin` | semua status |
| `BPH` | `pending_review`, `published`, `unpublished` |
| Koor divisi | semua status, tetapi difilter divisi |

### Delete Status

| Role | Bisa delete work dengan status |
|---|---|
| Koor divisi | `draft`, `rejected` |
| `BPH` | `published`, `unpublished` |
| `SuperAdmin` | `published`, `unpublished`, `rejected` |

### Moderation Transition

| Current status | Target status |
|---|---|
| `pending_review` | `published`, `rejected` |
| `published` | `unpublished` |
| `unpublished` | `published` |

## Pengurus Field Permission

Aturan ini berasal dari `RoleFieldPermission`.

| Group role | Editable fields |
|---|---|
| `admin` | `name`, `email`, `divisi`, `position`, `start_periode_year`, `end_periode_year`, `photo_url` |
| `koor` | `name`, `email`, `divisi`, `position`, `start_periode_year`, `end_periode_year`, `photo_url` |
| `pengurus` | `name`, `email`, `divisi`, `start_periode_year`, `end_periode_year`, `photo_url` |

Policy tambahan:

- Admin boleh update foto pengurus.
- Pengurus boleh update foto kalau update data sendiri.
- Koor tidak boleh update foto pengurus lain.
- Pengurus tidak boleh update data orang lain.
- Koor tidak boleh update target dengan level sama atau lebih tinggi.

## Position dan Division Rules

Aturan ini berasal dari `ValidPosition` dan `PositionGroup`.

| Division | Positions |
|---|---|
| `bph` | `ketuaUmum`, `kepalaBidangSumberDayaUmum`, `projectManagerI`, `projectManagerII`, `sekretarisUmumI`, `sekretarisUmumII`, `kepalaBidangHubunganMasyarakat`, `koordinatorHubunganMasyarakatExternal`, `hubunganMasyarakatExternal`, `bendaharaUmumI`, `bendaharaUmumII` |
| `pemro` | `koordinatorPemrograman`, `pemrogramanAnggota` |
| `jaringan` | `koordinatorJaringan`, `jaringanAnggota` |
| `medcrev` | `koordinatorMediaCreative`, `mediaCreativeAnggota` |
| `data` | `koordinatorData`, `dataAnggota` |

Admin boleh memilih division dan position dari request.

Non-admin:

- Division dipaksa mengikuti division role actor.
- Position harus valid di division actor.

## Auto Assign Role

Aturan ini berasal dari `AutoAsignRole`.

| Creator role | Assigned role |
|---|---|
| `KoorPemro` | `pemroAnggota` |
| `KoorJaringan` | `jaringanAnggota` |
| `KoorData` | `dataAnggota` |
| `KoorMedcrev` | `medcrevAnggota` |
| `BPH` | `BPHAnggota` |

Admin tidak memakai auto assign. Admin bisa memilih role target dari request.

## Permission Naming

Gunakan format:

```txt
resource:action
resource:action:scope
```

Pakai `resource:action` untuk permission umum.

Pakai `resource:action:scope` kalau scope benar-benar mengubah keputusan authorization, misalnya `self`, `division`, atau `any`.

Contoh permission yang disarankan:

```txt
blog:read:public
blog:read:admin
blog:create
blog:update
blog:delete
blog:set_status

gallery:read:public
gallery:create
gallery:delete

work:read:public
work:read:admin
work:create
work:update
work:delete
work:moderate_status

pengurus:read:public
pengurus:read:self
pengurus:read:division
pengurus:read:any
pengurus:create:self
pengurus:create:division
pengurus:create:any
pengurus:update:self
pengurus:update:division
pengurus:update:any
pengurus:delete:division
pengurus:delete:any

user:read:self
user:update:self
user:change_password:self
user:create:division
user:create:any
user:read:division
user:read:any
user:update:division
user:update:any
user:delete:division
user:delete:any
user:create_super_admin
user:read_super_admin
user:change_password:any

upload:read
upload:delete
```

## Mapping Permission Awal

### SuperAdmin

`SuperAdmin` sebaiknya memiliki semua permission internal/admin.

Minimal:

```txt
blog:read:admin
blog:create
blog:update
blog:delete
blog:set_status
gallery:create
gallery:delete
work:read:admin
work:create
work:update
work:delete
work:moderate_status
pengurus:read:any
pengurus:create:any
pengurus:update:any
pengurus:delete:any
user:read:any
user:create:any
user:update:any
user:delete:any
user:create_super_admin
user:read_super_admin
user:change_password:any
upload:read
upload:delete
```

### BPH

```txt
work:read:admin
work:create
work:moderate_status
work:delete
pengurus:read:division
pengurus:create:division
pengurus:update:division
pengurus:delete:division
user:read:division
user:create:division
user:update:division
user:delete:division
user:read:self
user:update:self
user:change_password:self
pengurus:read:self
pengurus:create:self
pengurus:update:self
upload:read
upload:delete
```

### KoorMedcrev

```txt
blog:read:admin
blog:create
blog:update
blog:delete
blog:set_status
gallery:create
gallery:delete
work:read:admin
work:create
work:update
work:delete
pengurus:read:division
pengurus:create:division
pengurus:update:division
pengurus:delete:division
user:read:division
user:create:division
user:update:division
user:delete:division
user:read:self
user:update:self
user:change_password:self
pengurus:read:self
pengurus:create:self
pengurus:update:self
upload:read
upload:delete
```

### KoorPemro, KoorJaringan, KoorData

```txt
work:read:admin
work:create
work:update
work:delete
pengurus:read:division
pengurus:create:division
pengurus:update:division
pengurus:delete:division
user:read:division
user:create:division
user:update:division
user:delete:division
user:read:self
user:update:self
user:change_password:self
pengurus:read:self
pengurus:create:self
pengurus:update:self
upload:read
upload:delete
```

### Anggota/Pengurus

Berlaku untuk:

- `pemroAnggota`
- `jaringanAnggota`
- `medcrevAnggota`
- `dataAnggota`
- `BPHAnggota`

```txt
user:read:self
user:update:self
user:change_password:self
pengurus:read:self
pengurus:create:self
pengurus:update:self
upload:read
upload:delete
```

## Database Schema Yang Disarankan

Minimal:

```sql
CREATE TABLE roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    group_name VARCHAR(50) NOT NULL,
    division VARCHAR(50),
    level INT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE permissions (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE role_permissions (
    role_id BIGINT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permission_id BIGINT NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);
```

Lalu migrasi user:

```sql
ALTER TABLE users ADD COLUMN role_id BIGINT REFERENCES roles(id);
```

Setelah semua data termigrasi dan kode sudah memakai `role_id`, kolom `users.role` bisa dihapus.

## Resource Policy Yang Tetap Dibutuhkan

Jangan semua aturan dipaksa masuk ke tabel `permissions`.

Aturan berikut tetap lebih sehat sebagai policy function:

- Same division check.
- Self ownership check.
- Role level comparison.
- Position harus sesuai division.
- Status transition work.
- Editable fields pengurus.
- Default status per role.
- Auto assign role saat koor create user.

Contoh pemakaian setelah refactor:

```go
RequirePermission("work:update")
policy.CanUpdateWork(actor, work)

RequirePermission("pengurus:update:division")
policy.CanUpdatePengurus(actor, targetPengurus)

RequirePermission("work:moderate_status")
policy.CanModerateWorkStatus(actor, currentStatus, targetStatus)
```

## Struktur Authorization Modular Monolith

Struktur yang cocok dengan arah project baru:

```txt
authorization/
  entity/
    role.go
    permission.go
  middleware.go
  module.go
  resource_policy.go
  service.go
```

Tanggung jawab setiap file:

| File | Tanggung jawab |
|---|---|
| `entity/role.go` | Model `Role` |
| `entity/permission.go` | Model `Permission` dan `RolePermission` |
| `service.go` | Load role, load permission, `HasPermission`, `GetActor` |
| `middleware.go` | `RequireAuth`, `RequirePermission`, `RequireAnyPermission` |
| `resource_policy.go` | Same division, ownership, level, status transition, field policy |
| `module.go` | Wiring dependency authorization module |

## Urutan Refactor

1. Tambah tabel `roles`, `permissions`, dan `role_permissions`.
2. Seed semua role dari constants lama.
3. Seed permission dari daftar permission awal.
4. Seed mapping role-permission.
5. Tambah `role_id` ke tabel `users`.
6. Backfill `users.role_id` dari `users.role`.
7. Ubah entity `User` dari `Role string` menjadi `RoleID`.
8. Ubah login/JWT agar membawa `role_id` dan `role_name`.
9. Buat authorization service database-driven.
10. Ubah middleware dari role-based menjadi permission-based.
11. Pindahkan `GetRoleInfo` agar membaca dari DB/cache, bukan constants.
12. Pindahkan hard-coded resource policy ke `resource_policy.go`.
13. Setelah stabil, hapus `users.role` dan constants role lama.

## Catatan Risiko

Ada beberapa inkonsistensi yang sebaiknya dibersihkan saat refactor:

- Route blog tidak mengizinkan `BPH`, tetapi blog policy punya rule untuk `BPH`.
- `CreateBlog` punya bug logic status karena memakai `||` bukan `&&`.
- `UpdateUserByID` saat ini route mengizinkan admin/koor, tetapi service tidak mengecek division atau level target.
- `DeleteMyPengurusProfile` route mengizinkan pengurus, tetapi service menolak pengurus delete data sendiri.
- Beberapa authorization check dobel di route, handler, dan service. Setelah refactor, middleware cukup mengecek permission umum, lalu service/policy mengecek resource-specific rule.

## Target Akhir

Target desain yang disarankan:

```txt
RBAC DB:
- siapa punya permission apa

Resource policy:
- permission itu boleh dipakai ke resource ini atau tidak
```

Dengan begitu, sistem menjadi fleksibel tanpa kehilangan aturan bisnis yang sudah ada.
