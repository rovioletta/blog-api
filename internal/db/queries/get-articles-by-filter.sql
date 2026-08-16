-- name: GetArticlesByFilter :many
SELECT
  *
FROM
  article
WHERE
  -- by title
  (@filter_search_title::text IS NULL OR title LIKE @filter_search_title::text)
  
  -- by tags
  AND (
    @filter_search_tags::text[] IS NULL 
    OR cardinality(@filter_search_tags::text[]) = 0 
    OR tags && @filter_search_tags::text[]
  )
  
  -- by created_at
  AND (@filter_created_from::timestamp IS NULL OR created_at >= @filter_created_from::timestamp)
  AND (@filter_created_to::timestamp IS NULL OR created_at <= @filter_created_to::timestamp)
  
  -- by updated_at
  AND (@filter_updated_from::timestamp IS NULL OR updated_at >= @filter_updated_from::timestamp)
  AND (@filter_updated_to::timestamp IS NULL OR updated_at <= @filter_updated_to::timestamp)
ORDER BY
  -- by title
  CASE WHEN @order_by_field::text = 'title' AND @order_type::text = 'asc' THEN title END ASC,
  CASE WHEN @order_by_field::text = 'title' AND @order_type::text = 'desc' THEN title END DESC,
  
  -- by created_at
  CASE WHEN @order_by_field::text = 'created_at' AND @order_type::text = 'asc' THEN created_at END ASC,
  CASE WHEN @order_by_field::text = 'created_at' AND @order_type::text = 'desc' THEN created_at END DESC,
  
  -- by updated_at
  CASE WHEN @order_by_field::text = 'updated_at' AND @order_type::text = 'asc' THEN updated_at END ASC,
  CASE WHEN @order_by_field::text = 'updated_at' AND @order_type::text = 'desc' THEN updated_at END DESC,
  
  -- default
  created_at DESC
LIMIT
  @limit_page::integer OFFSET @offset_page::integer;