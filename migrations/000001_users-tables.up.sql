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
    refresh_token VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS userHotelBooking (
    id UUID NOT NULL,
    user_id UUID NOT NULL,
    room_id UUID NOT NULL,
    will_arrive TIMESTAMP,
    will_leave TIMESTAMP,
    number_of_people INTEGER,
    is_canceled BOOLEAN DEFAULT FALSE,
    reason VARCHAR(255) IF is_canceled IS TRUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (room_id) REFERENCES establishments(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS userRestaurantBooking (
    id UUID NOT NULL,
    user_id UUID NOT NULL,
    table_id UUID NOT NULL,
    will_arrive TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (table_id) REFERENCES establishments(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);