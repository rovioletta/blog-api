-- name: UpdateArticle :exec
UPDATE
  article
SET
  title = COALESCE(sqlc.narg('new_title'), title),
  content = COALESCE(sqlc.narg('new_content'), content),
  tags = COALESCE(sqlc.narg('new_tags') :: text [], tags)
WHERE
  id = sqlc.arg('id') :: integer;