-- name: CreateArticle :exec
INSERT INTO article (title, content, tags)
VALUES ($1, $2, $3);