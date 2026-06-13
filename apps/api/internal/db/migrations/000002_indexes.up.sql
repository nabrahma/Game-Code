-- Users
CREATE INDEX idx_users_email    ON users(email);
CREATE INDEX idx_users_username ON users(username);

-- Problems
CREATE INDEX idx_problems_slug       ON problems(slug);
CREATE INDEX idx_problems_status     ON problems(status);
CREATE INDEX idx_problems_difficulty ON problems(difficulty);
CREATE INDEX idx_problems_order      ON problems(status, order_index);
CREATE INDEX idx_problems_status_diff ON problems(status, difficulty);

-- Full text search
CREATE INDEX idx_problems_fts ON problems USING GIN(search_vector);

-- Problem relationships
CREATE INDEX idx_problem_tags_tag     ON problem_tags(tag_id);
CREATE INDEX idx_problem_tags_problem ON problem_tags(problem_id);

-- Submissions
CREATE INDEX idx_submissions_user_id         ON submissions(user_id);
CREATE INDEX idx_submissions_problem_id      ON submissions(problem_id);
CREATE INDEX idx_submissions_user_problem    ON submissions(user_id, problem_id);
CREATE INDEX idx_submissions_created_at      ON submissions(created_at DESC);
CREATE INDEX idx_submission_results_sub      ON submission_test_results(submission_id);

-- Lists
CREATE INDEX idx_problem_list_items_list    ON problem_list_items(list_id);
CREATE INDEX idx_problem_lists_user         ON problem_lists(user_id);
CREATE INDEX idx_problem_lists_curated      ON problem_lists(is_curated);

-- Progress
CREATE INDEX idx_progress_user  ON user_problem_progress(user_id);

-- Discuss
CREATE INDEX idx_discuss_posts_problem     ON discuss_posts(problem_id);
CREATE INDEX idx_discuss_posts_user        ON discuss_posts(user_id);
CREATE INDEX idx_discuss_posts_problem_ts  ON discuss_posts(problem_id, created_at DESC);
CREATE INDEX idx_discuss_comments_post     ON discuss_comments(post_id);

-- OAuth
CREATE INDEX idx_oauth_user ON oauth_accounts(user_id);

-- Refresh tokens
CREATE INDEX idx_refresh_tokens_user ON refresh_tokens(user_id);
