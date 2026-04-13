CREATE TABLE IF NOT EXISTS work (
	id SERIAL PRIMARY KEY,
	pengurus_id INT REFERENCES pengurus(id), -- to link work to pengurus who did it for team recognition
	title VARCHAR(300) NOT NULL,
	tagline VARCHAR(100),
	description TEXT,
	slug TEXT,
	project_type varchar(50),
	technologies TEXT[],
  image_url TEXT,-- thumbnail
    gallery_id INT REFERENCES gallery(id) ON DELETE SET NULL,
    project_date TIMESTAMP
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- indexing
CREATE INDEX IF NOT EXISTS idx_work_image_url ON work(image_url);
