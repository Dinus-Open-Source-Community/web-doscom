CREATE TABLE blog_gallery (
	id SERIAL PRIMARY KEY,
	id_blog INT REFERENCES blog(id),
	id_gallery INT REFERENCES gallery(id)
);
