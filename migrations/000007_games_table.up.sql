CREATE TABLE IF NOT EXISTS games
(
    id          serial PRIMARY KEY,
    title       text NOT NULL,
    description text NOT NULL,
    user_id     int NOT NULL REFERENCES users (id),
    team_id     int NOT NULL REFERENCES teams (id),
    genre_id    int NOT NULL REFERENCES genres (id),
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now(),
    image_id    text NOT NULL,
    file_id     text NOT NULL,
    link        text
);


