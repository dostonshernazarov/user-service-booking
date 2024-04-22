CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    date_of_birth DATE,
    profile_img VARCHAR(255),
    card VARCHAR(255),
    gender VARCHAR(255),
    phone_number VARCHAR(255),
    role VARCHAR(255),
    establishment_id UUID,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);