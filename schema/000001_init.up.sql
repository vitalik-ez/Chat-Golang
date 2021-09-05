CREATE TABLE IF NOT EXISTS users (
    id serial not null UNIQUE,
    name varchar(255) not null,
    email varchar(255) not null UNIQUE,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS rooms (
    id serial not null UNIQUE,
    name varchar(255) not null,
    founder_id int REFERENCES users(id) on DELETE CASCADE not NULL
);

CREATE TABLE IF NOT EXISTS users_rooms (
    id serial not null UNIQUE,
    user_id int REFERENCES users(id) on DELETE CASCADE not NULL,
    room_id int REFERENCES rooms(id) on DELETE CASCADE not NULL
);

CREATE TABLE IF NOT EXISTS messages (
    id serial not null UNIQUE,
    room_id int REFERENCES rooms(id) on DELETE CASCADE not NULL,
    user_id int REFERENCES users(id) on DELETE CASCADE not NULL,
    timestamp timestamp default current_timestamp,
    text_message varchar(255) not null
);