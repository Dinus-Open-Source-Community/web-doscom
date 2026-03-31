CREATE TABLE history_photos (
  id SERIAL PRIMARY KEY,
  id_history INT REFERENCES history_timeline(id),
  imager_url TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- indexing
CREATE INDEX IF NOT EXISTS idx_history_photos_history_id ON history_photos(id_history);
