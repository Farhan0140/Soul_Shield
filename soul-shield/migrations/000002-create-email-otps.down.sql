-- +migrate Down
DROP INDEX IF EXISTS idx_email_otps_expires_at;
DROP INDEX IF EXISTS idx_email_otps_purpose;
DROP INDEX IF EXISTS idx_email_otps_email;

DROP TABLE IF EXISTS email_otps;