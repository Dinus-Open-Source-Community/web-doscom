CREATE TABLE history_timeline (
  id SERIAL PRIMARY KEY,
  id_author INT REFERENCES users(id),
  title VARCHAR(255) NOT NULL,
  year VARCHAR(50) NOT NULL,
  description TEXT,
  display_order INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- indexing
CREATE INDEX IF NOT EXISTS idx_history_timeline_disyplay_order ON history_timeline(display_order);
