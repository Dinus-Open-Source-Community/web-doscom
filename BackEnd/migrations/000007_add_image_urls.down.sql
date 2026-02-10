-- Remove image URL columns from tables
ALTER TABLE gallery DROP COLUMN IF EXISTS image_url;
ALTER TABLE blog DROP COLUMN IF EXISTS thumbnail_url;
ALTER TABLE work DROP COLUMN IF EXISTS image_url;
ALTER TABLE pengurus DROP COLUMN IF EXISTS photo_url;

-- Drop indexes
DROP INDEX IF EXISTS idx_gallery_image_url;
DROP INDEX IF EXISTS idx_blog_thumbnail_url;
DROP INDEX IF EXISTS idx_work_image_url;
DROP INDEX IF EXISTS idx_pengurus_photo_url;
