CREATE TABLE IF NOT EXISTS "user"
(
    id          INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        TEXT CONSTRAINT user_name_length CHECK (LENGTH(name) BETWEEN 2 AND 20) NOT NULL,
    surname     TEXT CONSTRAINT user_surname_length CHECK (LENGTH(surname) BETWEEN 2 AND 30) NOT NULL,
    email       TEXT CONSTRAINT user_email_domain CHECK (LENGTH(email) BETWEEN 6 AND 50) UNIQUE NOT NULL UNIQUE,
    password    BYTEA CONSTRAINT user_password_length CHECK (LENGTH(password) = 60) NOT NULL,
    birthday    DATE NULL
);

CREATE TABLE IF NOT EXISTS birthday_subscribing
(
    birthday_user_id        INTEGER CONSTRAINT foreign_key_bday_user CHECK (birthday_user_id > 0) REFERENCES "user" (id) ON DELETE CASCADE,
    subscribing_user_id     INTEGER CONSTRAINT foreign_key_sub_user CHECK (subscribing_user_id > 0) REFERENCES "user" (id) ON DELETE CASCADE,
    PRIMARY KEY (birthday_user_id, subscribing_user_id)
);