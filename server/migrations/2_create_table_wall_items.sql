CREATE TABLE IF NOT EXISTS wall_items (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	uid UUID NOT NULL DEFAULT uuid_generate_v4(),
	name VARCHAR NOT NULL,
	width NUMERIC (10,2),
	height NUMERIC (10,2),
	square_meters NUMERIC (10,2)
)

INSERT INTO wall_items (name, width, height, square_meters) 
VALUES
    ('Door', 0.80, 1.90, 2.70),
    ('Window', 2, 1.20, 2.40);