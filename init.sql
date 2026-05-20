CREATE TABLE IF NOT EXISTS orders (
    id         SERIAL PRIMARY KEY,
    name       TEXT NOT NULL,
    contact    TEXT NOT NULL,
    message    TEXT,
    status     TEXT DEFAULT 'new',
    created_at TIMESTAMP DEFAULT NOW()
);