CREATE TABLE blog (
    id SERIAL PRIMARY KEY,
    id_asset INT REFERENCES gallery(id),
    id_work INT REFERENCES works(id),
    id_pengurus INT REFERENCES pengurus(id),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    content TEXT,
    kategori ENUM('event','seminar','collaboration','education','technology','work','activity') NOT NULL,
    published_at TIMESTAMP,
    is_published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);