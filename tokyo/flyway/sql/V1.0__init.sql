CREATE TABLE users
(
    id              BIGINT PRIMARY KEY,
    uuid            UUID UNIQUE    NOT NULL,
    email           VARCHAR UNIQUE NOT NULL,
    hashed_password VARCHAR        NOT NULL
);