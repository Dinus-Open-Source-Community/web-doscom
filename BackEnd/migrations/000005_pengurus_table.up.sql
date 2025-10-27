CREATE TABLE IF NOT EXISTS pengurus (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES user(id), -- to link with user table on admin role
    id_asset INT REFERENCES gallery(id), -- to link with asset table for profile pictures
    divisi VARCHAR(50) CHECK (divisi IN ('bph', 'pemro', 'jaringan', 'medcrev', 'data')),
    name VARCHAR(150) NOT NULL,
    position VARCHAR(100) CHECK (position IN ('ketum', 'sdm','pr','pm', 'sekum', 'bendum','sek_ang','ben_ang', 'kor_pemro','kor_jaringan','kor_medcrev','kor_data','anggota','pemro_ang','jaringan_ang','medcrev_ang','data_ang')),
    period VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
