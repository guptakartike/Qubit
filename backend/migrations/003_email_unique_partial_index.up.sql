-- Drop the auto-generated unique constraint on email.
-- Replace with a partial unique index that ignores soft-deleted rows,
-- allowing re-registration after account deletion.
ALTER TABLE users DROP CONSTRAINT users_email_key;

CREATE UNIQUE INDEX users_email_active_idx
    ON users(email)
    WHERE deleted_at IS NULL;
