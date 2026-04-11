CREATE TABLE IF NOT EXISTS gallery (
    id SERIAL PRIMARY KEY,
    id_users INT REFERENCES users(id), -- link ke user pengupload
    file_upload_id INT REFERENCES file_uploads(id), -- link ke metadata file
    gallery_name VARCHAR(150) NOT NULL,
    gallery_type VARCHAR(50) CHECK (gallery_type IN ('fun', 'proker', 'achievment','work','activity','blog','pengurus','etc')),
    description TEXT,
    event_date TIMESTAMP,
    file_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
