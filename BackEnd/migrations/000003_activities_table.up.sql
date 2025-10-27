CREATE TABLE IF NOT EXISTS activities (
    id SERIAL PRIMARY KEY,
    id_asset INT REFERENCES gallery(id),
    activities_title VARCHAR(255) NOT NULL,
    activities_desc TEXT,
    activities_date TIMESTAMP,
    pengurus_id INT REFERENCES pengurus(id), -- to link activity to pengurus who organized it
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);