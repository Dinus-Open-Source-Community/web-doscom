CREATE TABLE IF NOT EXISTS activities (
    id SERIAL PRIMARY KEY,
    activities_title VARCHAR(255) NOT NULL,
    activities_desc TEXT,
    activities_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);