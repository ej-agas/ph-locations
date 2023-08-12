/*
|--------------------------------------------------------------------------
| Regions Table
|--------------------------------------------------------------------------
*/
CREATE TABLE regions (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    population BIGINT NOT NULL
);

CREATE UNIQUE INDEX regions_code__idx ON regions (code);
CREATE INDEX regions_name__idx ON regions (name);


/*
|--------------------------------------------------------------------------
| Provinces Table
|--------------------------------------------------------------------------
*/
CREATE TABLE provinces (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    income_class TEXT,
    population BIGINT NOT NULL,
    region_code TEXT references regions(code)
);

CREATE UNIQUE INDEX provinces_code__idx ON provinces (code);
CREATE INDEX provinces_name__idx ON provinces (name);
CREATE INDEX provinces_region_code__idx on provinces(region_code);

/*
|--------------------------------------------------------------------------
| Districts Table
|--------------------------------------------------------------------------
*/
CREATE TABLE districts (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    population BIGINT NOT NULL,
    region_code TEXT references regions(code)
);

CREATE UNIQUE INDEX districts_code__idx ON districts (code);
CREATE INDEX districts_name__idx ON districts (name);
CREATE INDEX districts_region_code__idx ON districts (region_code);


/*
|--------------------------------------------------------------------------
| Cities Table
|--------------------------------------------------------------------------
*/
CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    city_class TEXT,
    income_class TEXT,
    population BIGINT NOT NULL,
    province_code TEXT references provinces(code),
    district_code TEXT references districts(code)
);

CREATE UNIQUE INDEX cities_code__idx ON cities (code);
CREATE INDEX cities_name__idx ON cities (name);
CREATE INDEX cities_province_code__idx ON cities (province_code);
CREATE INDEX cities_district_code__idx ON cities (district_code);


/*
|--------------------------------------------------------------------------
| Municipalities Table
|--------------------------------------------------------------------------
*/
CREATE TABLE municipalities (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    income_class TEXT,
    population BIGINT NOT NULL,
    province_code TEXT references provinces(code),
    district_code TEXT references districts(code)
);

CREATE UNIQUE INDEX municipalities_code__idx ON municipalities (code);
CREATE INDEX municipalities_name__idx ON municipalities (name);
CREATE INDEX municipalities_province_code__idx ON municipalities (province_code);
CREATE INDEX municipalities_district_code__idx ON municipalities (district_code);


/*
|--------------------------------------------------------------------------
| Sub-Municipalities Table
|--------------------------------------------------------------------------
*/
CREATE TABLE sub_municipalities (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    population BIGINT NOT NULL,
    city_code TEXT references cities(code)
);

CREATE UNIQUE INDEX sub_municipalities_code__idx ON sub_municipalities (code);
CREATE INDEX sub_municipalities_name__idx ON sub_municipalities (name);
CREATE INDEX sub_municipalities_city_code__idx ON sub_municipalities (city_code);


/*
|--------------------------------------------------------------------------
| Special Government Units Table
|--------------------------------------------------------------------------
*/
CREATE TABLE special_government_units (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    province_code TEXT references provinces(code)
);

CREATE UNIQUE INDEX special_government_units_code__idx ON special_government_units (code);
CREATE INDEX special_government_units_name__idx ON special_government_units (name);
CREATE INDEX special_government_units_province_code__idx ON special_government_units (province_code);


/*
|--------------------------------------------------------------------------
| Barangays Table
|--------------------------------------------------------------------------
*/
CREATE TABLE barangays (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    urban_rural TEXT,
    population BIGINT NOT NULL,
    city_code TEXT references cities(code),
    municipality_code TEXT references municipalities(code),
    sub_municipality_code TEXT references sub_municipalities(code),
    special_government_unit_code TEXT references special_government_units(code)
);

CREATE UNIQUE INDEX barangays_code__idx ON barangays (code);
CREATE INDEX barangays_name__idx ON barangays (name);
CREATE INDEX barangays_city_code__idx ON barangays (city_code);
CREATE INDEX barangays_municipality_code__idx ON barangays (municipality_code);
CREATE INDEX barangays_sub_municipality_code__idx ON barangays (sub_municipality_code);
CREATE INDEX barangays_special_government_unit_code__idx ON barangays (special_government_unit_code);

