CREATE TABLE provinces (
   id SERIAL PRIMARY KEY,
   code TEXT NOT NULL,
   name TEXT NOT NULL,
   income_class TEXT,
   population BIGINT NOT NULL,
   region_code TEXT
);

CREATE UNIQUE INDEX provinces_code__idx ON provinces (code);
CREATE INDEX provinces_name__idx ON provinces (name);

ALTER TABLE provinces
ADD CONSTRAINT fk_provinces_region
FOREIGN KEY (region_code)
REFERENCES regions (code)