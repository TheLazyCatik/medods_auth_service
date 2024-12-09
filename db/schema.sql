CREATE TABLE
    users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        email TEXT NOT NULL
    );

CREATE TABLE
    sessions (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID REFERENCES users (id) ON DELETE CASCADE NOT NULL,
        hash_token TEXT NOT NULL,
        access_token_id UUID NOT NULL,
        expires_at TIMESTAMP NOT NULL
    );

CREATE OR REPLACE FUNCTION clean_expired_sessions()
RETURNS TRIGGER AS $$
BEGIN
  DELETE FROM sessions WHERE expires_at < NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER clean_sessions_trigger
AFTER INSERT OR UPDATE ON sessions
FOR EACH ROW
EXECUTE FUNCTION clean_expired_sessions();