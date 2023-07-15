CREATE TABLE sub_municipalities (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    population BIGINT NOT NULL,
    city_id INTEGER
);

CREATE UNIQUE INDEX sub_municipalities_code__idx ON sub_municipalities (code);
CREATE INDEX sub_municipalities_name__idx ON sub_municipalities (name);

ALTER TABLE sub_municipalities
ADD CONSTRAINT fk_sub_municipalities_cities
FOREIGN KEY (city_id)
REFERENCES cities (id)