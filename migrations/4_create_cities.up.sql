CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    city_class TEXT,
    income_class TEXT,
    population BIGINT NOT NULL,
    province_code TEXT,
    district_code TEXT
);

CREATE UNIQUE INDEX cities_code__idx ON cities (code);
CREATE INDEX cities_name__idx ON cities (name);

ALTER TABLE cities
ADD CONSTRAINT fk_cities_provinces
FOREIGN KEY (province_code)
REFERENCES provinces (code);

ALTER TABLE cities
ADD CONSTRAINT fk_cities_districts
FOREIGN KEY (district_code)
REFERENCES districts (code)