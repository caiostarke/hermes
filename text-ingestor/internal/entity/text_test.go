package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateText(t *testing.T) {
	txt, err := NewText("Hello my little friend!")
	assert.Nil(t, err)
	assert.NotNil(t, txt)
	assert.Equal(t, "Hello my little friend!", txt.Text)
}

func TestWhenTextIsRequired(t *testing.T) {
	txt, err := NewText("")
	assert.NotNil(t, err)
	assert.Nil(t, txt)

	assert.EqualError(t, err, ErrTextIsRequired.Error())
}
