CREATE TABLE districts (
   id SERIAL PRIMARY KEY,
   code VARCHAR(255) NOT NULL,
   name VARCHAR(255) NOT NULL,
   population BIGINT NOT NULL,
   region_id INTEGER
);

CREATE UNIQUE INDEX districts_code__idx ON districts (code);
CREATE INDEX districts_name__idx ON districts (name);

ALTER TABLE districts
ADD CONSTRAINT fk_districts_region
FOREIGN KEY (region_id)
REFERENCES regions (id)