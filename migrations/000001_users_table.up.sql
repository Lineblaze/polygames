CREATE TABLE users
(
    id               serial PRIMARY KEY,
    name             text        NOT NULL,
    surname          text        NOT NULL,
    email            text UNIQUE NOT NULL,
    username         text UNIQUE NOT NULL,
    gender           text        NOT NULL,
    encoded_password text        NOT NULL,
    salt             text        NOT NULL,
    role             int2        NOT NULL,
    image_id         text        NOT NULL,
    date_of_birth    timestamptz NOT NULL,
    created_at       timestamptz NOT NULL DEFAULT now(),
    updated_at       timestamptz NOT NULL DEFAULT now(),
    disabled_at      timestamptz
);

