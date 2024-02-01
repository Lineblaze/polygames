CREATE TABLE IF NOT EXISTS games_platforms
(
    game_id     int8 NOT NULL REFERENCES games (id),
    platform_id int4 NOT NULL REFERENCES platforms (id)
);
