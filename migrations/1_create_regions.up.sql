CREATE TABLE regions (
     id SERIAL PRIMARY KEY,
     code VARCHAR(255) NOT NULL,
     name VARCHAR(255) NOT NULL,
     population BIGINT NOT NULL
);

CREATE UNIQUE INDEX regions_code__idx ON regions (code);
CREATE INDEX regions_name__idx ON regions (name);