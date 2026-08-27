DROP INDEX IF EXISTS users_email_active_idx;

ALTER TABLE users ADD CONSTRAINT users_email_key UNIQUE (email);
