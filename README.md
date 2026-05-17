# 🧩 Web Doscom Project - Backend & Frontend

Web Doscom adalah aplikasi web modern untuk **Dinus Open Source Community (Doscom)** yang dibangun dengan arsitektur **Monolith**. Proyek ini memisahkan Backend (Golang) dan Frontend (Astro) yang saling berkomunikasi via REST API.

---

## Teknologi yang Digunakan

| Komponen | Teknologi | Deskripsi |
| :--- | :--- | :--- |
| **Backend** | ![Golang](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white) **Golang (Gin)** | RESTful API, Auth (JWT), Services |
| **Frontend** | ![Astro](https://img.shields.io/badge/Astro-BC52EE?style=flat&logo=astro&logoColor=white) **Astro** | Modern Static & Server Side Rendering |
| **Database** | ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=flat&logo=postgresql&logoColor=white) **PostgreSQL** | Relational Database Management |
| **Storage** | ![MinIO](https://img.shields.io/badge/MinIO-C72E49?style=flat&logo=minio&logoColor=white) **MinIO** | S3 Compatible Object Storage (File Uploads) |
| **Container** | ![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white) **Docker Compose** | Orchestration & Deployment |

---

##  Struktur Direktori

```bash
web-doscom/
│
├── BackEnd/               # Source code Backend (Go)
│   ├── cmd/api/           # Entry point (main.go)
│   ├── internal/          # Core Logic (Handler, Service, Model)
│   ├── migrations/        # SQL Migration files
│   └── Dockerfile         # Config Docker Backend
│
├── FrontEnd/              # Source code Frontend (Astro)
│   ├── src/               # Pages, Components, Layouts
│   └── Dockerfile         # Config Docker Frontend
│
├── docker-compose.yml     # Konfigurasi Multi-Container
└── perbaikan.md           # 📘 Panduan Teknis & Instruksi Tim (WAJIB BACA)
```

---

##  Cara Menjalankan Aplikasi

Pastikan PC Anda sudah terinstal **Docker** dan **Docker Compose**.

### 1. Clone & Masuk Direktori
```bash
git clone https://github.com/doscom/web-doscom.git
cd web-doscom
```

### 2. Jalankan Environment (Sekali Perintah)
Cukup jalankan satu perintah ini untuk menyalakan Backend, Frontend, Database, dan MinIO sekaligus:

```bash
docker compose up -d --build
```

*Tunggu beberapa saat hingga semua container `healthy` dan database selesai inisialisasi.*

### 3. Cek Status
Pastikan semua container berjalan:
```bash
docker ps
```

---

##  Akses & Kredensial (PENTING)

Setelah aplikasi berjalan, Anda bisa mengakses layanan berikut:

###  Aplikasi Utama
| Layanan | URL | Keterangan |
| :--- | :--- | :--- |
| **Frontend Website** | [http://localhost:4321](http://localhost:4321) | Halaman utama yang diakses user |
| **Backend API** | [http://localhost:8080](http://localhost:8080) | Base URL API |
| **Swagger Docs** | [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) | **Dokumentasi API Lengkap** |
| **MinIO Console** | [http://localhost:9001](http://localhost:9001) | Dashboard Cloud Storage |
| **Adminer** | [http://localhost:8081](http://localhost:8081) | Database GUI (Optional) |

### 👤 Akun Admin Default (Super Admin)
Gunakan akun ini untuk login pertama kali dan mengelola sistem:

*   **Username**: `admin`
*   **Password**: `admin123`
*   **Email**: `admin@doscom.org`

>  **Catatan**: Password ini di-seed otomatis saat pertama kali docker dijalankan. Segera ganti password di production!

---

##  Endpoint API Tersedia

Detail lengkap bisa dilihat di **Swagger UI**, berikut ringkasannya:

*   `POST /api/v1/auth/login` - Login User
*   `POST /api/v1/gallery` - Upload Galeri (MinIO)
*   `POST /api/v1/pengurus` - Register Pengurus (MinIO)
*   `POST /api/v1/blogs` - Create Blog Articles
*   `GET /api/v1/work` - List Works
*   ... dan lainnya.

---

##  Troubleshooting

**Backend restart terus?**
Cek log error-nya:
```bash
docker logs doscom-backend -f
```

**Database connection refused?**
Pastikan container `doscom-db` sudah `healthy`. Jika baru pertama kali run, inisialisasi database butuh waktu ~30 detik.

**Menghapus semua data (Reset)?**
```bash
docker compose down -v
```

---

*Dibuat oleh Tim Doscom Developers - 2025*
