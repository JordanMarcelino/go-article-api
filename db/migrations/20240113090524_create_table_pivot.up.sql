CREATE TABLE articles_tags
(
    tag_id     BIGINT       NOT NULL,
    article_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (tag_id, article_id),
    FOREIGN KEY fk_articles_tags_tag_id (tag_id) REFERENCES tags (id),
    FOREIGN KEY fk_articles_tags_article_id (article_id) REFERENCES articles (id)
) ENGINE = InnoDB;