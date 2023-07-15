CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    city_class VARCHAR(255),
    income_class VARCHAR(255),
    population BIGINT NOT NULL,
    province_id INTEGER,
    district_id INTEGER
);

CREATE UNIQUE INDEX cities_code__idx ON cities (code);
CREATE INDEX cities_name__idx ON cities (name);

ALTER TABLE cities
ADD CONSTRAINT fk_cities_provinces
FOREIGN KEY (province_id)
REFERENCES provinces (id);

ALTER TABLE cities
ADD CONSTRAINT fk_cities_districts
FOREIGN KEY (district_id)
REFERENCES districts (id)