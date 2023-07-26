CREATE TABLE barangays (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    urban_rural TEXT,
    population BIGINT NOT NULL,
    city_code TEXT,
    municipality_code TEXT,
    sub_municipality_code TEXT,
    special_government_unit_code TEXT
);

CREATE UNIQUE INDEX barangays_code__idx ON barangays (code);
CREATE INDEX barangays_name__idx ON barangays (name);

ALTER TABLE barangays
ADD CONSTRAINT fk_barangays_cities
FOREIGN KEY (city_code)
REFERENCES cities (code);

ALTER TABLE barangays
ADD CONSTRAINT fk_barangays_municipalities
FOREIGN KEY (municipality_code)
REFERENCES municipalities (code);

ALTER TABLE barangays
ADD CONSTRAINT fk_barangays_sub_municipalities
FOREIGN KEY (sub_municipality_code)
REFERENCES sub_municipalities (code);

ALTER TABLE barangays
ADD CONSTRAINT fk_barangays_special_government_units
FOREIGN KEY (special_government_unit_code)
REFERENCES special_government_units (code);