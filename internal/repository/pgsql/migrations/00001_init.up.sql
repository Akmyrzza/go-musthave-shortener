CREATE TABLE urls(
    id SERIAL PRIMARY KEY,
    originalURL VARCHAR(255) UNIQUE NOT NULL,
    shortURL VARCHAR(255) UNIQUE NOT NULL,
    userID VARCHAR(255),
    isDeleted BOOLEAN DEFAULT FALSE
);
