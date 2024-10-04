-- admins table for storing admin credentials
CREATE TABLE admins (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    CONSTRAINT admins_pk PRIMARY KEY (id)
) ENGINE = InnoDB;

-- cars table for storing car details
CREATE TABLE cars (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    available BOOLEAN DEFAULT TRUE,
    CONSTRAINT cars_pk PRIMARY KEY (id)
) ENGINE = InnoDB;

-- bookings table for storing user bookings
CREATE TABLE bookings (
    id INT NOT NULL AUTO_INCREMENT,
    car_id INT NOT NULL,
    user_id INT NOT NULL,
    booking_date DATETIME NOT NULL,
    CONSTRAINT bookings_pk PRIMARY KEY (id),
    CONSTRAINT fk_bookings_car_id FOREIGN KEY (car_id)
        REFERENCES cars (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_bookings_user_id FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE
) ENGINE = InnoDB;

-- users table for storing user details (optional for this example)
CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    CONSTRAINT users_pk PRIMARY KEY (id)
) ENGINE = InnoDB;
