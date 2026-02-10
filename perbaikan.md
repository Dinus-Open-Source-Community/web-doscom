# Panduan Perubahan & Pengembangan Backend Web Doscom 

Dokumen ini berisi rangkuman perubahan yang telah dilakukan (Migrasi MinIO & Refactoring) serta panduan bagi tim pengembang untuk melanjutkan pengembangan fitur selanjutnya dan juga sebagai dokumentasi.

---

##  1. Apa Saja yang Telah Diubah?

Berikut adalah komponen utama yang telah diperbarui dan fungsinya:

### A. Centralized Storage Service (`internal/service/storage.go`)
*   **Fungsi**: Service pusat untuk menangani semua upload file ke MinIO.
*   **Kenapa dibuat?**: Supaya tidak duplikasi code upload di setiap handler. Semua fitur (Gallery, Pengurus, Blog) wajib lewat service ini.
*   **Fitur**: Auto-rename file (UUID), validasi tipe & size file aman, retry logic, dan logging ke database `file_uploads`.

### B. Arsitektur Baru (Code Pattern)
Kami menerapkan pola **Clean Architecture** yang lebih rapi:
`Handler` ⇒ `Service` ⇒ `Model (Repository)`

*   **Service Layer**: Logic bisnis (seperti validasi, panggil storage service) dipindah ke sini. Handler hanya fokus terima request dan return response.
*   **Contoh Penerapan**:
    *   `Gallery`: sudah pakai `GalleryService`
    *   `Pengurus`: sudah pakai `PengurusService`
    *   `Blog`: sudah pakai `BlogService`

### C. Fitur Blog (`internal/handler/blog.go` & `service/blog_service.go`)
*   Fitur ini sebelumnya belum komplit, sekarang sudah **FULLY IMPLEMENTED**.
*   **CRUD Lengkap**: Create, List, GetByID, Update, Delete, dan ListByKategori.
*   **Swagger**: Dokumentasi API sudah update dan bisa dites.

### D. Perbaikan Crash & Type Safety
*   Memperbaiki bug `strconv` di `UploadHandler` yang bikin backend restart terus menerus.
*   Memperbaiki validasi `user_id` agar konsisten pakai tipe `int` (bukan string) sesuai middleware Auth.

---

## 🛠 2. Instruksi Penggunaan (Wajib Dibaca Tim)

Jika tim ingin membuat fitur baru atau melanjutkan fitur **Work** dan **User**, **WAJIB** mengikuti alur ini agar code tetap bersih dan tidak "spaghetti".

### Alur Kerja (Workflow)
Saat membuat fitur baru, jangan langsung coding di Handler! Ikuti urutan ini:

#### 1. Mulai dari Model (`internal/database/model/`)
Pastikan struct database dan interface ke GORM sudah siap dan support method yang dibutuhkan.
*   *Contoh*: Cek `internal/database/model/blog.go` untuk referensi.

#### 2. Buat Service Layer (`internal/service/`)  **PENTING**
Jangan akses DB langsung dari Handler! Buat file service baru.
*   Buat file misal: `internal/service/work_service.go`.
*   Inject model ke dalam service struct.
*   Masukkan logic bisnis di sini.
*   *Contoh*: Lihat `internal/service/blog_service.go`.

#### 3. Buat/Update Handler (`internal/handler/`)
Handler sekarang harusnya "bodoh" (hanya terima input & return JSON).
*   Inject Service ke dalam Handler struct (bukan DB langsung).
*   Panggil method service.
*   *Jangan ada logic upload manual di sini*, panggil `StorageService`.

#### 4. Update Routes (`internal/routes/routes.go`)
*   Inject dependency di `routes.go`: Model -> Service -> Handler.
*   Daftarkan route baru.

### Contoh Code Pattern (Copy-Paste Friendly)

**Service Pattern:**
```go
type WorkService struct {
    Model *model.WorkModel
}
func NewWorkService(m *model.WorkModel) *WorkService {
    return &WorkService{Model: m}
}
// Logic function
func (s *WorkService) CreateWork(data *model.Work) error {
    return s.Model.InsertWork(data)
}
```

**Handler Injection Pattern:**
```go
type WorkHandler struct {
    Service *service.WorkService // Pakai Service, jangan DB!
}
func NewWorkHandler(s *service.WorkService) *WorkHandler {
    return &WorkHandler{Service: s}
}
```

---

**Summary Akhir:**
Backend sekarang sudah **STABLE**. Swagger Documentation di endpoint `/swagger/index.html` adalah sumber kebenaran (Source of Truth) untuk frontend developer tapi tolong api tolong dites dulu sebelum digunakan (wajib).
