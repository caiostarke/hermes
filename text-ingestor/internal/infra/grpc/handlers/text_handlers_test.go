package handlers

import (
	"fmt"
	"testing"

	"github.com/caiostarke/hermes/text-ingestor/internal/entity"
	"github.com/caiostarke/hermes/text-ingestor/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupInMemoryDB(autoMigrateValue any) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(autoMigrateValue)

	return db, nil
}

func TestReturnFilteredTexts(t *testing.T) {
	db, err := setupInMemoryDB(entity.Text{})
	if err != nil {
		t.Error(err)
	}

	texts := []entity.Text{
		{Text: "Im gonna eat"},
		{Text: "Im gonna sit"},
		{Text: "Im gonna sleep"},
		{Text: "Apple"},
		{Text: "Orange"},
	}

	db.Create(&texts)
	textDB := database.NewText(db)

	handler := NewTextHandler(textDB)

	v := []entity.Text{}
	db.Find(&v)

	values, err := handler.ReturnFilteredTexts()
	assert.Nil(t, err)
	assert.True(t, ContainText(values, "Im gonna eat"))
	assert.True(t, ContainText(values, "Im gonna sit"))
	assert.True(t, ContainText(values, "Im gonna sleep"))
	assert.False(t, ContainText(values, "Apple"))

	fmt.Println(values)
}


func TestContainText(t *testing.T) {
	texts := []entity.Text{
		{Text: "lol"},
		{Text: "cc"},
	}

	tableTest := []struct {
		Name          string
		Value         string
		ExpectedValue bool
	}{
		{Name: "Test if lol exists", Value: "lol", ExpectedValue: true},
		{Name: "Test if cc exists", Value: "cc", ExpectedValue: true},
		{Name: "Test if lol exists with white spaces in right and left", Value: " lol ", ExpectedValue: true},
		{Name: "Test if empty exists ", Value: " ", ExpectedValue: false},
		{Name: "Test if a false value is false", Value: "Damn it god", ExpectedValue: false},
		{Name: "Test if a value containing a cc prefix is returned as true", Value: "cc xd", ExpectedValue: false},
	}

	for _, v := range tableTest {
		t.Run(v.Name, func(t *testing.T) {
			contain := ContainText(texts, v.Value)
			if contain != v.ExpectedValue {
				t.Errorf("test: %s : expected %t, got: %t", v.Name, v.ExpectedValue, contain)
			}
		})
	}
}