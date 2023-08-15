package study_case

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	"github.com/sqlc-dev/pqtype"
)

type Repository interface {
	GetStudyCase(ctx context.Context, id int64) (StudyCase, error)
	DeleteStudyCase(ctx context.Context, id int64) error
	ListStudyCases(ctx context.Context) ([]StudyCase, error)
	UpdateStudyCase(ctx context.Context, arg UpdateStudyCaseParams) (StudyCase, error)
	CreateStudyCase(ctx context.Context, arg CreateStudyCaseParams) (StudyCase, error)

	// FlashCard Related
	CreateFlashCard(ctx context.Context, arg CreateFlashCardParams) (FlashCard, error)
	DeleteFlashCard(ctx context.Context, id int64) error
	ListFlashCards(ctx context.Context, studyCaseID sql.NullInt64) ([]FlashCard, error)
}

type Service struct {
	queries Repository
}

func NewService(repo Repository) *Service {
	return &Service{queries: repo}
}

func (service *Service) ListStudyCases() (any, error) {
	data, err := service.queries.ListStudyCases(context.Background())
	if err != nil {
		return nil, err
	}

	return filterData(data), nil
}

func (service *Service) CreateStudyCase(name string, tags json.RawMessage, comment string, description string) (StudyCase, error) {
	s := CreateStudyCaseParams{
		Name: name,
		Tags: pqtype.NullRawMessage{RawMessage: tags, Valid: tags != nil},
		// I guess that is is not necessary populate Valid field, sql.Null[type] there's a method to validate it
		Comment:     sql.NullString{String: comment, Valid: comment != ""},
		Description: sql.NullString{String: description, Valid: description != ""},
	}

	return service.queries.CreateStudyCase(context.Background(), s)
}

func (service *Service) GetStudyCase(id string) (StudyCase, error) {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return StudyCase{}, err
	}

	sc, err := service.queries.GetStudyCase(context.Background(), idInt)
	if err != nil {
		return StudyCase{}, err
	}

	return sc, nil
}

func (s *Service) DeleteStudyCase(id string) error {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	return s.queries.DeleteStudyCase(context.Background(), idInt)
}

func (s *Service) UpdateStudyCase(id string, name string, tags json.RawMessage, comment string, description string, nextReview time.Time) (StudyCase, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return StudyCase{}, err
	}

	studyCase := UpdateStudyCaseParams{
		ID:          intID,
		Name:        name,
		Tags:        pqtype.NullRawMessage{RawMessage: tags},
		Comment:     sql.NullString{String: comment, Valid: comment != ""},
		Description: sql.NullString{String: description, Valid: description != ""},
		NextReview:  sql.NullTime{Time: nextReview, Valid: nextReview.Equal(time.Time{})},
	}

	return s.queries.UpdateStudyCase(context.Background(), studyCase)
}

func filterData(data []StudyCase) any {
	result := []struct {
		ID          int64
		Name        string
		Tags        string
		Comment     string
		Description string
		NextReview  time.Time
		CreatedAt   time.Time
	}{}

	for _, v := range data {
		d := struct {
			ID          int64
			Name        string
			Tags        string
			Comment     string
			Description string
			NextReview  time.Time
			CreatedAt   time.Time
		}{}

		d.ID = v.ID
		d.Name = v.Name
		d.Tags = string(v.Tags.RawMessage)
		d.Comment = v.Comment.String
		d.Description = v.Description.String
		d.NextReview = v.NextReview.Time
		d.CreatedAt = v.CreatedAt.Time

		result = append(result, d)
	}

	return result
}

// Implementation of service of FlasCards
// Would be loose coupled if flashcard repo and services was splitted in different package
func (s *Service) CreateFlashCard(front string, back string, nextReview time.Time, studyCaseID string) (FlashCard, error) {
	id, err := strconv.ParseInt(studyCaseID, 10, 64)
	if err != nil {
		return FlashCard{}, err
	}

	flashCard := CreateFlashCardParams{
		Front:       sql.NullString{String: front, Valid: front != ""},
		Back:        sql.NullString{String: back, Valid: back != ""},
		NextReview:  sql.NullTime{Time: nextReview, Valid: true},
		StudyCaseID: sql.NullInt64{Int64: id, Valid: id != 0},
	}

	return s.queries.CreateFlashCard(context.Background(), flashCard)
}

func (s *Service) DeleteFlashCard(id string) error {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	return s.queries.DeleteFlashCard(context.Background(), idInt)
}

func (s *Service) ListFlashCards(id string) ([]FlashCard, error) {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	return s.queries.ListFlashCards(context.Background(), sql.NullInt64{Int64: idInt, Valid: true})
}
