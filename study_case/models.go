// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package study_case

import (
	"database/sql"

	"github.com/sqlc-dev/pqtype"
)

type FlashCard struct {
	ID          int64
	Front       sql.NullString
	Back        sql.NullString
	NextReview  sql.NullTime
	StudyCaseID sql.NullInt64
	CreatedAt   sql.NullTime
}

type StudyCase struct {
	ID          int64
	Name        string
	Tags        pqtype.NullRawMessage
	Comment     sql.NullString
	Description sql.NullString
	NextReview  sql.NullTime
	CreatedAt   sql.NullTime
}
