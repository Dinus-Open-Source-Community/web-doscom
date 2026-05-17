package seeder

import (
	"gorm.io/gorm"
	"log"
)

type SeederFunc func(tx *gorm.DB)

type Seeder struct {
	Name string
	Fn   SeederFunc
}

func RunAllSeeders(db *gorm.DB) {
	log.Println("Starting database seeding...")

	seeders := []Seeder{
		{"users", SeedUsers},
		{"pengurus", RunSeedPengurus},
		{"work", RunSeedWorks},
		{"blog", RunSeedBlogs},
	}

	for _, s := range seeders {
		runSeeder(db, s.Name, s.Fn)
	}

	log.Println("Database seeding completed.")
}

func runSeeder(db *gorm.DB, name string, fn SeederFunc) {
	log.Printf("Seeding %s...", name)

	tx := db.Begin()
	if tx.Error != nil {
		log.Fatalf("failed to start transaction: %v", tx.Error)
	}

	fn(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Fatalf("seeder %s failed: %v", name, err)
	}

	log.Printf("Seeding %s done", name)
}
