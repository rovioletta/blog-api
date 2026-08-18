-- name: GetArticleByID :one
SELECT
  *
FROM
  article
WHERE
  id = @id :: integer
LIMIT
  1;