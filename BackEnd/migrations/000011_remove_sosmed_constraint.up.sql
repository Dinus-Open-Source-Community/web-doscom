-- Hapus constraint lama yang mematikan
ALTER TABLE pengurus DROP CONSTRAINT IF EXISTS pengurus_sosmed_check;

-- Pastikan kolom sosmed bisa menampung URL panjang
ALTER TABLE pengurus ALTER COLUMN sosmed TYPE VARCHAR(255);
