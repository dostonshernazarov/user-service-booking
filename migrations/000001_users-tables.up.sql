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
('1a8d3b24-3c29-4d21-a7d6-fcdd5a92c56a', 'Bekmurodov Avazbek', 'kimsansaney.1@gmail.com', 'hashed_password', '1990-05-15', 'https://example.com/profile_image.jpg', '1234 5678 9012 3456', 'Male', '+1234567890', 'admin', 'some_refresh_token', '2024-05-08 12:30:00', '2024-05-08 12:30:00', NULL),
('2e9ca276-5799-4f34-a0a6-938f7b0a5c8d', 'Jane Smith', 'jane@example.com', 'hashed_password', NULL, NULL, NULL, 'Female', NULL, 'user', NULL, '2024-05-08 13:45:00', NULL, NULL),
('3f1b673a-81f7-4f2e-891a-fcdd5a92c56a', 'Alice Johnson', 'alice@example.com', 'hashed_password', '1985-08-20', 'https://example.com/alice_profile.jpg', '9876 5432 1098 7654', 'Female', '+1987654321', 'admin', NULL, '2024-05-08 14:15:00', NULL, NULL),
('4d2a78b3-9c43-4e76-ae2e-938f7b0a5c8d', 'Bob Williams', 'bob@example.com', 'hashed_password', '1978-03-10', 'https://example.com/bob_profile.jpg', '8765 4321 0987 6543', 'Male', '+1765432987', 'user', NULL, '2024-05-08 14:45:00', NULL, NULL),

INSERT INTO users_establishment (user_id, establishment_id) VALUES
('1a8d3b24-3c29-4d21-a7d6-fcdd5a92c56a', 'e1a23456-7890-1234-5678-901234567890'),
('2e9ca276-5799-4f34-a0a6-938f7b0a5c8d', 'e2b34567-8901-2345-6789-012345678901'),
('3f1b673a-81f7-4f2e-891a-fcdd5a92c56a', 'e3c45678-9012-3456-7890-123456789012'),
('4d2a78b3-9c43-4e76-ae2e-938f7b0a5c8d', 'e4d56789-0123-4567-8901-234567890123');
