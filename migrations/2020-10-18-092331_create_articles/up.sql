-- Your SQL goes here
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE articles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    body TEXT NOT NULL
);