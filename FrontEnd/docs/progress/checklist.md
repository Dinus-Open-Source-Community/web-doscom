# Progress Checklist — Detail

Checklist granular per modul. Update centang saat selesai diimplementasi di **page/komponen** (bukan hanya di layer hooks/services).

Legenda: ✅ selesai · 🟡 partial · ⬜ belum

---

## Infrastruktur

| Item                    | Layer | Page/UI | Catatan                    |
| ----------------------- | ----- | ------- | -------------------------- |
| Axios + interceptor     | ✅    | —       | `lib/axios.ts`             |
| API paths               | ✅    | —       | `lib/api-path.ts`          |
| Types                   | ✅    | —       | `lib/types/`               |
| Services (7 domain)     | ✅    | —       | `src/services/`            |
| Hooks (43)              | ✅    | —       | `src/hooks/`               |
| QueryProvider           | ✅    | ⬜      | Ada tapi belum di layout   |
| Message / error helpers | ✅    | ⬜      | Belum dipakai di UI        |
| Routes constant         | ✅    | 🟡      | Admin routes belum lengkap |

---

## Auth & Session

| ID      | Requirement             | Service/Hook            | Page |
| ------- | ----------------------- | ----------------------- | ---- |
| AUTH-01 | Login email & password  | ✅ `useLoginMutation`   | ⬜   |
| AUTH-02 | Simpan token + redirect | ✅ `authService.login`  | ⬜   |
| AUTH-03 | Bearer otomatis         | ✅ interceptor          | —    |
| AUTH-04 | Logout + clear cache    | ✅ `useLogoutMutation`  | ⬜   |
| AUTH-05 | Refresh token (fase 2)  | ✅ `useRefreshMutation` | ⬜   |
| AUTH-06 | Route guard admin       | —                       | ⬜   |

---

## Dashboard

| ID      | Requirement      | Service/Hook  | Page    |
| ------- | ---------------- | ------------- | ------- |
| DASH-01 | Stat Core Team   | ✅ hook ready | ⬜ mock |
| DASH-02 | Stat Project     | ✅ hook ready | ⬜ mock |
| DASH-03 | Stat Article     | ✅ hook ready | ⬜ mock |
| DASH-04 | Stat Images      | ✅ hook ready | ⬜ mock |
| DASH-05 | Data real-time   | —             | ⬜      |
| DASH-06 | Loading skeleton | —             | ⬜      |
| DASH-07 | Error state      | —             | ⬜      |

---

## User & Super Admin

| ID      | Requirement              | Service/Hook    | Page       |
| ------- | ------------------------ | --------------- | ---------- |
| USER-01 | List super admin         | ✅              | ⬜ no page |
| USER-02 | Create super admin       | ✅              | ⬜         |
| USER-03 | Admin change password    | ✅              | ⬜         |
| USER-04 | List user                | ✅              | ⬜         |
| USER-05 | CRUD user                | ✅              | ✅         |
| USER-06 | Search/filter/pagination | —               | ⬜         |
| USER-07 | Form create user         | —               | ⬜         |
| PROF-01 | Tampil profil login      | ✅ `useMeQuery` | ⬜         |
| PROF-02 | Edit profil              | ✅              | ⬜         |
| PROF-03 | Ganti password           | ✅              | ⬜         |
| PROF-06 | Sidebar nama/email real  | ✅              | ⬜ mock    |

---

## Pengurus / Core Team

| ID      | Requirement            | Service/Hook | Page          |
| ------- | ---------------------- | ------------ | ------------- |
| PENG-01 | List + filter divisi   | ✅           | ⬜ mock table |
| PENG-02 | Tambah member          | ✅           | ⬜            |
| PENG-03 | Edit pengurus          | ✅           | ⬜            |
| PENG-04 | Hapus pengurus         | ✅           | ⬜            |
| PENG-05 | Search                 | —            | ⬜ UI only    |
| PENG-06 | Pagination             | —            | ⬜ mock       |
| PENG-07 | Upload foto            | ✅ service   | ⬜            |
| PENG-08 | Divisi scope (backend) | —            | —             |
| PROF-04 | Profil pengurus self   | ✅           | ⬜            |
| PROF-05 | Create profil pengurus | ✅           | ⬜            |

---

## Blog

| ID      | Requirement            | Service/Hook | Page       |
| ------- | ---------------------- | ------------ | ---------- |
| BLOG-01 | List + filter (public) | ✅           | ⬜ mock    |
| BLOG-02 | Create blog (admin)    | ✅           | ⬜ no page |
| BLOG-03 | Edit blog              | ✅           | ⬜         |
| BLOG-04 | Delete blog            | ✅           | ⬜         |
| BLOG-05 | Thumbnail di tabel     | —            | ⬜         |
| BLOG-06 | Status draft/published | —            | ⬜         |
| BLOG-07 | Max 3 kategori         | —            | ⬜         |

---

## Gallery

| ID     | Requirement          | Service/Hook | Page                |
| ------ | -------------------- | ------------ | ------------------- |
| GAL-01 | List + filter tahun  | ✅           | ⬜ placeholder page |
| GAL-02 | Upload batch (max 5) | ✅           | ⬜                  |
| GAL-03 | Hapus item           | ✅           | ⬜                  |
| GAL-04 | Form metadata        | —            | ⬜                  |
| GAL-05 | Preview grid         | —            | ⬜                  |

---

## Works

| ID      | Requirement            | Service/Hook | Page       |
| ------- | ---------------------- | ------------ | ---------- |
| WORK-01 | List admin             | ✅           | ⬜ no page |
| WORK-02 | Create work            | ✅           | ⬜         |
| WORK-03 | Edit work              | ✅           | ⬜         |
| WORK-04 | Delete work            | ✅           | ⬜         |
| WORK-05 | Public by project type | ✅           | ⬜ mock    |
| WORK-06 | Link pengurus owner    | —            | ⬜         |
| WORK-07 | Max 5 gallery images   | ✅ payload   | ⬜         |

---

## Upload & Media

| ID     | Requirement           | Service/Hook        | Page |
| ------ | --------------------- | ------------------- | ---- |
| UPL-01 | List file by category | ✅                  | ⬜   |
| UPL-02 | Delete file           | ✅                  | ⬜   |
| UPL-03 | Picker di blog/work   | —                   | ⬜   |
| UPL-04 | POST upload/image     | ⬜ backend disabled | ⬜   |

---

## Public Site Pages

| Halaman            | UI  | API | Catatan           |
| ------------------ | --- | --- | ----------------- |
| `/`                | ✅  | ⬜  | Static            |
| `/about`           | ✅  | ⬜  | Static            |
| `/blog`            | 🟡  | ⬜  | Mock cards        |
| `/works`           | 🟡  | ⬜  | Static sections   |
| `/gallery`         | ⬜  | ⬜  | Placeholder text  |
| `/story`           | ✅  | ⬜  | Static            |
| `/contact`         | 🟡  | ⬜  | Form tanpa submit |
| `/division`        | ✅  | ⬜  | Static            |
| `/division/[slug]` | 🟡  | ⬜  | Mock divisions    |

---

## Admin Pages

| Halaman            | UI  | API | Catatan           |
| ------------------ | --- | --- | ----------------- |
| `/admin/login`     | ✅  | ⬜  | No submit handler |
| `/admin`           | 🟡  | ⬜  | Hardcoded stats   |
| `/admin/core-team` | 🟡  | ⬜  | Mock 4 rows       |
| Super Admin        | ⬜  | ⬜  | —                 |
| BPH / Users        | ⬜  | ⬜  | —                 |
| Blog admin         | ⬜  | ⬜  | —                 |
| Gallery admin      | ⬜  | ⬜  | —                 |
| Works admin        | ⬜  | ⬜  | —                 |
| Profile            | ⬜  | ⬜  | —                 |
