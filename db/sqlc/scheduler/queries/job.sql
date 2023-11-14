-- name: CreateJob :one
INSERT INTO jobs (
  id, email_id, job_type, payload, recurring, frequency, next_run, timeout_after, job_status   
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetJobById :one
SELECT * FROM jobs
WHERE id = $1;