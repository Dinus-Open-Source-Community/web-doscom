package seeder

import (
	"log"

	"gorm.io/gorm"
)

// RunAllSeeders mengeksekusi semua seeder yang ada secara berurutan.
func RunAllSeeders(db *gorm.DB) {
	log.Println("Starting database seeding...")

	SeedUsers(db)
	SeedPengurus(db)
	SeedGallery(db)
	SeedFileUploads(db)
	SeedWorks(db)
	SeedBlogs(db)

	log.Println("Database seeding completed.")
}
