CREATE TABLE IF NOT EXISTS pengurus (
    id SERIAL PRIMARY KEY,
    divisi VARCHAR(50) CHECK (divisi IN ('bph', 'pemro', 'jaringan', 'medkrev', 'data')),
    name VARCHAR(150) NOT NULL,
    position VARCHAR(100),
    image_url VARCHAR(255),
    period VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);