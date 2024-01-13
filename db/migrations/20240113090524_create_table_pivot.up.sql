CREATE TABLE articles_tags
(
    tag_id     INT          NOT NULL,
    article_id VARCHAR(255) NOT NULL,
    FOREIGN KEY fk_articles_tags_tag_id (tag_id) REFERENCES tags (id),
    FOREIGN KEY fk_articles_tags_article_id (article_id) REFERENCES articles (id)
) ENGINE = InnoDB;

CREATE TABLE articles_comments
(
    comment_id VARCHAR(255) NOT NULL,
    article_id VARCHAR(255) NOT NULL,
    FOREIGN KEY fk_articles_comments_comment_id (comment_id) REFERENCES comments (id),
    FOREIGN KEY fk_articles_comments_article_id (article_id) REFERENCES articles (id)
) ENGINE = InnoDB;

CREATE TABLE users_comments
(
    user_id    VARCHAR(255) NOT NULL,
    comment_id VARCHAR(255) NOT NULL,
    FOREIGN KEY fk_users_comments_user_id (user_id) REFERENCES users (id),
    CONSTRAINT fk_users_comments_comment_id FOREIGN KEY (comment_id) REFERENCES comments (id)
) ENGINE = InnoDB;
