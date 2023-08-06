CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
id serial not null unique,
);