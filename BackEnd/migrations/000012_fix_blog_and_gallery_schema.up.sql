-- Ensure author_id column exists and references users(id) instead of pengurus(id)
DO $$
BEGIN
    -- Rename from id_pengurus if exists
    IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'blog' AND column_name = 'id_pengurus') THEN
        ALTER TABLE blog RENAME COLUMN id_pengurus TO author_id;
    ELSIF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'blog' AND column_name = 'author_id') THEN
        ALTER TABLE blog ADD COLUMN author_id INT;
    END IF;

    -- Drop old foreign key if it exists (might be pointing to pengurus)
    ALTER TABLE blog DROP CONSTRAINT IF EXISTS blog_id_pengurus_fkey;
    ALTER TABLE blog DROP CONSTRAINT IF EXISTS blog_author_id_fkey;
    
    -- Add new foreign key pointing to users(id)
    ALTER TABLE blog 
        ADD CONSTRAINT blog_author_id_fkey 
        FOREIGN KEY (author_id) 
        REFERENCES users(id);
END $$;

-- Ensure thumbnail_url exists
ALTER TABLE blog ADD COLUMN IF NOT EXISTS thumbnail_url TEXT;

-- Ensure status column and type exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'blog_status') THEN
        CREATE TYPE blog_status AS ENUM ('draft', 'published', 'scheduled');
    END IF;
END $$;

ALTER TABLE blog ADD COLUMN IF NOT EXISTS status blog_status NOT NULL DEFAULT 'draft';

-- Fix foreign key constraint for gallery table
ALTER TABLE gallery DROP CONSTRAINT IF EXISTS gallery_file_upload_id_fkey;
ALTER TABLE gallery 
    ADD CONSTRAINT gallery_file_upload_id_fkey 
    FOREIGN KEY (file_upload_id) 
    REFERENCES file_uploads(id) 
    ON DELETE CASCADE;

-- Ensure kategori is a text array if it was using VARCHAR
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'blog' AND column_name = 'kategori' AND data_type = 'character varying') THEN
        ALTER TABLE blog DROP CONSTRAINT IF EXISTS blog_kategori_check;
        ALTER TABLE blog ALTER COLUMN kategori TYPE text[] USING array[kategori];
    END IF;
END $$;
