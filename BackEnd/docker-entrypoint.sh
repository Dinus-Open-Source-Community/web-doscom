#!/bin/sh
set -e

echo "Starting backend initialization..."

# Install postgresql-client for database operations
echo "Installing postgresql-client..."
apk add --no-cache postgresql-client

# Wait for database to be ready
echo "Waiting for database..."
until PGPASSWORD=$DB_PASSWORD psql -h "db" -U "$DB_USER" -d "$DB_DATABASE" -c '\q' 2>/dev/null; do
  echo "Database is unavailable - sleeping"
  sleep 2
done

echo "Database is ready!"

# Run migrations
echo "Running database migrations..."
cd /app

# Check if migrate binary exists, if not skip migration
if command -v migrate >/dev/null 2>&1; then
    migrate -path migrations -database "$DBURL" up || echo "⚠️  Migration failed or already applied"
else
    echo "⚠️  migrate tool not found, running migrations manually..."
    for file in /app/migrations/*.up.sql; do
        if [ -f "$file" ]; then
            echo "Running $(basename $file)..."
            PGPASSWORD=$DB_PASSWORD psql -h "db" -U "$DB_USER" -d "$DB_DATABASE" -f "$file" 2>&1 | grep -v "already exists" || true
        fi
    done
fi

# Seed admin user if not exists
echo "🌱 Seeding admin user..."
PGPASSWORD=$DB_PASSWORD psql -h "db" -U "$DB_USER" -d "$DB_DATABASE" << 'EOF'
DO $$
BEGIN
    -- Check if admin user exists
    IF NOT EXISTS (SELECT 1 FROM users WHERE username = 'admin') THEN
        -- Insert admin user with password 'admin123' (hashed with bcrypt)
        INSERT INTO users (username, email, password, full_name, role, created_at, updated_at)
        VALUES (
            'admin',
            'admin@doscom.org',
            '$2a$10$4RmSchKkU25uPrzCgxHtbuFvTLDLB/lRr.JM9DwiMB27IStaoAl2K',
            'Administrator',
            'Super_Admin',
            CURRENT_TIMESTAMP,
            CURRENT_TIMESTAMP
        );
        RAISE NOTICE '✅ Admin user created: username=admin, password=admin123';
    ELSE
        RAISE NOTICE '⚠️  Admin user already exists';
    END IF;
END $$;
EOF

echo "✅ Initialization complete!"
echo "🚀 Starting backend server..."

# Start the application
exec go run ./cmd/api/main.go
