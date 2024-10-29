CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE cars (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    car_number VARCHAR(255),
    capacity INT
);
