-- name: CreateSubmission :one
INSERT INTO submissions (user_id, problem_id, language, code)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateSubmissionVerdict :one
UPDATE submissions
SET
    verdict           = $2,
    runtime_ms        = $3,
    memory_kb         = $4,
    error_message     = $5,
    passed_test_count = $6,
    total_test_count  = $7
WHERE id = $1
RETURNING *;

-- name: GetSubmissionByID :one
SELECT
    s.*,
    p.slug  AS problem_slug,
    p.title AS problem_title
FROM submissions s
JOIN problems p ON p.id = s.problem_id
WHERE s.id = $1;

-- name: ListUserSubmissions :many
SELECT
    s.id, s.language, s.verdict, s.runtime_ms, s.memory_kb, s.created_at,
    p.slug AS problem_slug, p.title AS problem_title
FROM submissions s
JOIN problems p ON p.id = s.problem_id
WHERE
    s.user_id = $1
    AND ($2::uuid IS NULL OR s.problem_id = $2)
    AND ($3::submission_verdict IS NULL OR s.verdict = $3)
ORDER BY s.created_at DESC
LIMIT $4 OFFSET $5;

-- name: InsertSubmissionTestResult :one
INSERT INTO submission_test_results
    (submission_id, test_case_id, input, expected, actual, passed, runtime_ms, memory_kb)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetSubmissionTestResults :many
SELECT * FROM submission_test_results WHERE submission_id = $1;
