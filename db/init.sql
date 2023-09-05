-- SET NAMES 'utf8';
-- SET time_zone = '+00:00';
-- SET foreign_key_checks = 0;
-- SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';


CREATE TABLE IF NOT EXISTS users (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(100),

    PRIMARY KEY (id),
    UNIQUE (name)
);


CREATE TABLE IF NOT EXISTS segments (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(100),

    PRIMARY KEY (id),
    UNIQUE (name)
);


CREATE TABLE IF NOT EXISTS user_segments (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    user_id BIGINT NOT NULL,
    segment_id BIGINT NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (segment_id) REFERENCES segments (id) ON DELETE CASCADE
);