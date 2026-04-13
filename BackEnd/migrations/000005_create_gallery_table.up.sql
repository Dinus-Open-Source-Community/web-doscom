CREATE TABLE IF NOT EXISTS gallery (
	id SERIAL PRIMARY KEY,
	id_users INT REFERENCES users(id), -- to link gallery to users who uploaded it
  file_upload_id INT REFERENCES file_uploads(id) ON DELETE CASCADE, -- fk to file_uploads table to save detail of file uploads
	gallery_name VARCHAR(150) NOT NULL,
	gallery_type VARCHAR(50) CHECK (gallery_type IN ('fun', 'proker', 'achievment','work','activity','blog','pengurus','etc')),
	description TEXT,
	event_date TIMESTAMP,
  file_url TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- indexing
-- CREATE INDEX IF NOT EXISTS idx_gallery_image_url ON gallery(image_url);
