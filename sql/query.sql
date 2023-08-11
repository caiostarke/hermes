-- name: GetStudyCase :one
SELECT * FROM study_case
WHERE id = $1 LIMIT 1;

-- name: ListStudyCases :many
SELECT * FROM study_case
ORDER BY name;

-- name: CreateStudyCase :one
INSERT INTO study_case (
  name, tags, comment, description, next_review
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteStudyCase :exec
DELETE FROM study_case
WHERE id = $1;

-- name: UpdateStudyCase :one
UPDATE study_case
  set name = $2,
  tags = $3,
  comment = $4,
  description = $5,
  next_review = $6
WHERE id = $1
RETURNING *;

-- name: GetFlashCard :one
SELECT * FROM flash_cards
WHERE id = $1 LIMIT 1;

-- name: ListFlashCards :many
SELECT * FROM flash_cards
WHERE study_case_id = $1
ORDER BY id;

-- name: CreateFlashCard :one
INSERT INTO flash_cards (
  front, back, next_review, study_case_id
  ) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteFlashCard :exec
DELETE FROM flash_cards
WHERE id = $1;

-- name: UpdateFlashCard :one
UPDATE flash_cards
  set front = $2,
  back = $3,
  next_review = $4
WHERE id = $1
RETURNING *;
