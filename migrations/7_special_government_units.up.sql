CREATE TABLE special_government_units (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    province_id INTEGER
);

CREATE UNIQUE INDEX special_government_units_code__idx ON special_government_units (code);
CREATE INDEX special_government_units_name__idx ON special_government_units (name);

ALTER TABLE special_government_units
ADD CONSTRAINT fk_special_government_units_provinces
FOREIGN KEY (province_id)
REFERENCES provinces (id);