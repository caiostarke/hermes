package entity

import (
	"errors"
	"time"
)

var (
	ErrTextIsRequired = errors.New("text is required")
)

type Text struct {
	ID         int64
	Text       string
	Created_at time.Time
}

func NewText(text string) (*Text, error) {
	txt := &Text{
		Text: text,
	}

	if err := txt.Validate(); err != nil {
		return nil, err
	}

	return txt, nil
}

func (t *Text) Validate() error {
	if t.Text == "" {
		return ErrTextIsRequired
	}

	return nil
}
