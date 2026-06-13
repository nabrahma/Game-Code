-- Trigger to keep search_vector up to date
CREATE OR REPLACE FUNCTION problems_fts_update() RETURNS trigger AS $$
BEGIN
    NEW.search_vector :=
        setweight(to_tsvector('english', coalesce(NEW.title, '')), 'A') ||
        setweight(to_tsvector('english', coalesce(NEW.description, '')), 'B');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER problems_fts_trigger
BEFORE INSERT OR UPDATE OF title, description
ON problems
FOR EACH ROW EXECUTE FUNCTION problems_fts_update();

-- Backfill existing rows
UPDATE problems SET updated_at = updated_at;
