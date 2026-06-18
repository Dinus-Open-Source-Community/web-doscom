# Shared Requirements

Kebutuhan yang berlaku untuk semua modul admin panel.

## Autentikasi Request

| Item | Requirement |
| --- | --- |
| Header | `Authorization: Bearer {access_token}` |
| Token storage | `localStorage` key `access_token` |
| Refresh token | HTTP-only cookie `RefreshToken`, path `/api/v1/auth` |
| Credentials | Axios `withCredentials: true` (untuk refresh/logout) |

Token di-set otomatis oleh `authService.login` / `authService.refresh` → `setAccessToken()`, dan interceptor `lib/axios.ts` attach header setiap request.

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

Frontend type: `PaginationQuery`, `PaginatedMeta` di `lib/types/common.ts`.

## Upload / Form Data

Endpoint create/update berikut memakai **`multipart/form-data`**:

| Modul | Field file |
| --- | --- |
| Blog (admin) | `files` (array, multiple) |
| Gallery (admin) | `files` (array, max 5) |
| Work (admin) | `files` (array, optional) |
| Pengurus (admin & self) | `file` (single, foto profil) |

Helper frontend: `toFormData()` di `lib/func/http.ts`.

## Error Handling UI

| Kebutuhan | Implementasi |
| --- | --- |
| Pesan error human-readable | `parseApiError(error)` dari `lib/message.ts` |
| Pesan sukses dari response | `translateSuccessMessage()` / `parseApiMessage()` |
| Toast UI | TBD di komponen (gunakan `UI_MESSAGES`) |

Axios interceptor melempar `ApiError` dengan `message`, `status`, `rawMessage`.

## Role-Based UI

Frontend **wajib** menyembunyikan/menonaktifkan aksi yang tidak diizinkan role user (defense in depth; backend tetap validasi).

| Fitur | Role minimum |
| --- | --- |
| Super Admin management | `SuperAdmin` |
| User CRUD | `ADMIN`, `KOOR`, `BPH` |
| Blog admin | `SuperAdmin`, `KoorMedcrev` |
| Gallery admin | `SuperAdmin`, `KoorMedcrev` |
| Work admin | `SuperAdmin`, koordinator (`RoleKoordinator`) |
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
