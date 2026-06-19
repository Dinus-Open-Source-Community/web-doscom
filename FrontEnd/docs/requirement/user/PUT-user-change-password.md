# PUT `/api/v1/user/change-password`

## Ringkasan

Ganti password akun sendiri (wajib old password).

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`, `BPH`, `PENGURUS`

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| old_password | string | True | Password lama |
| new_password | string | True | Password baru (min 8, huruf besar/kecil, angka, simbol) |

**Contoh:**

```json
{
  "old_password": "OldPass123!",
  "new_password": "NewPass456!"
}
```

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "password changed successfully, anjayy aku tau pw mu :)",
  "data": null,
  "error": null
}
```

## Response Error

**400** — body invalid / password lama salah.

## Catatan

- Backend mengembalikan status `201` (bukan `200`) meski `data` null.
- Validasi password mengikuti regex di `BackEnd/internal/constants/constants.go`.

## Frontend

- Service: `userService.changePassword`
- Hook: `useChangePasswordMutation`
- API path: `API_PATH.user.changePassword`
