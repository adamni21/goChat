CREATE TABLE IF NOT EXISTS user (
    id PRIMARY KEY UNIQUE NOT NULL,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    isVerified INTEGER NOT NULL,
    passwordString TEXT NOT NULL
) STRICT