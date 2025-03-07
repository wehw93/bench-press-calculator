CREATE TABLE users(
    id bigserial not null primary key,
    email varchar not null,
    encrypted_password varchar not null,
    weight float4
);