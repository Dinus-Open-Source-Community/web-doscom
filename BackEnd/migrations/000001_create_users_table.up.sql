CREATE TABLE IF NOT EXISTS users (
	 id SERIAL PRIMARY KEY,
	 email VARCHAR(100),
	 username VARCHAR(100) NOT NULL UNIQUE,
	 role VARCHAR(50) CHECK (role IN ('SuperAdmin', 'KoorPemro','KoorJaringan','KoorData','KoorMedcrev','BPH','pemroAnggota','jaringanAnggota','medcrevAnggota','dataAnggota','BPHAnggota')),
	 password VARCHAR(100) NOT NULL,
	 full_name VARCHAR(100),
	 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
