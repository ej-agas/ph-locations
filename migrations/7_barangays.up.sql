CREATE TABLE barangays (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    urban_rural VARCHAR(255),
    population BIGINT NOT NULL,
    city_id INTEGER,
    municipality_id INTEGER,
    sub_municipality_id INTEGER
);

CREATE UNIQUE INDEX barangays_code__idx ON barangays (code);
CREATE INDEX barangays_name__idx ON barangays (name);

ALTER TABLE barangays
ADD CONSTRAINT fk_barangays_cities
FOREIGN KEY (city_id)
REFERENCES cities (id);

ALTER TABLE barangays
ADD CONSTRAINT fk_barangays_municipalities
FOREIGN KEY (municipality_id)
REFERENCES municipalities (id);

ALTER TABLE barangays
ADD CONSTRAINT fk_barangays_sub_municipalities
FOREIGN KEY (sub_municipality_id)
REFERENCES sub_municipalities (id);