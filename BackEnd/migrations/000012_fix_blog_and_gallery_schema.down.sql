ALTER TABLE blog DROP COLUMN IF EXISTS status;

ALTER TABLE gallery DROP CONSTRAINT IF EXISTS gallery_file_upload_id_fkey;
ALTER TABLE gallery 
    ADD CONSTRAINT gallery_file_upload_id_fkey 
    FOREIGN KEY (file_upload_id) 
    REFERENCES file_uploads(id);
