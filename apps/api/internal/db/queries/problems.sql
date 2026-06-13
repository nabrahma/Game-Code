-- name: ListPublishedProblems :many
-- Used by: ProblemRepository.List
SELECT
    p.id, p.slug, p.title, p.difficulty, p.acceptance_rate,
    p.submission_count, p.order_index, p.updated_at,
    -- Aggregate tags as JSON array
    COALESCE(
        json_agg(DISTINCT jsonb_build_object('id', t.id, 'name', t.name, 'slug', t.slug))
        FILTER (WHERE t.id IS NOT NULL),
        '[]'
    ) AS tags,
    -- Aggregate languages as array
    COALESCE(array_agg(DISTINCT pl.language::text) FILTER (WHERE pl.language IS NOT NULL), '{}') AS languages
FROM problems p
LEFT JOIN problem_tags pt   ON pt.problem_id = p.id
LEFT JOIN tags t            ON t.id = pt.tag_id
LEFT JOIN problem_languages pl ON pl.problem_id = p.id
WHERE
    p.status = 'published'
    AND ($1::difficulty IS NULL OR p.difficulty = $1)
    AND ($2::text IS NULL OR p.search_vector @@ plainto_tsquery('english', $2))
GROUP BY p.id
ORDER BY
    CASE WHEN $3::text = 'newest'          THEN p.published_at END DESC,
    CASE WHEN $3::text = 'acceptance_asc'  THEN p.acceptance_rate END ASC,
    CASE WHEN $3::text = 'acceptance_desc' THEN p.acceptance_rate END DESC,
    p.order_index ASC
LIMIT $4 OFFSET $5;

-- name: CountPublishedProblems :one
SELECT COUNT(*) FROM problems p
WHERE
    p.status = 'published'
    AND ($1::difficulty IS NULL OR p.difficulty = $1)
    AND ($2::text IS NULL OR p.search_vector @@ plainto_tsquery('english', $2));

-- name: GetProblemBySlug :one
SELECT
    p.*,
    COALESCE(
        json_agg(DISTINCT jsonb_build_object('id', t.id, 'name', t.name, 'slug', t.slug, 'category', t.category))
        FILTER (WHERE t.id IS NOT NULL), '[]'
    ) AS tags,
    COALESCE(array_agg(DISTINCT pl.language::text) FILTER (WHERE pl.language IS NOT NULL), '{}') AS languages
FROM problems p
LEFT JOIN problem_tags pt      ON pt.problem_id = p.id
LEFT JOIN tags t               ON t.id = pt.tag_id
LEFT JOIN problem_languages pl ON pl.problem_id = p.id
WHERE p.slug = $1 AND p.status = 'published'
GROUP BY p.id;

-- name: GetProblemExamples :many
SELECT * FROM problem_examples WHERE problem_id = $1 ORDER BY order_index;

-- name: GetProblemHints :many
SELECT * FROM problem_hints WHERE problem_id = $1 ORDER BY order_index;

-- name: GetStarterCode :many
SELECT * FROM starter_code WHERE problem_id = $1;

-- name: GetVisibleTestCases :many
SELECT * FROM test_cases WHERE problem_id = $1 AND is_hidden = FALSE ORDER BY order_index;

-- name: GetAllTestCases :many
SELECT * FROM test_cases WHERE problem_id = $1 ORDER BY order_index;

-- name: GetEditorial :one
SELECT * FROM editorials WHERE problem_id = $1;

-- name: GetRelatedProblems :many
SELECT p.id, p.slug, p.title, p.difficulty, p.acceptance_rate
FROM problems p
JOIN problem_relations pr ON pr.target_id = p.id
WHERE pr.source_id = $1 AND p.status = 'published'
LIMIT 5;

-- name: IsFavorite :one
SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id = $1 AND problem_id = $2);

-- name: GetUserProblemStatus :one
SELECT status FROM user_problem_progress WHERE user_id = $1 AND problem_id = $2;

-- name: ToggleFavorite :exec
-- Called as two separate queries: insert or delete
INSERT INTO favorites (user_id, problem_id) VALUES ($1, $2)
ON CONFLICT (user_id, problem_id) DO NOTHING;

-- name: RemoveFavorite :exec
DELETE FROM favorites WHERE user_id = $1 AND problem_id = $2;
