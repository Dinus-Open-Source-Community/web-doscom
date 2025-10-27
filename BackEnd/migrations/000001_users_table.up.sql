CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100), 
    username VARCHAR(100) NOT NULL UNIQUE,
    role VARCHAR(50) CHECK (role IN ('Super_Admin', 'Kor_Pemro','Kor_Jaringan','Kor_Data','Kor_Medcrev','BPH','pemro_ang','jaringan_ang','medcrev_ang','data_ang','BPH_ang')),
    password VARCHAR(100) NOT NULL,
    full_name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- pemro, jarignan , data, medcrev, bph