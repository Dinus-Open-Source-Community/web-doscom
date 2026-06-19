# GET `/api/v1/upload/files`

## Ringkasan

List file yang sudah di-upload di folder MinIO per kategori.

## Authentication

Bearer atau cookie — semua user login (middleware auth tanpa role filter)

## Query Parameters

| Param | Tipe | Required | Keterangan |
| --- | --- | --- | --- |
| category | string | True | `gallery`, `blog`, `work`, atau `pengurus` |

## Response Success

**Status:** `200`

Format **flat** (bukan envelope standar):

```json
{
  "success": true,
  "message": "Files retrieved successfully",
  "files": [
    "2024/image-abc123.jpg",
    "2024/image-def456.png"
  ],
  "count": 2
}
```

## Response Error

**400** — category invalid / missing.

**401** — belum login.

**500** — MinIO error.

## Catatan

- Path file relatif terhadap kategori (prefix kategori di-strip dari full MinIO path).
- Dipakai untuk media picker modal saat create blog/work.

## Frontend

- Service: `uploadService.listFiles`
- Hook: `useUploadFilesQuery`
- API path: `API_PATH.upload.files`
