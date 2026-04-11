-- Remove added columns
ALTER TABLE work DROP COLUMN IF EXISTS gallery_id;
ALTER TABLE work DROP COLUMN IF EXISTS project_date;

-- Recreate works table if needed (optional, but good for completeness)
CREATE TABLE IF NOT EXISTS works (
    id SERIAL PRIMARY KEY,
    id_asset INT REFERENCES gallery(id),
    title VARCHAR(300) NOT NULL,
    description TEXT,
    project_date TEXT,
    pengurus_id INT REFERENCES pengurus(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
