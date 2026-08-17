-- name: GetArticlesByFilter :many
SELECT
  *
FROM
  article
WHERE
  -- by title
  (sqlc.narg('filter_search_title')::text IS NULL OR title LIKE CONCAT('%', sqlc.narg('filter_search_title')::text, '%'))
  
  -- by tags
  AND (
    @filter_search_tags::text[] IS NULL 
    OR cardinality(@filter_search_tags) = 0 
    OR tags && @filter_search_tags
  )
  
  -- by created_at
  AND (sqlc.narg('filter_created_from')::timestamp IS NULL OR created_at >= @filter_created_from)
  AND (sqlc.narg('filter_created_to')::timestamp IS NULL OR created_at <= @filter_created_to)
  
  -- by updated_at
  AND (sqlc.narg('filter_updated_from')::timestamp IS NULL OR updated_at >= @filter_updated_from)
  AND (sqlc.narg('filter_updated_to')::timestamp IS NULL OR updated_at <= @filter_updated_to)
ORDER BY
  -- by title
  CASE WHEN sqlc.narg('order_by_field')::text = 'title' AND sqlc.narg('order_type')::text = 'asc' THEN title END ASC,
  CASE WHEN sqlc.narg('order_by_field')::text = 'title' AND sqlc.narg('order_type')::text = 'desc' THEN title END DESC,
  
  -- by created_at
  CASE WHEN sqlc.narg('order_by_field')::text = 'created_at' AND sqlc.narg('order_type')::text = 'asc' THEN created_at END ASC,
  CASE WHEN sqlc.narg('order_by_field')::text = 'created_at' AND sqlc.narg('order_type')::text = 'desc' THEN created_at END DESC,
  
  -- by updated_at
  CASE WHEN sqlc.narg('order_by_field')::text = 'updated_at' AND sqlc.narg('order_type')::text = 'asc' THEN updated_at END ASC,
  CASE WHEN sqlc.narg('order_by_field')::text = 'updated_at' AND sqlc.narg('order_type')::text = 'desc' THEN updated_at END DESC,
  
  -- default
  created_at DESC
LIMIT
  sqlc.arg('limit')::integer OFFSET sqlc.arg('offset')::integer;