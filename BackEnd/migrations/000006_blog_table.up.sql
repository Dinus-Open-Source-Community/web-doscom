CREATE TABLE blog (
    id SERIAL PRIMARY KEY,
    id_asset INT REFERENCES gallery(id),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    content TEXT,
    kategori ENUM('event','seminar','collaboration','education','technology','work') NOT NULL,
    published_at TIMESTAMP,
    is_published BOOLEAN DEFAULT FALSE,
    id_work INT REFERENCES works(id),
    id_activity INT REFERENCES activities(id),
    id_pengurus INT REFERENCES pengurus(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);