CREATE TABLE comments
(
    id         VARCHAR(255) NOT NULL,
    body       VARCHAR(255) NOT NULL,
    created_at DATETIME     NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;