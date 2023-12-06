CREATE TABLE IF NOT EXISTS users
(
    id            bigserial primary key,
    username      text not null,
    password      text not null,
    email         text,
    gender        int not null,
    date_of_birth text not null,
    "group"       text not null
);

CREATE TABLE IF NOT EXISTS games
(
    id            bigserial primary key,
    user_id       text not null,
    team_id       text not null,
    description   text,
    genre_id      int not null,
    creation_data date,
    last_update   date,
    likes         int,
    views         int,
    downloads     int,
    image         text not null,
    file          text
);

CREATE TABLE IF NOT EXISTS genres
(
    id         bigserial primary key,
    name       text not null
);