CREATE TYPE blog_category AS ENUM (
'event',
'seminar',
'collaboration',
'education',
'technology',
'work',
'activity',
'sharing'
);

CREATE TYPE blog_status AS ENUM (
'draft',
'published',
'scheduled'
);

CREATE TABLE blog (
	id SERIAL PRIMARY KEY,
	author_id INT REFERENCES pengurus(id),
	title VARCHAR(255) NOT NULL,
	slug VARCHAR(255) NOT NULL UNIQUE,
	content TEXT,
  thumbnail_url TEXT,
	kategori blog_category[] NOT NULL,
	published_at TIMESTAMP,
	status blog_status NOT NULL DEFAULT 'draft',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- indexing
CREATE INDEX IF NOT EXISTS idx_blog_thumbnail_url ON blog(thumbnail_url);
