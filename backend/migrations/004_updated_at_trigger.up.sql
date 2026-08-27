-- Reusable trigger function that sets updated_at to now() on every update.
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER users_set_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER password_credentials_set_updated_at
    BEFORE UPDATE ON password_credentials
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
