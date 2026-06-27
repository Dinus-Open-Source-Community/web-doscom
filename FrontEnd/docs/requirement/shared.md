# Shared Requirements

Kebutuhan yang berlaku untuk semua modul admin panel.

## Autentikasi Request

| Item | Requirement |
| --- | --- |
| Cookie utama | `AccessToken` (HttpOnly, path `/`, ~15 menit) |
| Refresh token | `RefreshToken` (HttpOnly, path `/api/v1/auth`, ~2 jam) |
| Header opsional | `Authorization: Bearer {access_token}` (non-browser / legacy) |
| Credentials | Axios `withCredentials: true` (wajib untuk auth) |

Backend membaca token dari **cookie `AccessToken`** atau header `Authorization`. Setelah login/refresh, token disimpan di cookie — **bukan** di response body.

Interceptor `lib/axios.ts` masih attach Bearer dari `localStorage` jika ada (opsional); sesi utama mengandalkan cookie.

## Pagination

| Param | Tipe | Default | Keterangan |
| --- | --- | --- | --- |
| `page` | number | `1` | Halaman |
| `limit` | number | `10` | Item per halaman |

Response meta (nama field bervariasi per endpoint):

```json
{
  "totalPage": 5,
  "totalPages": 5,
  "currentPage": 1
}
```

**Catatan:** `GET /user` (list user) **tidak** mendukung pagination — mengembalikan seluruh array sesuai role.

Frontend type: `PaginationQuery`, `PaginatedMeta` di `lib/types/common.ts`.

## Upload / Form Data

Endpoint create/update berikut memakai **`multipart/form-data`**:

| Modul | Field file |
| --- | --- |
| Blog (admin) | `files` (array, multiple) |
| Gallery (admin) | `files` (array, max 5) |
| Work (admin) | `files` (array, optional), `existingID_image[]` |
| Pengurus (admin & self) | `file` (single, foto profil) |

Helper frontend: `toFormData()` di `lib/func/http.ts`; work: `buildWorkFormData()`.

## Error Handling UI

| Kebutuhan | Implementasi |
| --- | --- |
| Pesan error human-readable | `parseApiError(error)` dari `lib/message.ts` |
| Pesan sukses dari response | `translateSuccessMessage()` / `parseApiMessage()` |
| Toast UI | TBD di komponen (gunakan `UI_MESSAGES`) |

Axios interceptor melempar `ApiError` dengan `message`, `status`, `rawMessage`.

## Role-Based UI

Frontend **wajib** menyembunyikan/menonaktifkan aksi yang tidak diizinkan role user (defense in depth; backend tetap validasi).

Middleware group backend → role JWT:

| Group | Role JWT yang diizinkan |
| --- | --- |
| `ADMIN` | `SuperAdmin` |
| `KOOR` | `KoorPemro`, `KoorJaringan`, `KoorData`, `KoorMedcrev`, `BPH` |
| `BPH` | `BPH` |
| `PENGURUS` | `pemroAnggota`, `jaringanAnggota`, `medcrevAnggota`, `dataAnggota`, `BPHAnggota` |

| Fitur | Role minimum |
| --- | --- |
| Super Admin management | `SuperAdmin` (middleware `ADMIN`) |
| User CRUD | `ADMIN`, `KOOR`, `BPH` |
| Blog admin | `SuperAdmin`, `KoorMedcrev` |
| Gallery admin | `SuperAdmin`, `KoorMedcrev` |
| Work admin (CRUD) | `SuperAdmin`, koordinator (`KOOR` group) |
| Work status moderation | `SuperAdmin`, `BPH` |
| Pengurus admin | `KOOR`, `BPH`, `ADMIN` |
| Upload list/delete | Semua user login |

Ambil role dari `useMeQuery()` → field `role`.

## Hooks Global

| Kebutuhan UI | Hook |
| --- | --- |
| Data user login | `useMeQuery()` |
| Login form | `useLoginMutation()` |
| Logout sidebar | `useLogoutMutation()` |

## Status HTTP Umum

| Code | Arti untuk UI |
| --- | --- |
| 200/201 | Sukses |
| 400 | Validasi / body salah |
| 401 | Token expired / belum login → redirect login |
| 403 | Role tidak punya akses |
| 404 | Data tidak ditemukan |
| 409 | Duplikat (email, dll.) |
| 500 | Error server |

Pesan default per status ada di `HTTP_STATUS_MESSAGES` (`lib/message.ts`).

## Format Respons

### Envelope (mayoritas endpoint)

```json
{ "success": true, "message": "...", "data": {}, "error": null }
```

Berlaku untuk: auth, user, pengurus, works, gallery (public & admin).

### Flat (blog & upload)

JSON langsung tanpa field `success` — lihat modul `blog/` dan `upload/`.
