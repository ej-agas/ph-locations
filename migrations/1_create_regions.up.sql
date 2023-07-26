CREATE TABLE regions (
     id SERIAL PRIMARY KEY,
     code TEXT NOT NULL,
     name TEXT NOT NULL,
     population BIGINT NOT NULL
);

CREATE UNIQUE INDEX regions_code__idx ON regions (code);
CREATE INDEX regions_name__idx ON regions (name);