CREATE TYPE jobtype AS ENUM ('Simple', 'GenerateFibonacci');
CREATE TYPE jobstatus AS ENUM ('Scheduled', 'Finished');

CREATE TABLE IF NOT EXISTS jobs (
    id UUID PRIMARY KEY,
    email_id VARCHAR(256) NOT NULL,
    job_type jobtype NOT NULL,
    payload BYTEA CHECK (LENGTH(payload) <= 2000) NOT NULL,
    recurring BOOLEAN NOT NULL,
    frequency INTEGER NULL,
    next_run TIMESTAMP NULL,
    timeout_after INTEGER NOT NULL,
    job_status jobstatus NOT NULL
);