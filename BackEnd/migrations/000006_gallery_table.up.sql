CREATE TABLE IF NOT EXISTS gallery (
    id SERIAL PRIMARY KEY,
    gallery_name VARCHAR(150) NOT NULL,
    gallery_type VARCHAR(50) CHECK (gallery_type IN ('fun', 'proker', 'achievment')),
    description TEXT,
    event_date TIMESTAMP,
    image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);