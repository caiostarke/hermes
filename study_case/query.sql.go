// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package study_case

import (
	"context"
	"database/sql"

	"github.com/sqlc-dev/pqtype"
)

const createFlashCard = `-- name: CreateFlashCard :one
INSERT INTO flash_cards (
  front, back, next_review, study_case_id
  ) VALUES (
  $1, $2, $3, $4
)
RETURNING id, front, back, next_review, study_case_id, created_at
`

type CreateFlashCardParams struct {
	Front       sql.NullString
	Back        sql.NullString
	NextReview  sql.NullTime
	StudyCaseID sql.NullInt64
}

func (q *Queries) CreateFlashCard(ctx context.Context, arg CreateFlashCardParams) (FlashCard, error) {
	row := q.db.QueryRowContext(ctx, createFlashCard,
		arg.Front,
		arg.Back,
		arg.NextReview,
		arg.StudyCaseID,
	)
	var i FlashCard
	err := row.Scan(
		&i.ID,
		&i.Front,
		&i.Back,
		&i.NextReview,
		&i.StudyCaseID,
		&i.CreatedAt,
	)
	return i, err
}

const createStudyCase = `-- name: CreateStudyCase :one
INSERT INTO study_case (
  name, tags, comment, description, next_review
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, name, tags, comment, description, next_review, created_at
`

type CreateStudyCaseParams struct {
	Name        string
	Tags        pqtype.NullRawMessage
	Comment     sql.NullString
	Description sql.NullString
	NextReview  sql.NullTime
}

func (q *Queries) CreateStudyCase(ctx context.Context, arg CreateStudyCaseParams) (StudyCase, error) {
	row := q.db.QueryRowContext(ctx, createStudyCase,
		arg.Name,
		arg.Tags,
		arg.Comment,
		arg.Description,
		arg.NextReview,
	)
	var i StudyCase
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Tags,
		&i.Comment,
		&i.Description,
		&i.NextReview,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFlashCard = `-- name: DeleteFlashCard :exec
DELETE FROM flash_cards
WHERE id = $1
`

func (q *Queries) DeleteFlashCard(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFlashCard, id)
	return err
}

const deleteStudyCase = `-- name: DeleteStudyCase :exec
DELETE FROM study_case
WHERE id = $1
`

func (q *Queries) DeleteStudyCase(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteStudyCase, id)
	return err
}

const getFlashCard = `-- name: GetFlashCard :one
SELECT id, front, back, next_review, study_case_id, created_at FROM flash_cards
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFlashCard(ctx context.Context, id int64) (FlashCard, error) {
	row := q.db.QueryRowContext(ctx, getFlashCard, id)
	var i FlashCard
	err := row.Scan(
		&i.ID,
		&i.Front,
		&i.Back,
		&i.NextReview,
		&i.StudyCaseID,
		&i.CreatedAt,
	)
	return i, err
}

const getStudyCase = `-- name: GetStudyCase :one
SELECT id, name, tags, comment, description, next_review, created_at FROM study_case
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStudyCase(ctx context.Context, id int64) (StudyCase, error) {
	row := q.db.QueryRowContext(ctx, getStudyCase, id)
	var i StudyCase
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Tags,
		&i.Comment,
		&i.Description,
		&i.NextReview,
		&i.CreatedAt,
	)
	return i, err
}

const listFlashCards = `-- name: ListFlashCards :many
SELECT id, front, back, next_review, study_case_id, created_at FROM flash_cards
WHERE study_case_id = $1
ORDER BY id
`

func (q *Queries) ListFlashCards(ctx context.Context, studyCaseID sql.NullInt64) ([]FlashCard, error) {
	rows, err := q.db.QueryContext(ctx, listFlashCards, studyCaseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FlashCard
	for rows.Next() {
		var i FlashCard
		if err := rows.Scan(
			&i.ID,
			&i.Front,
			&i.Back,
			&i.NextReview,
			&i.StudyCaseID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listStudyCases = `-- name: ListStudyCases :many
SELECT id, name, tags, comment, description, next_review, created_at FROM study_case
ORDER BY name
`

func (q *Queries) ListStudyCases(ctx context.Context) ([]StudyCase, error) {
	rows, err := q.db.QueryContext(ctx, listStudyCases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StudyCase
	for rows.Next() {
		var i StudyCase
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Tags,
			&i.Comment,
			&i.Description,
			&i.NextReview,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFlashCard = `-- name: UpdateFlashCard :one
UPDATE flash_cards
  set front = $2,
  back = $3,
  next_review = $4
WHERE id = $1
RETURNING id, front, back, next_review, study_case_id, created_at
`

type UpdateFlashCardParams struct {
	ID         int64
	Front      sql.NullString
	Back       sql.NullString
	NextReview sql.NullTime
}

func (q *Queries) UpdateFlashCard(ctx context.Context, arg UpdateFlashCardParams) (FlashCard, error) {
	row := q.db.QueryRowContext(ctx, updateFlashCard,
		arg.ID,
		arg.Front,
		arg.Back,
		arg.NextReview,
	)
	var i FlashCard
	err := row.Scan(
		&i.ID,
		&i.Front,
		&i.Back,
		&i.NextReview,
		&i.StudyCaseID,
		&i.CreatedAt,
	)
	return i, err
}

const updateStudyCase = `-- name: UpdateStudyCase :one
UPDATE study_case
  set name = $2,
  tags = $3,
  comment = $4,
  description = $5,
  next_review = $6
WHERE id = $1
RETURNING id, name, tags, comment, description, next_review, created_at
`

type UpdateStudyCaseParams struct {
	ID          int64
	Name        string
	Tags        pqtype.NullRawMessage
	Comment     sql.NullString
	Description sql.NullString
	NextReview  sql.NullTime
}

func (q *Queries) UpdateStudyCase(ctx context.Context, arg UpdateStudyCaseParams) (StudyCase, error) {
	row := q.db.QueryRowContext(ctx, updateStudyCase,
		arg.ID,
		arg.Name,
		arg.Tags,
		arg.Comment,
		arg.Description,
		arg.NextReview,
	)
	var i StudyCase
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Tags,
		&i.Comment,
		&i.Description,
		&i.NextReview,
		&i.CreatedAt,
	)
	return i, err
}