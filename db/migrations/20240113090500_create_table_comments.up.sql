CREATE TABLE comments
(
    id         VARCHAR(255) NOT NULL,
    user_id    VARCHAR(255) NOT NULL,
    article_id VARCHAR(255) NOT NULL,
    body       VARCHAR(255) NOT NULL,
    created_at BIGINT       NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY fk_comments_user_id (user_id) REFERENCES users(id),
    FOREIGN KEY fk_comments_article_id (article_id) REFERENCES articles(id)
) ENGINE = InnoDB;