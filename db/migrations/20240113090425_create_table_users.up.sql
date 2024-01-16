CREATE TABLE users
(
    id         VARCHAR(255) NOT NULL,
    username   VARCHAR(100),
    password   VARCHAR(255) NOT NULL,
    email      VARCHAR(100) NOT NULL UNIQUE,
    phone      VARCHAR(100),
    avatar     VARCHAR(100),
    created_at BIGINT     NOT NULL,
    updated_at BIGINT,
    PRIMARY KEY (id)
) ENGINE = InnoDB;