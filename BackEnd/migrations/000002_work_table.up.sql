CREATE TABLE IF NOT EXISTS works (
    id SERIAL PRIMARY KEY,
    id_asset INT REFERENCES gallery(id),
    title VARCHAR(300) NOT NULL,
    description TEXT,
    project_date TEXT,
    pengurus_id INT REFERENCES pengurus(id), -- to link work to pengurus who did it for team recognition
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- join work activity blog user
-- image_url