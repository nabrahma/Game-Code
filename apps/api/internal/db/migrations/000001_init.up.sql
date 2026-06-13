-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "pgcrypto";    -- gen_random_uuid()
CREATE EXTENSION IF NOT EXISTS "pg_trgm";     -- trigram full-text search

-- ──────────────────────────────────────────────────────────────────
-- ENUMS
-- ──────────────────────────────────────────────────────────────────

CREATE TYPE difficulty AS ENUM ('easy', 'medium', 'hard');

CREATE TYPE problem_status AS ENUM ('draft', 'review', 'published', 'archived');

CREATE TYPE submission_verdict AS ENUM (
    'pending',
    'accepted',
    'wrong_answer',
    'time_limit_exceeded',
    'memory_limit_exceeded',
    'runtime_error',
    'compile_error',
    'internal_error'
);

CREATE TYPE language AS ENUM ('csharp', 'cpp', 'lua', 'gdscript');

CREATE TYPE user_role AS ENUM ('user', 'admin', 'content_editor');

CREATE TYPE post_type AS ENUM ('solution', 'question', 'note');

-- ──────────────────────────────────────────────────────────────────
-- USERS
-- ──────────────────────────────────────────────────────────────────

CREATE TABLE users (
    id              UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    email           TEXT        UNIQUE NOT NULL,
    email_verified  TIMESTAMPTZ,
    name            TEXT,
    username        TEXT        UNIQUE NOT NULL,
    avatar_url      TEXT,
    role            user_role   NOT NULL DEFAULT 'user',
    bio             TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- OAuth accounts linked to a user
CREATE TABLE oauth_accounts (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id             UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider            TEXT NOT NULL,   -- 'github' | 'google'
    provider_account_id TEXT NOT NULL,   -- provider's user id
    access_token        TEXT,
    refresh_token       TEXT,
    expires_at          TIMESTAMPTZ,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (provider, provider_account_id)
);

-- Refresh tokens (long-lived, stored server-side for revocation)
CREATE TABLE refresh_tokens (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash  TEXT        NOT NULL UNIQUE,   -- bcrypt hash of the token
    expires_at  TIMESTAMPTZ NOT NULL,
    revoked_at  TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Magic link tokens (email login)
CREATE TABLE magic_link_tokens (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    email       TEXT        NOT NULL,
    token_hash  TEXT        NOT NULL UNIQUE,
    expires_at  TIMESTAMPTZ NOT NULL,
    used_at     TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- ──────────────────────────────────────────────────────────────────
-- TAGS
-- ──────────────────────────────────────────────────────────────────

CREATE TABLE tags (
    id          UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT    UNIQUE NOT NULL,   -- "Pathfinding"
    slug        TEXT    UNIQUE NOT NULL,   -- "pathfinding"
    category    TEXT,                       -- "AI", "Math", "Rendering"
    description TEXT,
    order_index INT     NOT NULL DEFAULT 0
);

-- ──────────────────────────────────────────────────────────────────
-- PROBLEMS
-- ──────────────────────────────────────────────────────────────────

CREATE TABLE problems (
    id                  UUID            PRIMARY KEY DEFAULT gen_random_uuid(),
    slug                TEXT            UNIQUE NOT NULL,
    title               TEXT            NOT NULL,
    difficulty          difficulty      NOT NULL,
    status              problem_status  NOT NULL DEFAULT 'draft',
    description         TEXT            NOT NULL DEFAULT '',
    constraints         TEXT            NOT NULL DEFAULT '',
    game_engine_context TEXT,
    acceptance_rate     NUMERIC(5,2)    NOT NULL DEFAULT 0,
    submission_count    INT             NOT NULL DEFAULT 0,
    accepted_count      INT             NOT NULL DEFAULT 0,
    order_index         INT             NOT NULL DEFAULT 0,
    created_at          TIMESTAMPTZ     NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ     NOT NULL DEFAULT now(),
    published_at        TIMESTAMPTZ,
    -- Full-text search vector (updated via trigger)
    search_vector       TSVECTOR
);

CREATE TABLE problem_tags (
    problem_id  UUID NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    tag_id      UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (problem_id, tag_id)
);

CREATE TABLE problem_languages (
    problem_id  UUID     NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    language    language NOT NULL,
    PRIMARY KEY (problem_id, language)
);

CREATE TABLE problem_examples (
    id          UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    problem_id  UUID    NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    order_index INT     NOT NULL DEFAULT 0,
    input       TEXT    NOT NULL,
    output      TEXT    NOT NULL,
    explanation TEXT
);

CREATE TABLE problem_hints (
    id          UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    problem_id  UUID    NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    order_index INT     NOT NULL DEFAULT 0,
    content     TEXT    NOT NULL
);

CREATE TABLE test_cases (
    id           UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    problem_id   UUID    NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    input        TEXT    NOT NULL,
    output       TEXT    NOT NULL,
    is_hidden    BOOLEAN NOT NULL DEFAULT TRUE,
    order_index  INT     NOT NULL DEFAULT 0,
    time_limit   INT     NOT NULL DEFAULT 2000,   -- milliseconds
    memory_limit INT     NOT NULL DEFAULT 262144  -- KB (256 MB)
);

CREATE TABLE starter_code (
    id          UUID     PRIMARY KEY DEFAULT gen_random_uuid(),
    problem_id  UUID     NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    language    language NOT NULL,
    code        TEXT     NOT NULL,
    UNIQUE (problem_id, language)
);

CREATE TABLE editorials (
    id                UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    problem_id        UUID    UNIQUE NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    content           TEXT    NOT NULL DEFAULT '',
    time_complexity   TEXT,
    space_complexity  TEXT,
    author_notes      TEXT,
    unity_variant     TEXT,
    unreal_variant    TEXT,
    godot_variant     TEXT,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE problem_relations (
    source_id   UUID NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    target_id   UUID NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    reason      TEXT,
    PRIMARY KEY (source_id, target_id)
);

-- ──────────────────────────────────────────────────────────────────
-- SUBMISSIONS
-- ──────────────────────────────────────────────────────────────────

CREATE TABLE submissions (
    id                UUID                PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id           UUID                NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    problem_id        UUID                NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    language          language            NOT NULL,
    code              TEXT                NOT NULL,
    verdict           submission_verdict  NOT NULL DEFAULT 'pending',
    runtime_ms        INT,
    memory_kb         INT,
    error_message     TEXT,
    passed_test_count INT,
    total_test_count  INT,
    created_at        TIMESTAMPTZ         NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ         NOT NULL DEFAULT now()
);

CREATE TABLE submission_test_results (
    id              UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    submission_id   UUID    NOT NULL REFERENCES submissions(id) ON DELETE CASCADE,
    test_case_id    UUID    REFERENCES test_cases(id) ON DELETE SET NULL,
    input           TEXT    NOT NULL,
    expected        TEXT    NOT NULL,
    actual          TEXT    NOT NULL DEFAULT '',
    passed          BOOLEAN NOT NULL DEFAULT FALSE,
    runtime_ms      INT,
    memory_kb       INT
);

-- ──────────────────────────────────────────────────────────────────
-- FAVORITES & LISTS
-- ──────────────────────────────────────────────────────────────────

CREATE TABLE favorites (
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    problem_id  UUID NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, problem_id)
);

CREATE TABLE problem_lists (
    id          UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID    REFERENCES users(id) ON DELETE SET NULL,
    title       TEXT    NOT NULL,
    slug        TEXT    UNIQUE NOT NULL,
    description TEXT,
    is_public   BOOLEAN NOT NULL DEFAULT TRUE,
    is_curated  BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE problem_list_items (
    list_id     UUID NOT NULL REFERENCES problem_lists(id) ON DELETE CASCADE,
    problem_id  UUID NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    order_index INT  NOT NULL DEFAULT 0,
    added_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (list_id, problem_id)
);

-- ──────────────────────────────────────────────────────────────────
-- USER PROGRESS
-- ──────────────────────────────────────────────────────────────────

CREATE TABLE user_problem_progress (
    user_id       UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    problem_id    UUID NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    status        TEXT NOT NULL DEFAULT 'attempted', -- 'attempted' | 'solved'
    attempt_count INT  NOT NULL DEFAULT 0,
    last_attempt  TIMESTAMPTZ NOT NULL DEFAULT now(),
    solved_at     TIMESTAMPTZ,
    PRIMARY KEY (user_id, problem_id)
);

CREATE TABLE problem_notes (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    problem_id  UUID NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    content     TEXT NOT NULL DEFAULT '',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (user_id, problem_id)
);

-- ──────────────────────────────────────────────────────────────────
-- DISCUSS
-- ──────────────────────────────────────────────────────────────────

CREATE TABLE discuss_posts (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    problem_id  UUID        NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    user_id     UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type        post_type   NOT NULL DEFAULT 'solution',
    title       TEXT        NOT NULL,
    content     TEXT        NOT NULL,
    language    language,
    upvotes     INT         NOT NULL DEFAULT 0,
    is_pinned   BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE discuss_comments (
    id          UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    post_id     UUID    NOT NULL REFERENCES discuss_posts(id) ON DELETE CASCADE,
    user_id     UUID    NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content     TEXT    NOT NULL,
    upvotes     INT     NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
