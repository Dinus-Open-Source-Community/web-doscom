CREATE TYPE work_status AS ENUM (
'draft',
'published',
'unpublished', -- di un-publish manual
'rejected',
'pending_review'
);

CREATE TABLE IF NOT EXISTS work (
	id SERIAL PRIMARY KEY,
	pengurus_id INT REFERENCES pengurus(id), -- to link work to pengurus who did it for team recognition
	title VARCHAR(300) NOT NULL,
	tagline VARCHAR(100),
	description TEXT,
	slug TEXT,
	project_type varchar(50),
	technologies TEXT[],
  project_date TIMESTAMP,
  image_url TEXT,-- thumbnail
  status work_status NOT NULL DEFAULT 'draft',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- indexing
CREATE INDEX IF NOT EXISTS idx_work_image_url ON work(image_url);
