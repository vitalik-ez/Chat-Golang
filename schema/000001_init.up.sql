CREATE TABLE IF NOT EXISTS rooms (
    id serial not null UNIQUE,
    name varchar(255) not null UNIQUE
);

CREATE TABLE IF NOT EXISTS messages (
    id serial not null UNIQUE,
    room_id int REFERENCES rooms(id) on DELETE CASCADE not NULL,
    username varchar(255) not null,
    timestamp timestamp default current_timestamp,
    text_message varchar(255) not null
);