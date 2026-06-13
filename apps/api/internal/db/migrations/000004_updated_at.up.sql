-- Generic updated_at trigger function
CREATE OR REPLACE FUNCTION set_updated_at() RETURNS trigger AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Apply to all tables with updated_at
CREATE TRIGGER set_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER set_problems_updated_at
    BEFORE UPDATE ON problems
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER set_submissions_updated_at
    BEFORE UPDATE ON submissions
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER set_editorials_updated_at
    BEFORE UPDATE ON editorials
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER set_problem_lists_updated_at
    BEFORE UPDATE ON problem_lists
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER set_problem_notes_updated_at
    BEFORE UPDATE ON problem_notes
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER set_discuss_posts_updated_at
    BEFORE UPDATE ON discuss_posts
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
