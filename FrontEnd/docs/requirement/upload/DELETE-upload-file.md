# DELETE `/api/v1/upload/file`

## Ringkasan

Hapus file dari MinIO storage by nama file.

## Authentication

Bearer atau cookie — semua user login

## Headers

| Header | Nilai |
| --- | --- |
| Content-Type | `application/json` |

## Request Body

| Field | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| file_name | string | True | Nama/path file di MinIO |

**Contoh:**

```json
{
  "file_name": "gallery/2024/image-abc123.jpg"
}
```

## Response Success

**Status:** `200`

Format **flat**:

```json
{
  "success": true,
  "message": "File deleted successfully"
}
```

## Response Error

**400** — body invalid.

**500** — gagal hapus di MinIO.

## Catatan

- Kirim body via axios `{ data: payload }` (DELETE with body).
- Hati-hati: file yang masih direferensi gallery/blog/work bisa menyebabkan broken link.

## Frontend

- Service: `uploadService.deleteFile`
- Hook: `useDeleteUploadFileMutation`
- API path: `API_PATH.upload.file`
