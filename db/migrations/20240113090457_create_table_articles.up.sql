CREATE TABLE articles
(
    id          VARCHAR(255) NOT NULL,
    user_id     VARCHAR(255) NOT NULL,
    thumbnail   VARCHAR(100),
    title       VARCHAR(100),
    description VARCHAR(100),
    body        TEXT,
    created_at  BIGINT       NOT NULL,
    updated_at  BIGINT,
    PRIMARY KEY (id),
    CONSTRAINT fk_articles_user_id FOREIGN KEY (user_id) REFERENCES users (id)
)