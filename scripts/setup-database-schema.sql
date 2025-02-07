CREATE DATABASE IF NOT EXISTS exampledatabase;
USE exampledatabase;

CREATE TABLE IF NOT EXISTS products (
    id              INT                 PRIMARY KEY NOT NULL AUTO_INCREMENT,

    created_at      DATETIME            NOT NULL,
    updated_at      DATETIME            NOT NULL,
    deleted_at      DATETIME,

    code            VARCHAR(255)        NOT NULL,
    name            VARCHAR(255)        NOT NULL,
    description     VARCHAR(255)        NOT NULL,
    price           DOUBLE              NOT NULL
);
