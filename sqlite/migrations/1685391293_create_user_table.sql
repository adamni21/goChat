CREATE TABLE IF NOT EXISTS users (
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    isVerified INTEGER NOT NULL,
    passwordString TEXT NOT NULL
) STRICT
