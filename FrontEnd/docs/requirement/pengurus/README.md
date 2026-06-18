# Pengurus (Public & Self)

Modul: data pengurus public per divisi dan profil pengurus self-service.

**Hak akses:** Public (division) · `PENGURUS`, `KOOR`, `BPH`, `ADMIN` (profil sendiri)

Endpoint admin core team ada di [admin-pengurus/](../admin-pengurus/README.md).

## Functional Requirements

| ID | Requirement |
| --- | --- |
| PROF-04 | Tampilkan & edit profil pengurus jika sudah terdaftar |
| PROF-05 | Buat profil pengurus jika belum ada |

## UI Requirements — Profile Page

- Handle 404 dari `GET /pengurus/profile` → tampilkan CTA create
- Avatar upload via field `file`

### Field Permission by Role (Update Self)

| Role level | Field editable |
| --- | --- |
| Admin | Semua |
| Koordinator | name, email, divisi, position, periode, photo |
| Pengurus | name, email, divisi, periode |

## Endpoints

| Method | Path | Dokumen |
| --- | --- | --- |
| GET | `/api/v1/pengurus/division/{division}` | [GET-pengurus-by-division.md](./GET-pengurus-by-division.md) |
| GET | `/api/v1/pengurus/profile` | [GET-pengurus-profile.md](./GET-pengurus-profile.md) |
| POST | `/api/v1/pengurus` | [POST-pengurus-create-profile.md](./POST-pengurus-create-profile.md) |
| PUT | `/api/v1/pengurus/me` | [PUT-pengurus-me.md](./PUT-pengurus-me.md) |
| DELETE | `/api/v1/pengurus/me` | [DELETE-pengurus-me.md](./DELETE-pengurus-me.md) |

## Types & Hooks

- Types: `lib/types/pengurus.ts`
- Hooks: `hooks/pengurus.ts` — `usePengurusProfileQuery`, `useCreatePengurusProfileMutation`, dll.
- Service: `pengurusService`
