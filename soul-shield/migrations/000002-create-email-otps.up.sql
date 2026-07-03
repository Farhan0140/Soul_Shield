-- +migrate Up
CREATE TABLE email_otps(
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    otp VARCHAR(10) NOT NULL,
    purpose VARCHAR(30) NOT NULL CHECK(purpose IN('signup','forgot_password')),
    verified BOOLEAN NOT NULL DEFAULT FALSE,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_email_purpose UNIQUE(email, purpose)
);

CREATE INDEX idx_email_otps_email ON email_otps(email);
CREATE INDEX idx_email_otps_purpose ON email_otps(purpose);
CREATE INDEX idx_email_otps_expires_at ON email_otps(expires_at);