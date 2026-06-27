# Pengurus (Public & Self)

Modul: data core team publik dan self-service profil pengurus.

## Functional Requirements

| ID | Requirement |
| --- | --- |
| PENG-PUB-01 | Public dapat list pengurus per divisi |
| PENG-SELF-01 | User login dapat buat profil pengurus sendiri |
| PENG-SELF-02 | User login dapat lihat/edit/hapus profil sendiri |
| PENG-SELF-03 | Upload foto profil (single file `file`) |

## Divisi Valid

`bph`, `pemro`, `jaringan`, `medcrev`, `data`

## Posisi Valid (`constants.ValidPosition`)

| Key | Kategori |
| --- | --- |
| `ketuaUmum` | Admin |
| `kepalaBidangSumberDayaUmum`, `projectManagerI`, `projectManagerII` | Koordinator |
| `koordinatorPemrograman`, `koordinatorJaringan`, `koordinatorMediaCreative`, `koordinatorData` | Koordinator |
| `sekretarisUmumI`, `sekretarisUmumII`, `kepalaBidangHubunganMasyarakat`, `koordinatorHubunganMasyarakatExternal` | Koordinator |
| `bendaharaUmumI`, `bendaharaUmumII` | Koordinator |
| `hubunganMasyarakatExternal`, `pemrogramanAnggota`, `jaringanAnggota`, `mediaCreativeAnggota`, `dataAnggota` | Pengurus |

## Periode

Field periode menggunakan **`start_periode_year`** dan **`end_periode_year`** (integer tahun), bukan field `period` date.

## Endpoints

| Method | Path | Auth | Dokumen |
| --- | --- | --- | --- |
| GET | `/api/v1/pengurus/division/{division}` | Public | [GET-pengurus-by-division.md](./GET-pengurus-by-division.md) |
| POST | `/api/v1/pengurus` | Login (PENGURUS+) | [POST-pengurus-create-profile.md](./POST-pengurus-create-profile.md) |
| GET | `/api/v1/pengurus/profile` | Login (PENGURUS+) | [GET-pengurus-profile.md](./GET-pengurus-profile.md) |
| PUT | `/api/v1/pengurus/me` | Login (PENGURUS+) | [PUT-pengurus-me.md](./PUT-pengurus-me.md) |
| DELETE | `/api/v1/pengurus/me` | Login (PENGURUS+) | [DELETE-pengurus-me.md](./DELETE-pengurus-me.md) |

## Types & Hooks

- Types: `lib/types/pengurus.ts`
- Service: `pengurusService.*`
- Hooks: `hooks/pengurus.ts` — `usePengurusByDivisionQuery`, `usePengurusProfileQuery`, dll.
