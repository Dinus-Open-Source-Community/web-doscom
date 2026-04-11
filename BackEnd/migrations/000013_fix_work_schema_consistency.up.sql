-- Add missing columns to 'work' table
ALTER TABLE work ADD COLUMN IF NOT EXISTS gallery_id INT REFERENCES gallery(id) ON DELETE SET NULL;
ALTER TABLE work ADD COLUMN IF NOT EXISTS project_date TIMESTAMP;

-- Rename or ensure pengurus_id is consistent
-- (It already exists from migration 000004)

-- Clean up the plural 'works' table to avoid confusion
DROP TABLE IF EXISTS works CASCADE;
