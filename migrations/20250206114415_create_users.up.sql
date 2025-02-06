CREATE TABLE users(
    id bigserial not null primary key,
    email varchar not null unique,
    ecnrypted_password varchar not null,
    weight integer not null
);