

---

# 🧩 Web Doscom Project

Proyek ini merupakan aplikasi web **Doscom** yang terdiri dari tiga komponen utama:

* **Frontend** — dibangun menggunakan [Astro.js](https://astro.build/) dan Node.js
* **Backend** — ditulis menggunakan [Golang](https://go.dev/)
* **Database** — menggunakan [PostgreSQL](https://www.postgresql.org/)

---

## 📁 Struktur Proyek

```
web-doscom/
│
├── BackEnd/           # Source code backend Go
│   ├── cmd/           # Entry point aplikasi
│   ├── internal/      # Package internal backend
│   ├── migrations/    # File migrasi database
│   ├── go.mod
│   └── Dockerfile
│
├── FrontEnd/          # Source code frontend Astro
│   ├── public/
│   ├── src/
│   ├── package.json
│   └── Dockerfile
│
└── docker-compose.yml # Konfigurasi multi-container Docker
```

---

## 🚀 Menjalankan Proyek

### 1. Pastikan telah terinstal:

* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)

Cek versi:

```bash
docker -v
docker compose version
```

---

### 2. Clone repository ini

```bash
git clone https://github.com/<username>/web-doscom.git
cd web-doscom
```

---

### 3. Jalankan seluruh service

Gunakan perintah berikut untuk menjalankan **database**, **backend**, dan **frontend** sekaligus:

```bash
docker compose up -d
```

🧠 **Penjelasan singkat service:**

* `db` → PostgreSQL (port: 5432)
* `backend` → Go API Server (port: 8080)
* `frontend` → Astro Frontend (port: 4321)

---

### 4. Akses aplikasi

| Komponen    | URL Akses                                      | Keterangan                                 |
| ----------- | ---------------------------------------------- | ------------------------------------------ |
| Frontend    | [http://localhost:4321](http://localhost:4321) | Aplikasi utama                             |
| Backend API | [http://localhost:8080](http://localhost:8080) | Endpoint API Go                            |

---

### 5. Melihat log container

Untuk melihat log backend:

```bash
docker logs doscom-backend -f
```

Untuk melihat log frontend:

```bash
docker logs doscom-frontend -f
```

Untuk melihat log database:

```bash
docker logs doscom-db -f
```

---

### 6. Menghentikan semua container

```bash
docker compose down
```

Jika ingin juga menghapus volume (data database):

```bash
docker compose down -v
```

---

## ⚙️ Variabel Lingkungan (Environment Variables)

| Variabel              | Service  | Nilai Default | Keterangan                 |
| --------------------- | -------- | ------------- | -------------------------- |
| `POSTGRES_USER`       | db       | doscom        | Username database          |
| `POSTGRES_PASSWORD`   | db       | doscom        | Password database          |
| `POSTGRES_DB`         | db       | doscomdb      | Nama database              |
| `DB_HOST`             | backend  | db            | Host PostgreSQL            |
| `DB_PORT`             | backend  | 5432          | Port PostgreSQL            |
| `DB_USER`             | backend  | doscom        | Username PostgreSQL        |
| `DB_PASS`             | backend  | doscom        | Password PostgreSQL        |
| `DB_NAME`             | backend  | doscomdb      | Nama database              |
| `PORT`                | backend  | 8080          | Port aplikasi backend      |
| `CHOKIDAR_USEPOLLING` | frontend | true          | Untuk hot reload di Docker |

---

## 🧹 Tips Pengembangan

* Gunakan `docker compose logs -f` untuk debugging cepat.
* Perubahan pada file frontend/backend akan otomatis direload karena volume mount (`./FrontEnd:/usr/src/app` dan `./BackEnd:/app`).
* Jika dependency berubah (misal `go.mod` atau `package.json` diubah), rebuild image:

  ```bash
  docker compose build
  ```

---

## 📄 Lisensi

Proyek ini dikembangkan untuk keperluan internal **Doscom (Dinus Open Source Community)**.
Hak cipta © 2025 Doscom Developers.

---

Apakah kamu mau saya tambahkan contoh **endpoint API** backend (misalnya `/api/health` atau `/api/users`) di bagian dokumentasi bawahnya juga? Itu akan membantu kalau README-nya mau sekaligus jadi dokumentasi developer.
