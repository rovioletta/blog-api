-- name: DeleteArticleByID :exec
DELETE FROM article WHERE id = @id::integer;