-- name: GetArticlesByFilter :many
SELECT
  *
FROM
  article
WHERE
  CASE
    -- by title
    WHEN @filter_search_title :: text IS NULL THEN TRUE
    ELSE title LIKE @filter_search_title
  END
  AND CASE
    -- by tags
    WHEN @filter_search_tags :: text [] IS NULL
    OR cardinality(@filter_search_tags) IS NULL
    OR cardinality(@filter_search_tags) = 0 THEN TRUE
    ELSE tags & & @filter_search_tags
  END
  AND CASE
    -- by filter_created_from
    WHEN @filter_created_from :: timestamp IS NULL THEN TRUE
    ELSE created_at >= @filter_created_from
  END
  AND CASE
    -- by filter_created_to
    WHEN @filter_created_to :: timestamp IS NULL THEN TRUE
    ELSE created_at <= @filter_created_to
  END
  AND CASE
    -- by filter_updated_from
    WHEN @filter_updated_from :: timestamp IS NULL THEN TRUE
    ELSE updated_at >= @filter_updated_from
  END
  AND CASE
    -- by filter_updated_to
    WHEN @filter_updated_to :: timestamp IS NULL THEN TRUE
    ELSE updated_at <= @filter_updated_to
  END;