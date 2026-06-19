# Upload & Media

Modul: manajemen file MinIO — dipakai silang oleh blog, gallery, work, pengurus.

## Functional Requirements

| ID | Requirement |
| --- | --- |
| UPL-01 | List file per kategori folder MinIO |
| UPL-02 | Hapus file by nama |
| UPL-03 | Pilih file existing saat create blog/work (via gallery ID / upload list) |
| UPL-04 | Upload image langsung — **endpoint POST belum aktif di backend** |

## Role Access

| Endpoint | Role |
| --- | --- |
| `GET /upload/files` | Semua user login (auth middleware tanpa role filter) |
| `DELETE /upload/file` | Semua user login |
| `POST /upload/image` | **Disabled** (commented di backend) |

## Format Respons

Upload menggunakan **flat JSON** dengan field `success`, `message` — bukan envelope standar `{ data, error }`.

## UI Requirements

| Use case | UI |
| --- | --- |
| Media library modal | Grid file dari `useUploadFilesQuery` |
| Delete orphaned file | Tombol delete + konfirmasi |
| Category tabs | Switch `gallery` / `blog` / `work` / `pengurus` |

## Ketergantungan MinIO

Upload routes hanya diregister jika MinIO client tersedia di backend (`routes.go`). Jika MinIO down, endpoint upload tidak available.

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/upload/files` | [GET-upload-files.md](./GET-upload-files.md) |
| DELETE | `/api/v1/upload/file` | [DELETE-upload-file.md](./DELETE-upload-file.md) |

## Types & Hooks

- Types: `lib/types/upload.ts`
- Service: `uploadService`
- Hooks: `useUploadFilesQuery`, `useDeleteUploadFileMutation`
