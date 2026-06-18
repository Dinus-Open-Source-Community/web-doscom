#!/bin/sh
set -e

echo "Starting backend initialization..."

# Wait for database server to be ready (connect to 'postgres' db which always exists)
echo "Waiting for database server..."
until PGPASSWORD=$DB_PASSWORD psql -h "db" -U "$DB_USER" -d "postgres" -c '\q' 2>/dev/null; do
  echo "Database server is unavailable - sleeping"
  sleep 2
done

echo "Database server is ready!"

# Check if target database exists, create if not
echo "Checking target database: $DB_DATABASE..."
DB_EXISTS=$(PGPASSWORD=$DB_PASSWORD psql -h "db" -U "$DB_USER" -d "postgres" -tAc "SELECT 1 FROM pg_database WHERE datname='$DB_DATABASE'")
if [ "$DB_EXISTS" != "1" ]; then
    echo "Creating database $DB_DATABASE..."
    PGPASSWORD=$DB_PASSWORD psql -h "db" -U "$DB_USER" -d "postgres" -c "CREATE DATABASE $DB_DATABASE"
fi


# Run migrations
echo "Running database migrations..."
cd /app
./migrate down
./migrate up

echo "Running Seeder..."
./seeder
# Check if migrate binary exists, if not skip migration
# if command -v migrate >/dev/null 2>&1; then
#     migrate -path migrations -database "$DBURL" up || echo "⚠️  Migration failed or already applied"
# else
#     echo "⚠️  migrate tool not found, running migrations manually..."
#     for file in /app/migrations/*.up.sql; do
#         if [ -f "$file" ]; then
#             echo "Running $(basename $file)..."
#             PGPASSWORD=$DB_PASSWORD psql -h "db" -U "$DB_USER" -d "$DB_DATABASE" -f "$file" 2>&1 | grep -v "already exists" || true
#         fi
#     done
# fi

# Seed admin user if not exists
# echo " Seeding admin user..."
# PGPASSWORD=$DB_PASSWORD psql -h "db" -U "$DB_USER" -d "$DB_DATABASE" <<-'EOSQL'
#     -- Delete existing admin to ensure fresh start
#     DELETE FROM users WHERE email = 'admin@doscom.org' OR username = 'admin';
#
#     -- Insert fresh admin user with password 'admin'
#     INSERT INTO users (username, email, password, full_name, role, created_at, updated_at)
#     VALUES (
#         'admin',
#         'admin@doscom.org',
#         '$argon2id$v=19$m=65536,t=1,p=4$6it7jP6Yx3YshN2A2V3n8Q$Y9fF/D0+O+u6p9k7l0O2P3u5P6q7R8s9T0u1V2W3X4Y',
#         'Administrator',
#         'SuperAdmin',
#         CURRENT_TIMESTAMP,
#         CURRENT_TIMESTAMP
#     );
# EOSQL
# echo " Admin user re-created: email=admin@doscom.org, password=admin123"

echo "Initialization complete!"
echo "Starting backend server..."

# Start the application
exec ./app
# exec air -c .air.toml
