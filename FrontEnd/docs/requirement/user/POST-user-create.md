# POST `/api/v1/user`

## Ringkasan

Buat user baru oleh admin/koordinator/BPH. Role default di-assign backend berdasarkan creator role.

## Authentication

Bearer atau cookie — middleware `ADMIN`, `KOOR`, `BPH`

Role `PENGURUS` ditolak (**403**).

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| email | string | True | Email unik |
| fullname | string | True | Nama lengkap |
| username | string | False | Username (default dari fullname) |
| password | string | False | Password (default generated jika kosong) |
| role | string | False | Role key JWT |

**Contoh:**

```json
{
  "email": "111202412345@mhs.dinus.ac.id",
  "fullname": "Anggota Pemro",
  "role": "pemroAnggota"
}
```

## Response Success

**Status:** `201`

Format envelope `{ success, message, data, error }`.

```json
{
  "success": true,
  "message": "User created successfully",
  "data": null,
  "error": null
}
```

## Response Error

**400** — field wajib kosong.

**403** — role caller tidak diizinkan.

**409** — email sudah terdaftar.

## Frontend

- Service: `userService.create`
- Hook: `useCreateUserMutation`
- API path: `API_PATH.user.list`
