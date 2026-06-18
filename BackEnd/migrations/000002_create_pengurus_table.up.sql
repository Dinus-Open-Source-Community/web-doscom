CREATE TABLE IF NOT EXISTS pengurus (
  id SERIAL PRIMARY KEY ,
  id_user INT REFERENCES users(id) ON DELETE SET NULL,
  photo_url TEXT,
  email VARCHAR(100) UNIQUE NOT NULL,
  divisi VARCHAR(50) CHECK (divisi IN ('bph', 'pemro', 'jaringan', 'medcrev', 'data')),
  name VARCHAR(150) NOT NULL,
  position VARCHAR(100) CHECK (position IN ('ketuaUmum', 'kepalaBidangSumberDayaUmum','projectManagerI','projectManagerII','sekretarisUmumI', 'sekretarisUmumII', 'bendaharaUmumI','bendaharaUmumII','kepalaBidangHubunganMasyarakat','koordinatorHubunganMasyarakatExternal','koordinatorPemrograman','koordinatorData','koordinatorMediaCreative','koordinatorJaringan','hubunganMasyarakatExternal','pemrogramanAnggota','dataAnggota','jaringanAnggota','mediaCreativeAnggota')),
  start_periode_year  int NOT NULL,
  end_periode_year  int NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX IF NOT EXISTS idx_pengurus_photo_url ON pengurus(photo_url);
