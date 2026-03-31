CREATE TABLE IF NOT EXISTS file_uploads (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	category VARCHAR(50) NOT NULL CHECK (category IN ('gallery', 'blog', 'work', 'pengurus')),
	original_filename VARCHAR(255) NOT NULL,
	stored_filename VARCHAR(255) NOT NULL,
	file_size BIGINT NOT NULL,
	content_type VARCHAR(100) NOT NULL,
	file_url TEXT NOT NULL,
	uploaded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- indexing
CREATE INDEX idx_file_uploads_user_id ON file_uploads(user_id);
CREATE INDEX idx_file_uploads_category ON file_uploads(category);
CREATE INDEX idx_file_uploads_uploaded_at ON file_uploads(uploaded_at DESC);
