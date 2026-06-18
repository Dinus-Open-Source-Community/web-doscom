# Change Log — Sapto

Tracking commit dan perubahan lokal oleh **Sapto Gusty** (`zappto` / `saptogusty@gmail.com`).

Branch aktif: `frontend-v2`

## Index

| Tanggal | File | Tipe | Ringkasan |
| --- | --- | --- | --- |
| 2026-06-13 | [changes-2026-06-13-1be0749.md](./changes-2026-06-13-1be0749.md) | Commit | Fix Docker environment + Swagger docs |
| 2026-06-13 | [changes-2026-06-13-local-wip.md](./changes-2026-06-13-local-wip.md) | Local (belum commit) | Frontend data layer, docs, config fixes |
| 2026-06-12 | — | — | Tidak ada commit |

## Cara Penamaan

```
changes-[YYYY-MM-DD]-[commit-short-hash].md   → sudah di-commit
changes-[YYYY-MM-DD]-local-wip.md             → belum di-commit (working tree)
```

Contoh commit hash: `1be0749` (7 karakter pertama dari full hash).

## Status Working Tree (13 Jun 2026)

```
Modified:  11 files  (config, layout, QueryProvider)
Untracked: 103 files (src/lib, src/services, src/hooks, docs/)
```

Jalankan untuk update status:

```bash
git status FrontEnd/
git log --author="zappto" --since="1 day ago" --oneline
```

## Quick Reference

| Area | Path |
| --- | --- |
| Progress frontend | [../docs/progress/README.md](../docs/progress/README.md) |
| API requirements | [../docs/requirement/README.md](../docs/requirement/README.md) |
| Services | [../src/services/README.md](../src/services/README.md) |
| Hooks | [../src/hooks/README.md](../src/hooks/README.md) |
