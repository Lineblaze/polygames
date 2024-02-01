CREATE TABLE IF NOT EXISTS games
(
    id          bigserial PRIMARY KEY,
    user_id     int NOT NULL REFERENCES users (id),
    team_id     int NOT NULL,
    description text,
    genre_id    int NOT NULL REFERENCES genres (id),
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now(),
    likes       int,
    views       int,
    downloads   int,
    image       text NOT NULL,
    file        text NOT NULL
);


