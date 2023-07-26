CREATE TABLE sub_municipalities (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    population BIGINT NOT NULL,
    city_code TEXT
);

CREATE UNIQUE INDEX sub_municipalities_code__idx ON sub_municipalities (code);
CREATE INDEX sub_municipalities_name__idx ON sub_municipalities (name);

ALTER TABLE sub_municipalities
ADD CONSTRAINT fk_sub_municipalities_cities
FOREIGN KEY (city_code)
REFERENCES cities (code)