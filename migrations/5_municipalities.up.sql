CREATE TABLE municipalities (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    income_class VARCHAR(255),
    population BIGINT NOT NULL,
    province_id INTEGER,
    district_id INTEGER
);

CREATE UNIQUE INDEX municipalities_code__idx ON municipalities (code);
CREATE INDEX municipalities_name__idx ON municipalities (name);

ALTER TABLE municipalities
ADD CONSTRAINT fk_municipalities_provinces
FOREIGN KEY (province_id)
REFERENCES provinces (id);

ALTER TABLE municipalities
ADD CONSTRAINT fk_municipalities_districts
FOREIGN KEY (district_id)
REFERENCES districts (id);