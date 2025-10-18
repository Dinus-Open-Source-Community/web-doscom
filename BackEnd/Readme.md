
# Web Doscom Backend — Dokumentasi

## Struktur Proyek

- `cmd/api/main.go` — Entry point aplikasi Gin API
- `internal/server/routes.go` — Semua definisi route API (ping, user, dst)
- `internal/handler/` — Handler (jika ingin custom handler)
- `internal/database/` — Koneksi dan model database (GORM)
- `.env` — Konfigurasi environment (DB, PORT, dll)

## Setup & Menjalankan

1. **Clone repo & install dependency**
    ```sh
    git clone https://github.com/Dinus-Open-Source-Community/web-doscom.git
    cd web-doscom/BackEnd
    go mod tidy
    ```

2. **Buat file `.env`**
    ```env
    PORT=3001
    DBURL=postgres://webdoscom:webdoscom123@localhost:5432/web_doscom?sslmode=disable
    DB_TIMEZONE=public
    # ...tambahkan variabel lain sesuai kebutuhan
    ```

3. **Jalankan aplikasi**
    ```sh
    go run ./cmd/api/main.go
    # atau jika pakai fish shell dan ingin load .env manual:
    # . load-env.fish; go run ./cmd/api/main.go
    # cd to backend folder go run ./cmd.api/mai.go
    ```

## Endpoints API

- **GET `/api/ping`**
    - Cek server hidup
    - Response: `{ "message": "pong" }`

- **POST `/api/user`**
    - Contoh create user (dummy, belum ke database)
    - Response:
      ```json
      {
        "message": "User created successfully",
        "user": {
          "id": 1,
          "name": "Test User"
        }
      }
      ```

## Tips & Troubleshooting

- Jika error `.env` tidak terbaca, pastikan `env.LoadEnv()` dipanggil sebelum koneksi DB.
- Jika port sudah dipakai, ganti `PORT` di `.env` dan restart server.
- Untuk error fish shell saat load .env, gunakan script `load-env.fish` seperti di bawah:

```fish
for line in (grep -v '^s*#' .env | sed '/^s*$/d')
    set parts (string split -m1 "=" $line)
    set key $parts[1]
    set value (string join "=" $parts[2..-1])
    set -gx $key $value
end
```

- Untuk pengembangan, gunakan [direnv](https://direnv.net/) agar .env otomatis di-load.

## Kontribusi

1. Fork & clone repo
2. Buat branch baru
3. Commit perubahan
4. Push dan buat Pull Request

---
Maintainer: Dinus Open Source Community