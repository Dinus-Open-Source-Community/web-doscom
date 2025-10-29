CREATE TABLE IF NOT EXISTS pengurus (
    id CHAR(8) PRIMARY KEY ,
    id_user INT REFERENCES users(id), -- to link with user table on admin role
    URLAsset TEXT, -- get api to asset table gallery for profile pictures
    email VARCHAR(100) UNIQUE NOT NULL,
    divisi VARCHAR(50) CHECK (divisi IN ('bph', 'pemro', 'jaringan', 'medcrev', 'data')),
    name VARCHAR(150) NOT NULL,
    position VARCHAR(100) CHECK (position IN ('ketum', 'sdm','pr','pm','pm_ang', 'sekum', 'bendum','sek_ang','ben_ang', 'kor_pemro','kor_jaringan','kor_medcrev','kor_data','anggota','pemro_ang','jaringan_ang','medcrev_ang','data_ang')),
    sosmed ENUM('instagram', 'linkedin', 'github'),
    period DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
