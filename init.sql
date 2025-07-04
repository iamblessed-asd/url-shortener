CREATE TABLE IF NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    short TEXT UNIQUE NOT NULL,
    original TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
