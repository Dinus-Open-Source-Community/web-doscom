# Blog (Public)

Modul public: list dan detail artikel blog.

**Hak akses:** Public (tanpa autentikasi)

Endpoint admin ada di [admin-blog/](../admin-blog/README.md).

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/blogs` | [GET-blogs-list.md](./GET-blogs-list.md) |
| GET | `/api/v1/blogs/{id}` | [GET-blogs-by-id.md](./GET-blogs-by-id.md) |

## Types & Hooks

- Types: `lib/types/blog.ts`
- Hooks: `hooks/blog.ts` — `useBlogsQuery`, `useBlogQuery`
- Service: `blogService`
