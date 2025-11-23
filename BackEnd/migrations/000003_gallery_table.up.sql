CREATE TABLE IF NOT EXISTS gallery (
    id SERIAL PRIMARY KEY,
    id_pengurus INT REFERENCES pengurus(id), -- to link gallery to pengurus who uploaded it
    gallery_name VARCHAR(150) NOT NULL,
    gallery_type VARCHAR(50) CHECK (gallery_type IN ('fun', 'proker', 'achievment','work','activity','blog','pengurus','etc')),
    description TEXT,
    event_date TIMESTAMP,
    file_size BIGINT,
    mime_type VARCHAR(100),
    asset_url VARCHAR(255),
    kategori VARCHAR(100) CHECK (kategori IN ('image','video')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
