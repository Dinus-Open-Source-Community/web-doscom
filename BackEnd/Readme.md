# Readme — Menjalankan Go app jika `go run ./cmd/api/main.go` gagal

Dokumentasi singkat untuk menjalankan aplikasi Go di shell fish ketika perintah `go run ./cmd/api/main.go` tidak berhasil, dan solusi untuk error fish saat mencoba memuat .env:

## Masalah umum
- Error yang muncul:
    ```
    set: : invalid variable name. See `help identifiers`
    ```
    Penyebab: skrip yang membaca `.env` mencoba menjalankan `set -gx` dengan nama variabel kosong (mis. baris kosong atau komentar di `.env`). Fish tidak menerima nama variabel kosong.

## Cara cepat menjalankan aplikasi
- Coba jalankan modul package langsung (jika menggunakan module mode):
    ```
    go run ./cmd/api
    ```

## Memuat .env dengan benar di fish
Gunakan loop yang mengabaikan komentar/baris kosong dan memisahkannya dengan aman:

```fish
# load-env.fish
for line in (grep -v '^\s*#' .env | sed '/^\s*$/d')
        set parts (string split -m1 "=" $line)
        set key $parts[1]
        # gabungkan sisa jadi value (untuk = di value)
        set value (string join "=" $parts[2..-1])
        set -gx $key $value
end

# lalu jalankan aplikasi
# source load-env.fish  # di fish: . load-env.fish
go run ./cmd/api
```

Penjelasan singkat:
- `grep -v '^\s*#'` menghapus komentar.
- `sed '/^\s*$/d'` menghapus baris kosong sehingga tidak ada nama variabel kosong.
- `string split -m1 "="` memecah menjadi key dan sisa value (menghindari masalah jika `=` ada di value).

## Alternatif: muat .env dari Go
Tambahkan library seperti `github.com/joho/godotenv` ke proyek agar aplikasi yang dijalankan (tanpa perlu export manual) membaca `.env` saat startup:

```go
import "github.com/joho/godotenv"

func main() {
        _ = godotenv.Load() // load .env secara otomatis
        // ...
}
```

## Tips tambahan
- Untuk development, gunakan `direnv` atau `envchain`/`dotenv` tool agar tidak perlu script manual.
- Pastikan modul Go (go.mod) berada di root proyek sehingga `go run ./cmd/api` bekerja.
- Selalu periksa baris kosong dan komentar di `.env` saat memuat di fish.

Jika ingin, bisa tambahkan file `load-env.fish` ke repo dan perintah run singkat di README. 