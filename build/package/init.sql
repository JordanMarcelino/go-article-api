DELIMITER //

CREATE PROCEDURE GetArticleTags(IN p_article_id VARCHAR(255))
BEGIN
    SELECT tags.id, tags.name
    FROM tags
             JOIN articles_tags ON tags.id = articles_tags.tag_id
    WHERE articles_tags.article_id = p_article_id;
END //

DELIMITER ;


DELIMITER //

CREATE PROCEDURE GetArticleComments(IN p_article_id VARCHAR(255))
BEGIN
    SELECT comments.id, comments.body, comments.created_at
    FROM comments
             JOIN articles_comments ON comments.id = articles_comments.comment_id
    WHERE articles_comments.article_id = p_article_id;
END //

DELIMITER ;


DELIMITER //

CREATE PROCEDURE GetUserComments(IN p_user_id VARCHAR(255))
BEGIN
    SELECT comments.id, comments.body, comments.created_at
    FROM comments
             JOIN users_comments ON comments.id = users_comments.comment_id
    WHERE users_comments.user_id = p_user_id;
END //

DELIMITER ;