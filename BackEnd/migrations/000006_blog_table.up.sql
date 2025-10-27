CREATE TABLE blog (
    id SERIAL PRIMARY KEY,
    id_asset INT REFERENCES gallery(id),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    content TEXT,
    published_at TIMESTAMP,
    is_published BOOLEAN DEFAULT FALSE,
    work_id INT REFERENCES works(id),
    activity_id INT REFERENCES activities(id),
    pengurus_id INT REFERENCES pengurus(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);