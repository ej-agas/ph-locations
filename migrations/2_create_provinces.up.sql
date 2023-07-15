CREATE TABLE provinces (
   id SERIAL PRIMARY KEY,
   code VARCHAR(255) NOT NULL,
   name VARCHAR(255) NOT NULL,
   income_class VARCHAR(255),
   population BIGINT NOT NULL,
   region_id INTEGER
);

CREATE UNIQUE INDEX provinces_code__idx ON provinces (code);
CREATE INDEX provinces_name__idx ON provinces (name);

ALTER TABLE provinces
ADD CONSTRAINT fk_provinces_region
FOREIGN KEY (region_id)
REFERENCES regions (id)