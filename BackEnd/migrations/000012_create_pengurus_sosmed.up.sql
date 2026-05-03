CREATE TABLE pengurus_sosmed (
  id SERIAL PRIMARY KEY,
  pengurus_id INT NOT NULL REFERENCES pengurus(id) ON DELETE CASCADE,
  platform VARCHAR(50) NOT NULL,       -- instagram, github, dll
  username VARCHAR(100),               -- username akun
  url TEXT,                            -- full link profile
  is_primary BOOLEAN DEFAULT FALSE,    -- akun utama atau tidak
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
