-- Add image URL columns to existing tables
ALTER TABLE gallery ADD COLUMN IF NOT EXISTS image_url VARCHAR(500);
ALTER TABLE blog ADD COLUMN IF NOT EXISTS thumbnail_url VARCHAR(500);
ALTER TABLE work ADD COLUMN IF NOT EXISTS image_url VARCHAR(500);
ALTER TABLE pengurus ADD COLUMN IF NOT EXISTS photo_url VARCHAR(500);

-- Add indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_gallery_image_url ON gallery(image_url);
CREATE INDEX IF NOT EXISTS idx_blog_thumbnail_url ON blog(thumbnail_url);
CREATE INDEX IF NOT EXISTS idx_work_image_url ON work(image_url);
CREATE INDEX IF NOT EXISTS idx_pengurus_photo_url ON pengurus(photo_url);
