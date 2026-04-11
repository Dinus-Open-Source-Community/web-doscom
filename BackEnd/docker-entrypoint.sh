#!/bin/sh
set -e

echo "Starting backend initialization..."

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
PGPASSWORD=$DB_PASSWORD psql -h "db" -U "$DB_USER" -d "$DB_DATABASE" <<-EOSQL
    -- Delete existing admin to ensure fresh start
    DELETE FROM users WHERE email = 'admin@doscom.org' OR username = 'admin';
    
    -- Insert fresh admin user with password 'admin123'
    INSERT INTO users (username, email, password, full_name, role, created_at, updated_at)
    VALUES (
        'admin',
        'admin@doscom.org',
        '\$argon2id\$v=19\$m=65536,t=3,p=2\$TUV6VEpxR3VpYllRcUJpRA\$7H+D7Cg0zKEnshgU0fQ7u3f9mZ4Kj9C6vS8aF9D9m6k',
        'Administrator',
        'SuperAdmin',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );
EOSQL
echo "✅ Admin user re-created: email=admin@doscom.org, password=admin123"

echo "✅ Initialization complete!"
echo "🚀 Starting backend server..."

# Start the application
exec ./app
