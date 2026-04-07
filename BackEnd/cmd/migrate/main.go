// cmd/migration/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	env "web_doscom/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Load environment variables
	env.LoadEnv()

	// Define flags
	var (
		action  string
		steps   int
		version int
	)

	flag.StringVar(&action, "action", "", "Migration action: up, down, force, version, drop, steps")
	flag.IntVar(&steps, "steps", 0, "Number of steps for migration (use with up/down)")
	flag.IntVar(&version, "version", 0, "Force migration to specific version")
	flag.Parse()

	// Fallback to old behavior for backward compatibility
	if action == "" && len(os.Args) > 1 {
		action = os.Args[1]
	}

	if action == "" {
		printUsage()
		os.Exit(1)
	}

	// Get database URL
	dbURL := os.Getenv("DBURL")
	if dbURL == "" {
		log.Fatal("❌ DBURL environment variable is not set")
	}

	// Create migration instance
	m, err := migrate.New("file://./migrations", dbURL)
	if err != nil {
		log.Fatal("❌ Failed to create migration instance:", err)
	}
	defer m.Close()

	// Execute migration action
	switch action {
	case "up":
		if steps > 0 {
			if err := m.Steps(steps); err != nil && err != migrate.ErrNoChange {
				log.Fatal("❌ Migration failed:", err)
			}
			fmt.Printf("✅ Migrated up %d step(s)\n", steps)
		} else {
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatal("❌ Migration failed:", err)
			}
			fmt.Println("✅ All migrations applied successfully!")
		}

	case "down":
		if steps > 0 {
			if err := m.Steps(-steps); err != nil && err != migrate.ErrNoChange {
				log.Fatal("❌ Migration failed:", err)
			}
			fmt.Printf("✅ Rolled back %d step(s)\n", steps)
		} else {
			if err := m.Down(); err != nil && err != migrate.ErrNoChange {
				log.Fatal("❌ Migration failed:", err)
			}
			fmt.Println("✅ All migrations rolled back!")
		}

	case "drop":
		fmt.Println("⚠️  WARNING: This will drop all tables!")
		fmt.Print("Are you sure? (yes/no): ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm == "yes" {
			if err := m.Drop(); err != nil {
				log.Fatal("❌ Drop failed:", err)
			}
			fmt.Println("✅ All tables dropped!")
		} else {
			fmt.Println("❌ Operation cancelled")
		}

	case "force":
		if version == 0 {
			log.Fatal("❌ Please provide version with -version flag")
		}
		if err := m.Force(version); err != nil {
			log.Fatal("❌ Force failed:", err)
		}
		fmt.Printf("✅ Forced migration to version %d\n", version)

	case "version":
		v, dirty, err := m.Version()
		if err != nil {
			log.Fatal("❌ Failed to get version:", err)
		}
		fmt.Printf("📊 Current version: %d\n", v)
		fmt.Printf("📊 Dirty state: %t\n", dirty)

	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run cmd/migration/main.go -action=<action> [options]")
	fmt.Println()
	fmt.Println("Actions:")
	fmt.Println("  up              Apply all pending migrations")
	fmt.Println("  down            Rollback all migrations")
	fmt.Println("  drop            Drop all tables (requires confirmation)")
	fmt.Println("  force           Force migration to specific version")
	fmt.Println("  version         Show current migration version")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -steps=N        Run N migration steps (use with up/down)")
	fmt.Println("  -version=N      Force to specific version (use with force)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run cmd/migration/main.go -action=up")
	fmt.Println("  go run cmd/migration/main.go -action=down -steps=1")
	fmt.Println("  go run cmd/migration/main.go -action=force -version=1")
	fmt.Println("  go run cmd/migration/main.go -action=version")
	fmt.Println()
	fmt.Println("Backward compatible (old style):")
	fmt.Println("  go run cmd/migration/main.go up")
	fmt.Println("  go run cmd/migration/main.go down")
}
