CREATE TABLE IF NOT EXISTS gallery (
    id SERIAL PRIMARY KEY,
    gallery_name VARCHAR(150) NOT NULL,
    gallery_type VARCHAR(50) CHECK (gallery_type IN ('fun', 'proker', 'achievment','work','activity','blog','pengurus')),
    description TEXT,
    event_date TIMESTAMP,
    asset_url VARCHAR(255),
    kategori VARCHAR(100) CHECK (kategori IN ('image', 'video')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);