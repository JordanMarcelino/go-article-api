CREATE TABLE comments
(
    id         VARCHAR(255) NOT NULL,
    user_id    VARCHAR(255) NOT NULL,
    article_id VARCHAR(255) NOT NULL,
    body       VARCHAR(255) NOT NULL,
    created_at BIGINT       NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_comments_user_id FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_comments_article_id FOREIGN KEY (article_id) REFERENCES articles (id)
)