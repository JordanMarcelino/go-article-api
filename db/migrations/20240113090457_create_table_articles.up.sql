CREATE TABLE articles
(
    id          VARCHAR(255) NOT NULL,
    thumbnail   VARCHAR(100),
    title       VARCHAR(100),
    description VARCHAR(100),
    body        TEXT,
    created_at  DATETIME     NOT NULL,
    updated_at  DATETIME,
    user_id     VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_articles_user_id FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE = InnoDB;