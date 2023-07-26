CREATE TABLE municipalities (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    income_class TEXT,
    population BIGINT NOT NULL,
    province_code TEXT,
    district_code TEXT
);

CREATE UNIQUE INDEX municipalities_code__idx ON municipalities (code);
CREATE INDEX municipalities_name__idx ON municipalities (name);

ALTER TABLE municipalities
ADD CONSTRAINT fk_municipalities_provinces
FOREIGN KEY (province_code)
REFERENCES provinces (code);

ALTER TABLE municipalities
ADD CONSTRAINT fk_municipalities_districts
FOREIGN KEY (district_code)
REFERENCES districts (code);