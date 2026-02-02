-- Drop trigger and function
DROP TRIGGER IF EXISTS trigger_update_file_uploads_updated_at ON file_uploads;
DROP FUNCTION IF EXISTS update_file_uploads_updated_at();

-- Drop indexes
DROP INDEX IF EXISTS idx_file_uploads_uploaded_at;
DROP INDEX IF EXISTS idx_file_uploads_category;
DROP INDEX IF EXISTS idx_file_uploads_user_id;

-- Drop table
DROP TABLE IF EXISTS file_uploads;
