BEGIN;

ALTER TABLE users DROP COLUMN is_premium;

COMMIT;