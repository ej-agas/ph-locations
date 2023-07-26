CREATE TABLE districts (
   id SERIAL PRIMARY KEY,
   code TEXT NOT NULL,
   name TEXT NOT NULL,
   population BIGINT NOT NULL,
   region_code TEXT
);

CREATE UNIQUE INDEX districts_code__idx ON districts (code);
CREATE INDEX districts_name__idx ON districts (name);

ALTER TABLE districts
ADD CONSTRAINT fk_districts_region
FOREIGN KEY (region_code)
REFERENCES regions (code)