CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(100),
    price INT NOT NULL
);