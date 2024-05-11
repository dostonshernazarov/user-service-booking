CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL UNIQUE,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    date_of_birth DATE,
    profile_img VARCHAR(255),
    card VARCHAR(255),
    gender VARCHAR(255),
    phone_number VARCHAR(255),
    role VARCHAR(255),
    refresh_token VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users_establishment (
    user_id UUID NOT NULL,
    establishment_id UUID NOT NULL UNIQUE,
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

INSERT INTO users (id, full_name, email, password, date_of_birth, profile_img, card, gender, phone_number, role, refresh_token, created_at, updated_at, deleted_at) VALUES
('1a8d3b24-3c29-4d21-a7d6-fcdd5a92c56a', 'Bekmurodov Avazbek', 'kimsansaney.1@gmail.com', 'hashed_password', '2005-02-06', 'https://example.com/profile_image.jpg', '1234 5678 9012 3456', 'Male', '+998337709978', 'admin', 'some_refresh_token', '2024-05-08 12:30:00', '2024-05-08 12:30:00', NULL),
('2e9ca276-5799-4f34-a0a6-938f7b0a5c8d', 'Jane Smith', 'jane@example.com', 'hashed_password', '1990-05-15', 'https://example.com/profile_image.jpg', 'NULL', 'Female', 'NULL', '+1234567890', 'NULL', '2024-05-08 13:45:00', '2024-05-08 13:45:00', NULL),
('3f1b673a-81f7-4f2e-891a-fcdd5a92c56a', 'Alice Johnson', 'alice@example.com', 'hashed_password', '1985-08-20', 'https://example.com/alice_profile.jpg', '9876 5432 1098 7654', 'Female', '+1987654321', 'admin', 'NULL', '2024-05-08 14:15:00', '2024-05-08 13:45:00', NULL),
('4d2a78b3-9c43-4e76-ae2e-938f7b0a5c8d', 'Bob Williams', 'bob@example.com', 'hashed_password', '1978-03-10', 'https://example.com/bob_profile.jpg', '8765 4321 0987 6543', 'Male', '+1765432987', 'user', 'NULL', '2024-05-08 14:45:00', '2024-05-08 13:45:00', NULL);

INSERT INTO users_establishment (user_id, establishment_id) VALUES
('1a8d3b24-3c29-4d21-a7d6-fcdd5a92c56a', 'e1a23456-7890-1234-5678-901234567890'),
('2e9ca276-5799-4f34-a0a6-938f7b0a5c8d', 'e2b34567-8901-2345-6789-012345678901'),
('3f1b673a-81f7-4f2e-891a-fcdd5a92c56a', 'e3c45678-9012-3456-7890-123456789012'),
('4d2a78b3-9c43-4e76-ae2e-938f7b0a5c8d', 'e4d56789-0123-4567-8901-234567890123');


CREATE TABLE "location_table"(
    "location_id" UUID PRIMARY KEY NOT NULL,
    "establishment_id" UUID NOT NULL,
    "address" VARCHAR(255) DEFAULT '',
    "latitude" FLOAT DEFAULT 0,
    "longitude" FLOAT DEFAULT 0,
    "country" VARCHAR(255) DEFAULT '',
    "city" VARCHAR(255) DEFAULT '',
    "state_province" VARCHAR(255) DEFAULT '',
    "created_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(0)
);

CREATE TABLE "room_table"(
    "room_id" UUID PRIMARY KEY NOT NULL,
    "hotel_id" UUID NOT NULL,
    "price" FLOAT DEFAULT 0,
    "description" TEXT DEFAULT '',
    "number_of_rooms" BIGINT DEFAULT 0,
    "holidays" VARCHAR(255) DEFAULT '',
    "free_days" VARCHAR(255) DEFAULT '',
    "discount" FLOAT DEFAULT 0,
    "created_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(0)
);

CREATE TABLE "favourite_table"(
    "favourite_id" UUID PRIMARY KEY NOT NULL,
    "establishment_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "created_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(0)
);

CREATE TABLE "review_table"(
    "review_id" UUID PRIMARY KEY NOT NULL,
    "establishment_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "rating" FLOAT DEFAULT 0,
    "comment" TEXT DEFAULT '',
    "created_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(0)
);

CREATE TABLE "image_table"(
    "image_id" UUID PRIMARY KEY NOT NULL,
    "establishment_id" UUID NOT NULL,
    "image_url" VARCHAR(255) DEFAULT '',
    "created_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(0)
);

CREATE TABLE "restaurant_table"(
    "restaurant_id" UUID PRIMARY KEY NOT NULL,
    "owner_id" UUID NOT NULL,
    "restaurant_name" VARCHAR(255) DEFAULT '',
    "description" TEXT DEFAULT '',
    "rating" FLOAT DEFAULT 0,
    "opening_hours" VARCHAR(255) DEFAULT '',
    "contact_number" VARCHAR(255) DEFAULT '',
    "licence_url" VARCHAR(255) DEFAULT '',
    "website_url" VARCHAR(255) DEFAULT '',
    "created_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(0)
);

CREATE TABLE "attraction_table"(
    "attraction_id" UUID PRIMARY KEY NOT NULL,
    "owner_id" UUID NOT NULL,
    "attraction_name" VARCHAR(255) DEFAULT '',
    "description" VARCHAR(255) DEFAULT '',
    "rating" FLOAT DEFAULT 0,
    "contact_number" VARCHAR(255) DEFAULT '',
    "licence_url" VARCHAR(255) DEFAULT '',
    "website_url" VARCHAR(255) DEFAULT '',
    "created_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(0)
);

CREATE TABLE "hotel_table"(
    "hotel_id" UUID PRIMARY KEY NOT NULL,
    "owner_id" UUID NOT NULL,
    "hotel_name" VARCHAR(255) DEFAULT '',
    "description" TEXT DEFAULT '',
    "rating" FLOAT DEFAULT 0,
    "contact_number" VARCHAR(255) DEFAULT '',
    "licence_url" VARCHAR(255) DEFAULT '',
    "website_url" VARCHAR(255) DEFAULT '',
    "created_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(0)
);