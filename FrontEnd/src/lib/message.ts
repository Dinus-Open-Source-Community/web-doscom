export interface ApiMessageEnvelope {
  success?: boolean;
  message?: string;
  error?: string | Record<string, unknown>;
  errors?: string;
}

export type ApiErrorEnvelope = ApiMessageEnvelope;

export const MESSAGES = {
  error: {
    default: "Terjadi kesalahan. Silakan coba lagi.",
    network: "Tidak dapat terhubung ke server. Periksa koneksi internet Anda.",
  },
  success: {
    default: "Operasi berhasil.",
  },
  info: {
    loading: "Memuat data...",
    empty: "Belum ada data.",
  },
  warning: {
    unsavedChanges: "Perubahan belum disimpan. Yakin ingin keluar?",
  },
} as const;

export const DEFAULT_ERROR_MESSAGE = MESSAGES.error.default;
export const NETWORK_ERROR_MESSAGE = MESSAGES.error.network;
export const DEFAULT_SUCCESS_MESSAGE = MESSAGES.success.default;

export const ERROR_MESSAGES: Record<string, string> = {
  "failed to read req body":
    "Data yang dikirim tidak valid. Periksa kembali formulir Anda.",
  "invalid email or password": "Email atau kata sandi salah.",
  "invalid email or password ": "Email atau kata sandi salah.",
  "acces denied": "Akses ditolak. Akun Anda tidak memiliki izin login.",
  "invalid request body or missing fields": "Data registrasi tidak lengkap.",
  "all fields (username, email, password, role, fullname) are required":
    "Semua field wajib diisi: username, email, password, role, dan fullname.",
  "password must be at least 8 characters": "Kata sandi minimal 8 karakter.",
  "failed to register user": "Gagal mendaftarkan pengguna. Silakan coba lagi.",
  "cookie not found, what are you doing here????":
    "Sesi login tidak ditemukan. Silakan login kembali.",
  "there was an error deleting the refresh token":
    "Gagal mengakhiri sesi. Silakan coba logout lagi.",
  "refresh token not found, what are you doing heree????":
    "Token refresh tidak ditemukan. Silakan login kembali.",
  "refresh token invalid or expired":
    "Sesi Anda telah berakhir. Silakan login kembali.",

  "authroization header is required":
    "Anda belum login. Silakan login terlebih dahulu.",
  "authorization header must be bearer token":
    "Format token autentikasi tidak valid.",

  forbidden: "Anda tidak memiliki izin untuk mengakses fitur ini.",
  "you are not allowed access this resource":
    "Anda tidak memiliki izin untuk mengakses resource ini.",
  "token expired": "Sesi Anda telah berakhir. Silakan login kembali.",
  "invalid token bro": "Token autentikasi tidak valid. Silakan login kembali.",
  "invalid token": "Token autentikasi tidak valid. Silakan login kembali.",

  "role not valid, who are u??": "Role akun tidak valid.",
  "role not valid": "Role akun tidak valid.",
  "you are not allowed to access this!!":
    "Anda tidak memiliki izin untuk melakukan aksi ini.",
  "missing fields, all fields are required": "Semua field wajib diisi.",
  "failed while create user": "Gagal membuat pengguna. Silakan coba lagi.",
  "not allowed to access this route, nakal yaa!!":
    "Anda tidak memiliki izin mengakses halaman ini.",
  "some field are missing": "Beberapa field wajib belum diisi.",
  "email already registered": "Email sudah terdaftar.",
  "failed to create user": "Gagal membuat pengguna. Silakan coba lagi.",
  "invalid user id": "ID pengguna tidak valid.",
  "user not found": "Pengguna tidak ditemukan.",
  "unable to process the request": "Permintaan tidak dapat diproses.",
  "failed to fetch users data": "Gagal mengambil data pengguna.",
  "role not valid to proced this action":
    "Role Anda tidak diizinkan melakukan aksi ini.",
  "invalid id": "ID tidak valid.",
  "invalild id": "ID tidak valid.",
  "failed to update data user": "Gagal memperbarui data pengguna.",
  "cannot proced": "Aksi tidak dapat dilanjutkan.",
  "you are not allowed": "Anda tidak memiliki izin untuk aksi ini.",
  "error while delete user": "Gagal menghapus pengguna.",
  "failed to change password": "Gagal mengubah kata sandi.",
  "you are not allowed to access this route,, hayo ngapain masuk sini":
    "Anda tidak memiliki izin mengakses halaman ini.",
  "failed to change password, nakal ganti pw wong liyo, mentang mentang superadmin":
    "Gagal mengubah kata sandi pengguna.",
  "terjadi kesalahan ketika ambil data":
    "Gagal mengambil data. Silakan coba lagi.",

  forbiddennnnn: "Anda tidak memiliki izin untuk mengakses blog ini.",
  "failed to read file": "Gagal membaca file yang diunggah.",
  "failed to insert data": "Gagal menyimpan data blog.",
  "failed to fetch blog data": "Gagal mengambil data blog.",
  "blog not found": "Artikel blog tidak ditemukan.",
  "invalid id blog": "ID blog tidak valid.",
  "failed to update blog, something went wrong":
    "Gagal memperbarui blog. Silakan coba lagi.",
  "max 3 kategory allowed": "Maksimal 3 kategori diperbolehkan.",
  "failed to fetch data": "Gagal mengambil data.",
  "terjadi kesalahan ketika mengambil data":
    "Gagal mengambil data. Silakan coba lagi.",
  "terjadi kesalahan ketika menghapus data :)":
    "Gagal menghapus data. Silakan coba lagi.",

  "you're not allowed broo": "Anda tidak memiliki izin mengunggah galeri.",
  "no files uploaded": "Tidak ada file yang diunggah.",
  "max upload 5 file": "Maksimal 5 file dapat diunggah sekaligus.",
  "failed while opening file upload": "Gagal memproses file upload.",
  "failed to insert and upload gallery": "Gagal menyimpan galeri.",
  "failed to get data gallery, some issue at backend :))":
    "Gagal mengambil data galeri.",
  "you cannot access this resource":
    "Anda tidak memiliki izin mengakses galeri ini.",
  "request cannot be processed": "Permintaan tidak dapat diproses.",
  "role not valid or have no permission":
    "Role Anda tidak memiliki izin untuk aksi ini.",
  "gallery not found": "Galeri tidak ditemukan.",

  "role is not allowed to access this resource":
    "Anda tidak memiliki izin mengakses data proyek.",
  "invalid request body": "Data proyek tidak valid.",
  "failed to insert data, something went wrong":
    "Gagal menyimpan proyek. Silakan coba lagi.",
  "failed to fetch data, something went wrong": "Gagal mengambil data proyek.",
  "invalid id format": "Format ID proyek tidak valid.",
  "something went wrong while fetching data": "Gagal mengambil detail proyek.",
  "forbidden to access this resource":
    "Anda tidak memiliki izin mengakses proyek ini.",
  "invalid request body, budy": "Data permintaan tidak valid.",
  "failed to update work, something went wrong": "Gagal memperbarui proyek.",
  "failed to delete work, something went wrong": "Gagal menghapus proyek.",

  "failed to open file": "Gagal membuka file yang diunggah.",
  "gagal mendaftarkan pengurus: email sudah digunakan":
    "Email sudah digunakan oleh pengurus lain.",
  "gagal mendaftarkan pengurus: input tidak valid":
    "Data pengurus tidak valid. Periksa kembali formulir Anda.",
  "failed to create data pengurus, server error":
    "Gagal menyimpan data pengurus.",
  "invalid pengurus id": "ID pengurus tidak valid.",
  "error while getting the data or data not found":
    "Data pengurus tidak ditemukan.",
  "failed to get data": "Gagal mengambil data pengurus.",
  "failed to fetch pengurus data": "Gagal mengambil data pengurus.",
  "failed to get data somting wong": "Gagal mengambil data pengurus.",
  "failed to bind data, please use form-data for updates":
    "Format data tidak valid. Gunakan form-data untuk pembaruan.",
  "akses ditolak untuk memperbarui data ini":
    "Anda tidak memiliki izin memperbarui data ini.",
  "pengurus data not found": "Data pengurus tidak ditemukan.",
  "failed to update data pengurus, server error":
    "Gagal memperbarui data pengurus.",
  "pengurus not found": "Pengurus tidak ditemukan.",

  "failed to delete file": "Gagal menghapus file.",
  "failed to list files": "Gagal mengambil daftar file.",

  "email sudah terdaftar di pengurus":
    "Email sudah digunakan oleh pengurus lain.",
  "user_id tidak ditemukan": "Pengguna tidak ditemukan.",
  "koordinator tidak dapat memperbarui foto pengurus":
    "Koordinator tidak dapat memperbarui foto pengurus.",
  "forbidden, you can't see other data":
    "Anda tidak memiliki izin melihat data ini.",
  "you can't see other data bro": "Anda tidak memiliki izin melihat data ini.",
  "you can not see other division bro":
    "Anda tidak memiliki izin melihat divisi lain.",
  "you can not delete data from other division":
    "Anda tidak dapat menghapus data dari divisi lain.",
  "you can't delete your own data": "Anda tidak dapat menghapus data sendiri.",
  "work not found or something wrong while fetch data":
    "Proyek tidak ditemukan.",
  "work not found, the process cannot continue": "Proyek tidak ditemukan.",
  "you can only set 5 gallery to this work":
    "Maksimal 5 gambar galeri per proyek.",
  "invalid gallery ids": "ID galeri tidak valid.",
  "pengurus tidak valid atau tidak ditemukan":
    "Pengurus tidak valid atau tidak ditemukan.",
};

export const SUCCESS_MESSAGES: Record<string, string> = {
  "login success bolo, nasi padang satu bungkus": "Login berhasil.",
  "login success, nasi padangnya sebungkus bolo": "Login berhasil.",
  "user created successfully": "Pengguna berhasil dibuat.",
  "logout success, nasi padang satu bungkus": "Logout berhasil.",
  "refresh token success": "Sesi berhasil diperbarui.",

  "superadmin created successfully": "Super admin berhasil dibuat.",
  "get user": "Data pengguna berhasil diambil.",
  "list of users data": "Daftar pengguna berhasil diambil.",
  "successfully update user data": "Data pengguna berhasil diperbarui.",
  "user deleted info kopi dan gorengan bolo": "Pengguna berhasil dihapus.",
  "password changed successfully": "Kata sandi berhasil diubah.",
  "password changed successfully, anjayy aku tau pw mu :)":
    "Kata sandi berhasil diubah.",
  "get super admin data": "Data super admin berhasil diambil.",
  "current user data": "Data profil berhasil diambil.",

  "successfully create blog": "Blog berhasil dibuat.",
  "succsess get all blogs": "Data blog berhasil diambil.",
  "blog found successfully": "Blog berhasil ditemukan.",
  "successfully update blog": "Blog berhasil diperbarui.",
  "successfully fetch data": "Data berhasil diambil.",
  "successfully delete data": "Blog berhasil dihapus.",

  "successfully insert data": "Galeri berhasil disimpan.",
  "successfully get data": "Data berhasil diambil.",
  "successfully delete gallery, bang nasi padang satu bungkus bang":
    "Galeri berhasil dihapus.",

  "work created successfully": "Proyek berhasil dibuat.",
  "success fetching all works": "Data proyek berhasil diambil.",
  "success fetching work detail": "Detail proyek berhasil diambil.",
  "work updated successfully": "Proyek berhasil diperbarui.",
  "work deleted successfully": "Proyek berhasil dihapus.",

  "pengurus created successfully": "Data pengurus berhasil dibuat.",
  success: "Operasi berhasil.",
  "list of pengurus": "Daftar pengurus berhasil diambil.",
  "successfully update pengurus data": "Data pengurus berhasil diperbarui.",
  "pengurus deleted": "Data pengurus berhasil dihapus.",

  "file deleted successfully": "File berhasil dihapus.",
  "files retrieved successfully": "Daftar file berhasil diambil.",
  "file uploaded successfully": "File berhasil diunggah.",
};

export const UI_MESSAGES = {
  auth: {
    loginSuccess: "Login berhasil. Selamat datang!",
    logoutSuccess: "Logout berhasil.",
    registerSuccess: "Registrasi berhasil.",
    sessionExpired: "Sesi Anda telah berakhir. Silakan login kembali.",
  },
  common: {
    saveSuccess: "Data berhasil disimpan.",
    updateSuccess: "Data berhasil diperbarui.",
    deleteSuccess: "Data berhasil dihapus.",
    loadError: "Gagal memuat data. Silakan coba lagi.",
    confirmDelete: "Yakin ingin menghapus data ini?",
  },
  blog: {
    createSuccess: "Blog berhasil dibuat.",
    updateSuccess: "Blog berhasil diperbarui.",
    deleteSuccess: "Blog berhasil dihapus.",
  },
  gallery: {
    createSuccess: "Galeri berhasil disimpan.",
    deleteSuccess: "Galeri berhasil dihapus.",
  },
  work: {
    createSuccess: "Proyek berhasil dibuat.",
    updateSuccess: "Proyek berhasil diperbarui.",
    deleteSuccess: "Proyek berhasil dihapus.",
  },
  pengurus: {
    createSuccess: "Data pengurus berhasil dibuat.",
    updateSuccess: "Data pengurus berhasil diperbarui.",
    deleteSuccess: "Data pengurus berhasil dihapus.",
  },
  user: {
    createSuccess: "Pengguna berhasil dibuat.",
    updateSuccess: "Data pengguna berhasil diperbarui.",
    deleteSuccess: "Pengguna berhasil dihapus.",
    passwordChanged: "Kata sandi berhasil diubah.",
  },
  upload: {
    deleteSuccess: "File berhasil dihapus.",
  },
} as const;

export const ERROR_PATTERNS: Array<{ pattern: RegExp; message: string }> = [
  {
    pattern: /validation failed/i,
    message: "Validasi gagal. Periksa kembali data yang Anda isi.",
  },
  {
    pattern: /token expired|jwt.*expired/i,
    message: "Sesi Anda telah berakhir. Silakan login kembali.",
  },
  {
    pattern: /invalid token/i,
    message: "Token autentikasi tidak valid. Silakan login kembali.",
  },
  {
    pattern: /not found|tidak ditemukan/i,
    message: "Data tidak ditemukan.",
  },
  {
    pattern: /forbidden|not allowed|akses ditolak|you are not allowed/i,
    message: "Anda tidak memiliki izin untuk aksi ini.",
  },
  {
    pattern: /unauthorized|authorization header/i,
    message: "Anda belum login. Silakan login terlebih dahulu.",
  },
  {
    pattern: /duplicate|already registered|sudah terdaftar|sudah digunakan/i,
    message: "Data sudah terdaftar. Gunakan data lain.",
  },
  {
    pattern: /password/i,
    message: "Terjadi masalah pada kata sandi. Periksa kembali input Anda.",
  },
];

export const SUCCESS_PATTERNS: Array<{ pattern: RegExp; message: string }> = [
  {
    pattern: /created successfully|successfully create|berhasil dibuat/i,
    message: "Data berhasil dibuat.",
  },
  {
    pattern: /updated successfully|successfully update|berhasil diperbarui/i,
    message: "Data berhasil diperbarui.",
  },
  {
    pattern: /deleted successfully|successfully delete|berhasil dihapus/i,
    message: "Data berhasil dihapus.",
  },
  {
    pattern: /successfully fetch|success fetching|successfully get|get data/i,
    message: "Data berhasil diambil.",
  },
  {
    pattern: /login success|refresh token success/i,
    message: "Autentikasi berhasil.",
  },
  {
    pattern: /logout success/i,
    message: "Logout berhasil.",
  },
];

export const HTTP_STATUS_MESSAGES: Record<number, string> = {
  400: "Permintaan tidak valid. Periksa kembali data yang dikirim.",
  401: "Sesi Anda telah berakhir. Silakan login kembali.",
  403: "Anda tidak memiliki izin untuk aksi ini.",
  404: "Data tidak ditemukan.",
  409: "Data sudah ada. Gunakan data lain.",
  422: "Data yang dikirim tidak valid.",
  429: "Terlalu banyak permintaan. Coba lagi nanti.",
  500: "Terjadi kesalahan pada server. Silakan coba lagi.",
  502: "Server sedang tidak tersedia. Coba lagi nanti.",
  503: "Layanan sedang sibuk. Coba lagi nanti.",
};

export {
  ApiError,
  normalizeErrorKey,
  normalizeMessageKey,
  parseApiError,
  toApiError,
  translateErrorMessage,
} from "./func/error";
export {
  parseApiMessage,
  translateApiMessage,
  translateSuccessMessage,
} from "./func/message";
