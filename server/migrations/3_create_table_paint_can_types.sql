CREATE TABLE IF NOT EXISTS wall_items (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	uid UUID NOT NULL DEFAULT uuid_generate_v4(),
	name VARCHAR NOT NULL,
	liters NUMERIC (10,2),
)

INSERT INTO wall_items (name, liters) 
VALUES
    ('Tin 0.5L', 0.5),
    ('Tin 2.5L', 2.5);
	('Tin 3.6L', 3.6),
    ('Tin 18L', 18);
